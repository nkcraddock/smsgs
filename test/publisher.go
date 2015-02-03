package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/nkcraddock/smsgs/queue"
)

const (
	publisher = "qt"
	endpoint  = "http://localhost:3001/events"
)

func main() {
	for {
		evt := queue.GenerateRandomEvent(publisher)
		b, _ := json.Marshal(evt)
		http.Post(endpoint, "application/json", bytes.NewBuffer(b))
	}

}
