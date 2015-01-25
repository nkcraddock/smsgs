package main

import (
	"github.com/go-martini/martini"
	"github.com/nkcraddock/smessages/queue"
	"github.com/nkcraddock/smessages/webhooks"
)

const (
	rabbitUri = "amqp://guest:guest@172.17.0.24:5672"
	exchange  = "smsgs.evt"
	publisher = "qt"
)

func main() {
	m := martini.Classic()
	p := webhooks.NewMemPersister()
	q := queue.OpenRabbit(rabbitUri)

	m.Map(p)
	m.Map(q)

	m.Get("/webhooks", GetWebhooks)
	m.Post("/webhooks", AddWebhook)
	m.Delete("/webhooks", DeleteWebhook)

	m.RunOnAddr(":3001")
}
