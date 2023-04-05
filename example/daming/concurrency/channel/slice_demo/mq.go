package map_demo

import (
	"errors"
	"sync"
)

type Broker struct {
	mutex sync.RWMutex
	chans []chan Msg
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

func (t *Broker) Subscribe(capacity int) (<-chan Msg, error) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	c := make(chan Msg, capacity)
	t.chans = append(t.chans, c)
	return c, nil
}

func (t *Broker) Close() error {
	t.mutex.Lock()
	chans := t.chans
	t.chans = nil
	t.mutex.Unlock()

	// 避免了重复close chan的问题
	for _, ch := range chans {
		close(ch)
	}
	return nil
}

type Msg struct {
	Topic   string
	Content string
}
