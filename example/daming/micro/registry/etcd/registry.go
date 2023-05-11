package etcd

import (
	"context"
	"encoding/json"
	"example/daming/micro/registry"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type Registry struct {
	sess *concurrency.Session
	c    *clientv3.Client
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
	//TODO implement me
	panic("implement me")
}

func (r *Registry) Subscribe(serviceName string) (<-chan registry.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Registry) Close() error {
	err := r.sess.Close() // 调用完session的Close，不用调用UnRegister
	return err
}

func (r *Registry) instanceKey(si registry.ServiceInstance) string {
	// 可以考虑直接使用 Address
	// 也可以说在 si 里面引入一个 InstanceName 的字段
	return fmt.Sprintf("/micro/%s/%s", si.Name, si.Address)
}

func (r *Registry) serviceKey(si registry.ServiceInstance) string {
	return fmt.Sprintf("/micro/%s", si.Name)
}
