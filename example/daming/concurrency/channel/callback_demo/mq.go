package map_demo

import (
	"sync"
)

type Listener func(msg Msg)
type Broker struct {
	mutex     sync.RWMutex
	listeners []Listener
}

func (t *Broker) Send(m Msg) error {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	for _, listener := range t.listeners {
		listener(m)
	}
	return nil
}

func (t *Broker) Subscribe(cb func(msg Msg)) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.listeners = append(t.listeners, cb)
	return nil
}

type Msg struct {
	Topic   string
	Content string
}
