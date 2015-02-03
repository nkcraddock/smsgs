package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nkcraddock/smsgs/queue"
	"github.com/nu7hatch/gouuid"
)

const (
	rabbitUri = "amqp://guest:guest@localhost:5672"
	sub       = "sub-qt"
	hookUrl   = "http://localhost:3001/mock-subscriber"
	exch      = "smsgs.evt"
)

func main() {
	q := queue.OpenRabbit(rabbitUri)

	defer q.Close()

	queue := getId(hookUrl)
	in := q.Listen(queue)

	for evt := range in {
		b, _ := json.Marshal(evt)
		http.Post(hookUrl, "application/json", bytes.NewBuffer(b))
		fmt.Println("DISPATCHED to", hookUrl)
	}

}

func getId(url string) string {
	u, _ := uuid.NewV5(uuid.NamespaceURL, []byte(url))
	return u.String()
}
