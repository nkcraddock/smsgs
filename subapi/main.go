package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/nkcraddock/smsgs/queue"
	"github.com/nkcraddock/smsgs/webhooks"
)

const (
	rabbitUri = "amqp://guest:guest@localhost:5672"
	exchange  = "smsgs.evt"
	publisher = "qt"
)

func main() {
	m := martini.Classic()
	p := webhooks.NewMemPersister()
	q := queue.OpenRabbit(rabbitUri)

	publish := q.Publish(exchange)

	m.Map(p)
	m.Map(q)
	m.Map(publish)

	m.Get("/webhooks", GetWebhooks)
	m.Post("/webhooks", AddWebhook)
	m.Delete("/webhooks", DeleteWebhook)

	m.Post("/events", AddEvent)
	m.Post("/mock-subscriber", ReceivedEvent)

	m.RunOnAddr(":3001")
}

func AddEvent(res http.ResponseWriter, req *http.Request, p chan<- queue.Event) {
	dec := json.NewDecoder(req.Body)
	var evt queue.Event
	err := dec.Decode(&evt)

	if err != nil {
		fmt.Fprintf(res, "AddEvent failed: %s", err)
		return
	}
	go func() {
		p <- evt
	}()
}

func ReceivedEvent(res http.ResponseWriter, req *http.Request) {
	b, _ := ioutil.ReadAll(req.Body)
	fmt.Println("RECEIVED MESSAGE:", string(b))
}
