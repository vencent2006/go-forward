/**
 * @Author: vincent
 * @Description:
 * @File:  event
 * @Version: 1.0.0
 * @Date: 2021/8/12 18:26
 */

package kim

import (
	"sync"
	"sync/atomic"
)

type Event struct {
	fired int32
	c     chan struct{}
	o     sync.Once
}

/*
第一次执行, 返回true，且，关闭c
之后执行, 返回false
*/
func (e *Event) Fire() bool {
	ret := false
	e.o.Do(func() {
		atomic.StoreInt32(&e.fired, 1)
		close(e.c)
		ret = true
	})
	return ret
}

func (e *Event) Done() <-chan struct{} {
	return e.c
}

// todo: Fire() HasFired() 还是不太清楚
func (e *Event) HasFired() bool {
	return atomic.LoadInt32(&e.fired) == 1
}

func NewEvent() *Event {
	return &Event{c: make(chan struct{})}
}
