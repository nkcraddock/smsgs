package main

import (
	"fmt"

	"github.com/nkcraddock/smessages/queue"
)

const (
	rabbitUri = "amqp://guest:guest@172.17.0.24:5672"
	sub       = "sub-qt"
	exch      = "smsgs.evt"
)

func main() {
	q := queue.OpenRabbit(rabbitUri)

	defer q.Close()

	in := q.Listen(sub)

	for x := range in {
		fmt.Println("RECEIVED:", x)
	}
}
