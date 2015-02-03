package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/nkcraddock/smsgs/webhooks"
)

const (
	publisher = "qt"
)

var endpoint string

func main() {
	port := flag.Int("p", 3001, "the port to listen on")
	flag.StringVar(&endpoint, "u", "http://smsgs/webhooks", "the uri of the webhooks endpoint")
	flag.Parse()

	m := martini.Classic()

	m.Post("/event", ReceivedEvent)

	addWebHook("qt", "user", "*")
	m.RunOnAddr(fmt.Sprintf(":%d", *port))
}

func addWebHook(pub string, typ string, key string) {
	wh := webhooks.Webhook{
		Url: "http://subtest/event",
		Pub: pub,
		Typ: typ,
		Key: key,
	}

	b, _ := json.Marshal(wh)
	http.Post(endpoint, "application/json", bytes.NewBuffer(b))
}

func ReceivedEvent(res http.ResponseWriter, req *http.Request) {
	b, _ := ioutil.ReadAll(req.Body)
	fmt.Println("RECEIVED MESSAGE:", string(b))
}
