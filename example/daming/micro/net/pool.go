package net

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"
)

type Pool struct {
	// 空闲连接队列
	idlesConns chan *idleConn
	// 请求(阻塞)队列
	reqQueue []connReq

	// 最大连接数
	maxCnt int
	// 当前连接数——你已经建好了的
	cnt int
	// 最大空闲时间
	maxIdleTime time.Duration
	// 工厂方法
	factory func() (net.Conn, error)

	// 锁
	lock sync.Mutex
}

func NewPool(initCnt int, maxIdleCnt int, maxCnt int, maxIdleTime time.Duration, factor func() (net.Conn, error)) (*Pool, error) {
	if initCnt > maxIdleCnt {
		return nil, errors.New("micro: 初始连接数量不能大于最大空闲连接数量")
	}

	idlesConns := make(chan *idleConn, maxIdleCnt)
	for i := 0; i < initCnt; i++ {
		conn, err := factor()
		if err != nil {
			return nil, err
		}
		idlesConns <- &idleConn{c: conn, lastActiveTime: time.Now()}
	}
	res := &Pool{
		idlesConns:  idlesConns,
		maxCnt:      maxCnt,
		cnt:         0,
		maxIdleTime: maxIdleTime,
		factory:     factor,
	}
	return res, nil
}

func (p *Pool) Get(ctx context.Context) (net.Conn, error) {
	// 检查是否ctx超时
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	for {
		select {
		case ic := <-p.idlesConns:
			// 有空闲连接
			if ic.lastActiveTime.Add(p.maxIdleTime).Before(time.Now()) {
				// 超过最大空闲时间
				_ = ic.c.Close()
				continue // 进入下一次循环
			}
			return ic.c, nil
		default:
			// 没有空闲连接
			p.lock.Lock()
			if p.cnt >= p.maxCnt {
				// 超过上限了
				req := connReq{connChan: make(chan net.Conn, 1)}
				p.reqQueue = append(p.reqQueue, req)
				p.lock.Unlock() // 注意解锁的位置
				select {
				case <-ctx.Done():
					// 选项1：从队列里面删除req自己; 我们的逻辑不支持删除自己
					// 选项2：在这里转发
					go func() {
						c := <-req.connChan
						_ = p.Put(context.Background(), c)
					}()
					return nil, ctx.Err()
				case c := <-req.connChan:
					return c, nil
				}
			} else {
				c, err := p.factory()
				if err != nil {
					return nil, err
				}
				p.cnt++ // 计数加1
				p.lock.Unlock()
				return c, nil
			}
		}
	}
}

func (p *Pool) Put(ctx context.Context, c net.Conn) error {
	p.lock.Lock()
	if len(p.reqQueue) > 0 {
		// 有阻塞的请求
		req := p.reqQueue[0] // 如果从队尾取，效率更高
		p.reqQueue = p.reqQueue[1:]
		p.lock.Unlock() // 解锁的位置很重要
		req.connChan <- c
		return nil
	} else {
		// 没有阻塞的请求
		defer p.lock.Unlock()
		ic := &idleConn{c: c, lastActiveTime: time.Now()}
		select {
		case p.idlesConns <- ic:
		default:
			// 空闲队列满了
			_ = c.Close()
			p.cnt--
		}
		return nil
	}
}

type idleConn struct {
	c              net.Conn
	lastActiveTime time.Time
}

type connReq struct {
	connChan chan net.Conn
}
