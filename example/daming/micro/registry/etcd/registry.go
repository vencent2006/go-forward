package etcd

import (
	"context"
	"encoding/json"
	"example/daming/micro/registry"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"sync"
)

type Registry struct {
	c       *clientv3.Client
	sess    *concurrency.Session
	cancels []func()
	mutex   sync.Mutex // 也可以换成读写锁
}

func NewRegistry(c *clientv3.Client) (*Registry, error) {
	sess, err := concurrency.NewSession(c)
	if err != nil {
		return nil, err
	}
	return &Registry{
		sess: sess,
		c:    c,
	}, nil
}

func (r *Registry) Register(ctx context.Context, si registry.ServiceInstance) error {
	val, err := json.Marshal(si)
	if err != nil {
		return err
	}
	_, err = r.c.Put(ctx, r.instanceKey(si), string(val), clientv3.WithLease(r.sess.Lease()))
	return err
}

func (r *Registry) UnRegister(ctx context.Context, si registry.ServiceInstance) error {
	_, err := r.c.Delete(ctx, r.instanceKey(si))
	return err
}

func (r *Registry) ListServices(ctx context.Context, serviceName string) ([]registry.ServiceInstance, error) {
	getResp, err := r.c.Get(ctx, r.serviceKey(serviceName), clientv3.WithPrefix()) // 取前缀
	if err != nil {
		return nil, err
	}
	res := make([]registry.ServiceInstance, 0, len(getResp.Kvs))
	for _, kv := range getResp.Kvs {
		var si registry.ServiceInstance
		err = json.Unmarshal(kv.Value, &si)
		if err != nil {
			return nil, err
		}
		res = append(res, si)
	}
	return res, nil
}

func (r *Registry) Subscribe(serviceName string) (<-chan registry.Event, error) {
	ctx, cancel := context.WithCancel(context.Background())
	r.mutex.Lock()
	r.cancels = append(r.cancels, cancel)
	r.mutex.Unlock()
	ctx = clientv3.WithRequireLeader(ctx) // 表示从etcd的leader，避免因为选举时，造成的数据不一致
	watchResp := r.c.Watch(ctx, r.serviceKey(serviceName), clientv3.WithPrefix())
	res := make(chan registry.Event)
	go func() {
		for {
			select {
			case resp := <-watchResp:
				if resp.Err() != nil {
					// return 老师说，这块的continue和Err的区别不大
					continue
				}
				if resp.Canceled {
					return
				}

				for range resp.Events { // 还能 for range 连着写呢
					res <- registry.Event{}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return res, nil
}

func (r *Registry) Close() error {
	r.mutex.Lock()
	cancels := r.cancels
	r.cancels = nil
	r.mutex.Unlock()
	for _, cancel := range cancels {
		cancel()
	}
	err := r.sess.Close() // 调用完session的Close，不用调用UnRegister
	return err
}

func (r *Registry) instanceKey(si registry.ServiceInstance) string {
	// 可以考虑直接使用 Address
	// 也可以说在 si 里面引入一个 InstanceName 的字段
	return fmt.Sprintf("/micro/%s/%s", si.Name, si.Address)
}

func (r *Registry) serviceKey(serviceName string) string {
	return fmt.Sprintf("/micro/%s", serviceName)
}
