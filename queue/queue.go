package queue

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

const (
	exchangeName = "smsgs.evt"
	bufferSize   = 10
)

type Event struct {
	Publisher string
	Key       string
	EventType string
	Payload   interface{}
}

type EventPublisher interface {
	Close()
	Publish(*Event)
}

type EventListener interface {
	Close()
	Channel() chan Event
	Listen(string)
}

type RabbitEventPublisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbit(uri string) EventPublisher {
	conn, _ := amqp.Dial(uri)
	ch, _ := conn.Channel()
	ch.ExchangeDeclare(exchangeName, "topic", false, false, false, false, nil)
	newPublisher := new(RabbitEventPublisher)
	newPublisher = &RabbitEventPublisher{
		conn: conn,
		ch:   ch,
	}
	return newPublisher
}

// Listener

type RabbitEventListener struct {
	conn   *amqp.Connection
	ch     *amqp.Channel
	Events chan Event
}

func (l RabbitEventListener) Listen(topic string) {
	l.ch.QueueBind(exchangeName, topic, exchangeName, false, nil)
}

func NewRabbitListener(uri string) EventListener {
	conn, _ := amqp.Dial(uri)
	ch, _ := conn.Channel()
	ch.QueueDeclare(exchangeName, false, false, false, false, nil)
	newListener := new(RabbitEventListener)
	newListener = &RabbitEventListener{
		conn: conn,
		ch:   ch,
	}
	newListener.Events = make(chan Event, bufferSize)
	msgs, _ := ch.Consume(exchangeName, "", false, false, false, false, nil)
	go func() {
		for e := range msgs {
			var evt Event
			json.Unmarshal(e.Body, &evt)
			newListener.Events <- evt
			e.Ack(false)
		}
	}()
	return newListener
}

func (l RabbitEventListener) Channel() chan Event {
	return l.Events
}

func (l RabbitEventListener) Close() {
	l.conn.Close()
	l.ch.Close()
}

func body(evt *Event) ([]byte, error) {
	data, _ := json.Marshal(evt)
	return data, nil
}

func (p *RabbitEventPublisher) Publish(evt *Event) {
	topic := fmt.Sprintf("%s.%s.%s", evt.Publisher, evt.EventType, evt.Key)
	b, _ := body(evt)
	p.ch.Publish(
		exchangeName,
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})
	fmt.Println("Published a smsg", string(b))
}

func (p *RabbitEventPublisher) Close() {
	p.ch.Close()
	p.conn.Close()
}
