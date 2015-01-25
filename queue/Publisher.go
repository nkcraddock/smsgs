package queue

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type EventPublisher interface {
	Close()
	Channel() chan<- Event
}

// Rabbit

type RabbitEventPublisher struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	exchange string
	c        chan Event
}

func NewRabbitPublisher(uri string, exchange string) EventPublisher {
	conn, _ := amqp.Dial(uri)
	ch, _ := conn.Channel()
	ch.ExchangeDeclare(exchange, "topic", false, false, false, false, nil)
	newPublisher := new(RabbitEventPublisher)
	newPublisher = &RabbitEventPublisher{
		conn:     conn,
		ch:       ch,
		c:        make(chan Event),
		exchange: exchange,
	}

	go func() {
		for event := range newPublisher.c {
			newPublisher.Publish(&event)
		}
	}()

	return newPublisher
}

func body(evt *Event) ([]byte, error) {
	data, _ := json.Marshal(evt)
	return data, nil
}

func (p *RabbitEventPublisher) Channel() chan<- Event {
	return p.c
}

func (p *RabbitEventPublisher) Publish(evt *Event) {
	topic := fmt.Sprintf("%s.%s.%s", evt.Publisher, evt.EventType, evt.Key)
	b, _ := body(evt)
	p.ch.Publish(
		p.exchange,
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
	close(p.c)
	p.ch.Close()
	p.conn.Close()
}
