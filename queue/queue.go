package queue

type Q interface {
	Publish(string) chan<- Event
	Listen(string) <-chan Event
	Subscribe(string)
	Bind(string, string, string)
	Unbind(string, string, string)
	Close()
	Purge(string)
}

type Event struct {
	Publisher string      `json:"publisher"`
	Key       string      `json:"key"`
	EventType string      `json:"eventType"`
	Payload   interface{} `json:"payload"`
}
