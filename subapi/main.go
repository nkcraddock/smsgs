package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
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

	m.Map(p)

	m.Get("/webhooks", func() string {
		response, _ := json.Marshal(p.GetHooks())
		return string(response)
	})

	m.Post("/webhooks", AddWebhook)

	m.RunOnAddr(":3001")
}

func AddWebhook(res http.ResponseWriter, req *http.Request, p webhooks.Persister) {
	dec := json.NewDecoder(req.Body)
	var hook webhooks.Webhook
	err := dec.Decode(&hook)

	if err != nil {
		fmt.Fprintf(res, "AddWebhook failed: %s", err)
		return
	}

	p.AddHook(hook)
}
