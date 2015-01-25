package queue

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type EventListener interface {
	Close()
	Channel() chan Event
	Listen(string)
}

// Rabbit

type RabbitEventListener struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	exchange string
	Events   chan Event
}

func NewRabbitListener(uri string, exchangeName string) EventListener {
	bufferSize := 10
	conn, _ := amqp.Dial(uri)
	ch, _ := conn.Channel()
	ch.QueueDeclare(exchangeName, false, false, false, false, nil)
	newListener := new(RabbitEventListener)
	newListener = &RabbitEventListener{
		conn:     conn,
		ch:       ch,
		exchange: exchangeName,
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

func (l RabbitEventListener) Listen(topic string) {
	l.ch.QueueBind(l.exchange, topic, l.exchange, false, nil)
}

func (l RabbitEventListener) Channel() chan Event {
	return l.Events
}

func (l RabbitEventListener) Close() {
	l.conn.Close()
	l.ch.Close()
}
