package webhooks

import "fmt"

type Webhook struct {
	Url string `json:"url"`
	Pub string `json:"pub"`
	Typ string `json:"typ"`
	Key string `json:"key"`
}

func (wh *Webhook) Topic() string {
	return fmt.Sprintf("%s.%s.%s", wh.Pub, wh.Typ, wh.Key)
}
