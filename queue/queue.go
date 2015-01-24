package queue

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Event struct {
	Publisher string
	EventType string
	Payload   string
}

type EventPublisher interface {
	Close()
	Publish(*Event)
}

type RabbitEventPublisher struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewRabbit(uri string) EventPublisher {
	conn, _ := amqp.Dial(uri)
	ch, _ := conn.Channel()
	q, _ := ch.QueueDeclare("evt", false, false, false, false, nil)
	newPublisher := new(RabbitEventPublisher)
	newPublisher = &RabbitEventPublisher{
		conn: conn,
		ch:   ch,
		q:    q,
	}
	return newPublisher
}

func (p *RabbitEventPublisher) Publish(evt *Event) {
	p.ch.Publish(
		"",
		p.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(evt.Payload),
		})
	fmt.Println("Published a smsg")
}

func (p *RabbitEventPublisher) Close() {
	p.ch.Close()
	p.conn.Close()
}
