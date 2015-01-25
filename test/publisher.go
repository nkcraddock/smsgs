package main

import "github.com/nkcraddock/smessages/queue"

const (
	rabbitUri = "amqp://guest:guest@localhost:5672"
	exchange  = "smsgs.evt"
	publisher = "qt"
)

func main() {
	c := queue.OpenRabbit(rabbitUri)
	defer c.Close()

	events := c.Publish(exchange)

	for {
		events <- queue.GenerateRandomEvent(publisher)
	}
}
