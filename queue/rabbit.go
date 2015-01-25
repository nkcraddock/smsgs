package queue

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Q interface {
	Publish(string) chan<- Event
	Listen(string) <-chan Event
	Subscribe(string)
	Bind(string, string, string)
	Close()
	Purge(string)
}

type rabbitConnection struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func OpenRabbit(uri string) Q {
	conn, err := amqp.Dial(uri)
	failOnError(err)

	ch, err := conn.Channel()
	failOnError(err)

	return rabbitConnection{
		conn: conn,
		ch:   ch,
	}
}

func (c rabbitConnection) Purge(name string) {
	c.ch.ExchangeDelete(name, false, false)
	c.ch.QueueDelete(name, false, false, false)
}

func (c rabbitConnection) Publish(exchangeName string) chan<- Event {
	c.ch.ExchangeDeclare(exchangeName, "topic", false, false, false, false, nil)

	eventChannel := make(chan Event)

	go func() {
		for event := range eventChannel {
			c.publish(exchangeName, &event)
		}
		c.Close()
	}()

	return eventChannel
}

func (c rabbitConnection) Listen(queueName string) <-chan Event {
	c.Subscribe(queueName)
	eventChannel := make(chan Event)
	msgs, _ := c.ch.Consume(queueName, "", false, false, false, false, nil)
	go func() {
		for e := range msgs {
			var evt Event
			json.Unmarshal(e.Body, &evt)
			eventChannel <- evt
			e.Ack(false)
		}
	}()
	return eventChannel
}

func (c rabbitConnection) Subscribe(queueName string) {
	c.ch.ExchangeDeclare(queueName, "topic", false, false, false, false, nil)
	c.ch.QueueDeclare(queueName, false, false, false, false, nil)
	c.ch.QueueBind(queueName, "#", queueName, false, nil)
}

func (c rabbitConnection) Bind(src string, target string, filter string) {
	fmt.Println("BINDING:", src, target, filter)
	c.ch.ExchangeDeclare(src, "topic", false, false, false, false, nil)
	c.ch.ExchangeDeclare(target, "topic", false, false, false, false, nil)
	c.ch.ExchangeBind(target, filter, src, false, nil)
}

func (c rabbitConnection) Close() {
	c.ch.Close()
	c.conn.Close()
}

func (c rabbitConnection) publish(exchangeName string, evt *Event) {
	b, _ := body(evt)
	c.ch.Publish(
		exchangeName,
		topic(evt),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})
	fmt.Println("PUBLISHED:", topic(evt))
}

func body(evt *Event) ([]byte, error) {
	data, _ := json.Marshal(evt)
	return data, nil
}

func topic(evt *Event) string {
	return fmt.Sprintf("%s.%s.%s", evt.Publisher, evt.EventType, evt.Key)
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}
