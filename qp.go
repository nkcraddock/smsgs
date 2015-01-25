package main

import (
	"fmt"

	"github.com/nkcraddock/smessages/queue"
)

func main() {
	q := queue.NewRabbitListener("amqp://guest:guest@172.17.0.24:5672", "smsgs.evt")
	defer q.Close()

	bind(q)

	for x := range q.Channel() {
		fmt.Println("RECEIVED:", x)
	}

}

func bind(q queue.EventListener) {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUV")
	types := []string{"group", "user", "computer"}

	for i := 0; i < len(types); i++ {
		for l := 0; l < len(letters); l++ {
			q.Listen(fmt.Sprintf("qt.%s.%s", types[i], string(letters[l])))
		}
	}

}
