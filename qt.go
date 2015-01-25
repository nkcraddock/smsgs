package main

import (
	"fmt"
	"math/rand"

	"github.com/nkcraddock/smessages/queue"
)

const publisher = "qt"

func main() {
	q := queue.NewRabbitPublisher("amqp://guest:guest@172.17.0.24:5672", "smsgs.evt")
	c := q.Channel()
	defer q.Close()
	for {
		go func() {
			msg := createRandomMessage()
			c <- *msg
		}()
	}
}

func createRandomMessage() *queue.Event {
	types := []string{"user", "group", "computer", "gun"}
	keys := createKeys(100)

	t := types[rand.Intn(len(types))]
	k := keys[rand.Intn(len(keys))]

	var p string

	switch t {
	case "user":
		p = fmt.Sprintf(`{ "uid" : "%d" }`, rand.Intn(100)+1000)
	case "group":
		p = fmt.Sprintf(`{ "gid" : "%d" }`, rand.Intn(100)+1000)
	case "computer":
		p = fmt.Sprintf(`{ "cid" : "%d" }`, rand.Intn(100)+1000)
	}

	return &queue.Event{
		Publisher: publisher,
		EventType: t,
		Key:       k,
		Payload:   p,
	}
}

var letters []rune = []rune("ABCDEFGHIJKLMNOPQRSTUV1234567890")

func createKeys(howmany int) []string {
	k := make([]string, howmany)

	for i := 0; i < len(k); i++ {
		s := make([]rune, 1)

		for c := range s {
			s[c] = letters[rand.Intn(len(letters))]
		}

		k[i] = string(s)
	}

	return k
}