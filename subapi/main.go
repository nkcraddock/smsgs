package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/nkcraddock/smessages/queue"
	"github.com/nkcraddock/smessages/webhooks"
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
