package main

import (
	"fmt"

	"github.com/nkcraddock/smessages/queue"
	"github.com/nu7hatch/gouuid"
)

const (
	rabbitUri = "amqp://guest:guest@localhost:5672"
	sub       = "sub-qt"
	hookUrl   = "http://localhost:3001/test"
	exch      = "smsgs.evt"
)

func main() {
	q := queue.OpenRabbit(rabbitUri)

	defer q.Close()

	queue := getId(hookUrl)
	in := q.Listen(queue)

	for x := range in {
		fmt.Println("RECEIVED:", x)
	}

}

func getId(url string) string {
	u, _ := uuid.NewV5(uuid.NamespaceURL, []byte(url))
	return u.String()
}
