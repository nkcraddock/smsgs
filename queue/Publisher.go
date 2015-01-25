package queue

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

type rabbitEventPublisher struct {
	conn     *amqp.Connection
	ch       *amqp.Channel
	exchange string
	c        chan Event
}

func Publisher(uri string, exchange string) chan Event {
	conn, _ := amqp.Dial(uri)
	ch, _ := conn.Channel()
	ch.ExchangeDeclare(exchange, "topic", false, false, false, false, nil)
	newPublisher := new(rabbitEventPublisher)
	newPublisher = &rabbitEventPublisher{
		conn:     conn,
		ch:       ch,
		c:        make(chan Event),
		exchange: exchange,
	}

	go func() {
		for event := range newPublisher.c {
			newPublisher.Publish(&event)
		}
		newPublisher.Close()
	}()

	return newPublisher.c
}

func body(evt *Event) ([]byte, error) {
	data, _ := json.Marshal(evt)
	return data, nil
}

func (p *rabbitEventPublisher) Channel() chan<- Event {
	return p.c
}

func (p *rabbitEventPublisher) Publish(evt *Event) {
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
	fmt.Println("PUBLISHED:", string(b))
}

func (p *rabbitEventPublisher) Close() {
	close(p.c)
	p.ch.Close()
	p.conn.Close()
}
