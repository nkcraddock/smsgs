package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nkcraddock/smsgs/queue"
	"github.com/nkcraddock/smsgs/webhooks"
)

func AddWebhook(res http.ResponseWriter, req *http.Request, p webhooks.Persister, q queue.Q) {
	dec := json.NewDecoder(req.Body)
	var hook webhooks.Webhook
	err := dec.Decode(&hook)

	if err != nil {
		fmt.Fprintf(res, "AddWebhook failed: %s", err)
		return
	}

	if p.AddHook(hook) {
		id := p.GetQueue(hook.Url)
		q.Subscribe(id)
		q.Bind(exchange, id, hook.Topic())
	}
}
func GetWebhooks(res http.ResponseWriter, req *http.Request, p webhooks.Persister) string {
	response, _ := json.Marshal(p.GetHooks())
	return string(response)
}

func DeleteWebhook(res http.ResponseWriter, req *http.Request, p webhooks.Persister, q queue.Q) {
	dec := json.NewDecoder(req.Body)
	var hook webhooks.Webhook
	err := dec.Decode(&hook)

	if err != nil {
		fmt.Fprintf(res, "DeleteWebhook failed: %s", err)
		return
	}

	if p.DeleteHook(hook) {
		id := p.GetQueue(hook.Url)
		q.Unbind(exchange, id, hook.Topic())
	}
}
