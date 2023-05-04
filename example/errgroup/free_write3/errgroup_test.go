package errgroup

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_group(t *testing.T) {
	testCases := []struct {
		name    string
		exec    func() error
		wantErr error
	}{
		{
			name: "normal",
			exec: func() error {
				var g Group
				g.Go(func() error {
					fmt.Printf("number:%d\n", 1)
					return nil
				})
				g.Go(func() error {
					fmt.Printf("number:%d\n", 2)
					return nil
				})
				return g.Wait()
			},
			wantErr: nil,
		},
		{
			name: "set limit 1",
			exec: func() error {
				var g Group
				g.SetLimit(1)
				g.Go(func() error {
					fmt.Printf("number:%d\n", 1)
					return nil
				})
				g.Go(func() error {
					fmt.Printf("number:%d\n", 2)
					return nil
				})
				return g.Wait()
			},
			wantErr: nil,
		},
		{
			name: "with cancel",
			exec: func() error {
				g, ctx := WithContext(context.Background())
				g.Go(func() error {
				LOOP1:
					for {
						select {
						case <-ctx.Done():
							break LOOP1
						default:
							fmt.Printf("number:%d running\n", 1)
							time.Sleep(time.Second)
						}
					}
					fmt.Printf("number:%d finished\n", 1)
					return nil
				})
				g.Go(func() error {
					var i int
				LOOP2:
					for {
						select {
						case <-ctx.Done():
							break LOOP2
						default:
							fmt.Printf("number:%d running\n", 2)
							time.Sleep(time.Second)
							i++
							if i > 3 {
								return fmt.Errorf("timeout err happened")
							}
						}
					}
					fmt.Printf("number:%d finished\n", 2)
					return nil
				})
				return g.Wait()
			},
			wantErr: fmt.Errorf("timeout err happened"),
		},
		//{
		//	name: "set limit 0", // 这个会阻塞住, 但却不会检测为死锁，比较奇怪；n=0，应该是无法使用的
		//	exec: func() error {
		//		var g Group
		//		g.SetLimit(0)
		//		errChan := make(chan error, 10)
		//
		//		fmt.Println("prepare to block")
		//
		//		select {
		//		//case errChan <- goWait(1):
		//		//case errChan <- goWait(2):
		//		case <-time.After(time.Second * 1):
		//			fmt.Println("time over, g.Go not return")
		//		case errChan <- func(i int) error {
		//			g.Go(func() error {
		//				fmt.Printf("number:%d\n", i)
		//				return nil
		//			})
		//			return nil
		//		}(1):
		//			fmt.Println("i wait")
		//		}
		//
		//		return g.Wait()
		//	},
		//	wantErr: nil,
		//},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.exec()
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

func Test_limit_myGroup(t *testing.T) {
	var eg Group
	eg.SetLimit(0)
	ch := make(chan int, 1)
	select {
	case ch <- func() int {
		eg.Go(func() error {
			return nil
		})
		return 1
	}():
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
	//eg.Go(func() error {
	//	fmt.Println("finish job")
	//	return nil
	//})
	err := eg.Wait()
	t.Log(err)
}

func Test_multiCancel(t *testing.T) {

	//context控制并发  WithCancel  cancel()
	ctx, cancel := context.WithCancel(context.Background())
	go Worker(ctx, "01")
	go Worker(ctx, "02")
	go Worker(ctx, "03")
	time.Sleep(time.Second * 5)
	fmt.Println("stop all")
	cancel()
	time.Sleep(time.Second * 1)
	cancel()
	cancel()
	cancel()
	cancel()
	cancel() // cancel()是幂等的，因为一旦err设置过，就不会传播cancel调用了
}

func Worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " stop the goroutine!")
			return
		default:
			fmt.Println(name, " ing")
			time.Sleep(time.Second * 1)
		}
	}
}
