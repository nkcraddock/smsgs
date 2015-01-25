package webhooks

import "fmt"

type Persister interface {
	AddHook(hook Webhook)
	GetHooks() []Webhook
}

type MemPersister struct {
	Hooks map[string]map[string]Webhook // Hooks["URL"]["TOPIC"]Webhook = {}
}

func NewMemPersister() Persister {
	return MemPersister{
		Hooks: make(map[string]map[string]Webhook),
	}
}

func (p MemPersister) AddHook(hook Webhook) {
	t := hook.Topic()
	h := p.GetSub(hook.Url)

	if _, ok := h[t]; !ok {
		h[t] = hook
		fmt.Println("Added Hook:", hook)
	}
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

func (p MemPersister) GetSub(url string) map[string]Webhook {
	if sub, ok := p.Hooks[url]; ok {
		return sub
	}

	sub := make(map[string]Webhook)

	p.Hooks[url] = sub

	return sub
}
