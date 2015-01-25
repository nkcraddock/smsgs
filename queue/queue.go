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
	Publisher string
	Key       string
	EventType string
	Payload   interface{}
}
