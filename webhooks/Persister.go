package webhooks

import (
	"fmt"

	"github.com/nu7hatch/gouuid"
)

type Persister interface {
	AddHook(hook Webhook) bool
	DeleteHook(hook Webhook) bool
	GetHooks() []Webhook
	GetQueue(string) string
}

type MemPersister struct {
	Hooks  map[string]map[string]Webhook // Hooks["URL"]["TOPIC"]Webhook
	Queues map[string]string             // Queues["URL"]hash
}

func NewMemPersister() Persister {
	return MemPersister{
		Hooks:  make(map[string]map[string]Webhook),
		Queues: make(map[string]string),
	}
}

func (p MemPersister) AddHook(hook Webhook) bool {
	t := hook.Topic()
	h := p.GetSub(hook.Url)

	if _, ok := h[t]; !ok {
		h[t] = hook
		fmt.Println("Added Hook:", hook)
		return true
	}

	return false
}

func (p MemPersister) DeleteHook(hook Webhook) bool {
	t := hook.Topic()
	h := p.GetSub(hook.Url)

	if _, ok := h[t]; ok {
		delete(h, t)
		fmt.Println("Removed hook:", hook)
		return true
	}
	return false
}

func (p MemPersister) GetHooks() []Webhook {
	v := make([]Webhook, len(p.Hooks))
	i := 0
	for _, s := range p.Hooks {
		for _, h := range s {
			v[i] = h
			i++
		}
	}
	return v
}

func (p MemPersister) GetQueue(url string) string {
	if id, ok := p.Queues[url]; ok {
		return id
	}
	return ""
}

func (p MemPersister) GetSub(url string) map[string]Webhook {
	if sub, ok := p.Hooks[url]; ok {
		return sub
	}

	sub := make(map[string]Webhook)

	p.Hooks[url] = sub
	p.Queues[url] = getId(url)

	return sub
}

func getId(url string) string {
	u, _ := uuid.NewV5(uuid.NamespaceURL, []byte(url))
	return u.String()
}
