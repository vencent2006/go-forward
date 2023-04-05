package map_demo

import (
	"errors"
	"sync"
)

type Broker struct {
	mutex sync.RWMutex
	chans map[string]chan Msg
}

func (t *Broker) Send(m Msg) error {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	for _, ch := range t.chans {
		select {
		case ch <- m:
		default:
			return errors.New("消息队列已满")
		}
	}
	return nil
}

func (t *Broker) Subscribe(topic string, capacity int) (<-chan Msg, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	if t.chans == nil {
		t.chans = make(map[string]chan Msg)
	}
	ch, isHas := t.chans[topic]
	if isHas {
		return ch, nil
	}
	c := make(chan Msg, capacity)
	t.chans[topic] = c
	return c, nil
}

type Msg struct {
	Topic   string
	Content string
}
