package queue

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type Q interface {
	Publish(string) chan<- Event
	//	Listen(string, <-chan Event)
	//	Bind(string, string, string)
	Close()
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
	fmt.Println("PUBLISHED:", string(b))
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
