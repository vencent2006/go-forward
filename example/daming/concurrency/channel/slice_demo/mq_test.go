package map_demo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBroker_Send(t *testing.T) {
	// producer
	b := &Broker{}
	go func() {
		for {
			err := b.Send(Msg{Content: time.Now().String()})
			if err != nil {
				t.Error(err)
				return
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// consumer
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		topic := fmt.Sprintf("消费者 %d", i)
		go func() {
			defer wg.Done()
			msgs, err := b.Subscribe(100)
			if err != nil {
				t.Error(err)
				return
			}
			for msg := range msgs {
				t.Log(topic, msg.Content)
			}
		}()
	}
	wg.Wait()

}
