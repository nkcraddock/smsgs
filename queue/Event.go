package queue

type Event struct {
	Publisher string
	Key       string
	EventType string
	Payload   interface{}
}
