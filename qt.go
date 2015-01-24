package main

import (
	"fmt"

	"github.com/nkcraddock/smessages/queue"
)

func main() {
	q := queue.NewRabbit("amqp://guest:guest@172.17.0.24:5672")
	defer q.Close()
	for i := 0; i < 10; i++ {
		publishMessage(q, "change", fmt.Sprintf("{ a: 5, b: %d}", i))
	}
}

func publishMessage(q queue.EventPublisher, eventType string, payload string) {
	q.Publish(&queue.Event{
		Publisher: "yourmom",
		EventType: eventType,
		Payload:   payload,
	})

}
