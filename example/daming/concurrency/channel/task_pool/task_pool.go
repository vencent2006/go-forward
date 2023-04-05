package task_pool

import "context"

type Task func()

type TaskPool struct {
	tasks chan Task
	//close *atomic.Bool
	close chan struct{}
}

// numG 是goroutine的数量，就是你要控制住的
// capacity 是缓存的容量
func NewTaskPool(numG int, capacity int) *TaskPool {
	res := &TaskPool{
		tasks: make(chan Task, capacity),
		//close: atomic.NewBool(false),
		close: make(chan struct{}),
	}

	// 这个东西，要是没有退出 goroutine的机制，那就是妥妥的 goroutine 泄露
	for i := 0; i < numG; i++ {
		go func() {
			for {
				select {
				case <-res.close:
					return
				case t := <-res.tasks:
					t()
				}
			}
			//for t := range res.tasks {
			//	if res.close.Load() {
			//		return
			//	}
			//	t()
			//}
		}()
	}
	return res
}

// 提交任务
func (t *TaskPool) Submit(ctx context.Context, task Task) error {
	select {
	case t.tasks <- task:
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

// 会释放资源，重复调用panic
func (t *TaskPool) Close() error {
	//t.close.Store(true)
	//这种方法不行
	//t.close <- struct {}{}
	close(t.close) // 重复调用会panic
	return nil
}
