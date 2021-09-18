/**
 * @Author: vincent
 * @Description:
 * @File:  channelImpl
 * @Version: 1.0.0
 * @Date: 2021/8/12 18:07
 */

package kim

import (
	"errors"
	"fmt"
	"go-examples/demo/chat_demo_im/logger"
	"sync"
	"time"
)

type ChannelImpl struct {
	// 保护性措施
	sync.Mutex           // 防止ReadLoop多次重入
	once       sync.Once // 保证Close只运行1次

	// attr
	id string
	Conn

	// action data struct
	readWait  time.Duration
	writeWait time.Duration
	writeChan chan []byte
	closed    *Event
}

func NewChannel(id string, conn Conn) Channel {
	// log withFields
	log := logger.WithFields(logger.Fields{
		"module": "channel",
		"id":     id,
	})

	// 填充变量
	ch := &ChannelImpl{
		id:        id,
		Conn:      conn,
		readWait:  DefaultReadWait, // default value
		writeWait: DefaultReadWait, // default value
		writeChan: make(chan []byte, 5),
		closed:    NewEvent(),
	}

	//go writeLoop
	go func() {
		err := ch.writeLoop()
		if err != nil {
			log.Error(err)
		}
	}()

	return ch
}

// todo: 为什么writeLoop不放在interface Channel
func (ch *ChannelImpl) writeLoop() error {
	for {
		select {
		// write payload
		case payload := <-ch.writeChan:
			err := ch.WriteFrame(OpBinary, payload)
			if err != nil {
				return err
			}

			// todo: 这是为了再多write点吗，这是合理的姿势吗
			chanLen := len(ch.writeChan)
			for i := 0; i < chanLen; i++ {
				payload = <-ch.writeChan
				err := ch.WriteFrame(OpBinary, payload)
				if err != nil {
					return err
				}
			}

			// todo: Flush的必要性, ch.Flush是干啥用的呢
			err = ch.Conn.Flush()
			if err != nil {
				return err
			}
		case <-ch.closed.Done():
			return nil
		}
	}

}

func (ch *ChannelImpl) ID() string {
	return ch.id
}

// 这是实现interface Agent
func (ch *ChannelImpl) Push(payload []byte) error {
	if ch.closed.HasFired() {
		return fmt.Errorf("channel %s has closed", ch.id)
	}
	// 异步写
	ch.writeChan <- payload
	return nil
}

func (ch *ChannelImpl) ReadLoop(lis MessageListener) error {
	// 毕竟是暴露方法, 确保不会并发进入
	ch.Lock()
	defer ch.Unlock()

	// add log field
	log := logger.WithFields(logger.Fields{
		"struct": "ChannelImpl",
		"func":   "ReadLoop",
		"id":     ch.ID(),
	})

	// loop
	for {
		// todo: 这个deadline设置的比较啰嗦，
		// 不如仿照WriteFrame在ReadFrame里添加deadline设置
		_ = ch.SetReadDeadline(time.Now().Add(ch.readWait))
		frame, err := ch.ReadFrame()
		if err != nil {
			return err
		}

		// todo: 用switch会更好
		if frame.GetOpCode() == OpClose {
			return errors.New("remote side close the channel")
		}

		if frame.GetOpCode() == OpPing {
			log.Info("recv a ping; resp with a pong")
			_ = ch.WriteFrame(OpPong, nil)
			continue
		}

		//真正的消息
		payload := frame.GetPayload()
		if len(payload) == 0 {
			// 长度为0"
			log.Info("recv a msg; but payload len is 0")
			continue
		}

		// todo: Optimization point
		go lis.Receive(ch, payload)

	}
}

func (ch *ChannelImpl) SetWriteWait(duration time.Duration) {
	if duration == 0 {
		return
	}
	ch.writeWait = duration
}

func (ch *ChannelImpl) SetReadWait(duration time.Duration) {
	if duration == 0 {
		return
	}
	ch.readWait = duration
}

// overwrite Conn: WriteFrame 主要是多了Deadline
func (ch *ChannelImpl) WriteFrame(code OpCode, payload []byte) error {
	_ = ch.Conn.SetWriteDeadline(time.Now().Add(ch.writeWait))
	return ch.Conn.WriteFrame(code, payload)
}

func (ch *ChannelImpl) Close() error {
	// 因为有chan，所以只能关闭一次
	ch.once.Do(func() {
		close(ch.writeChan)
		ch.closed.Fire()
	})
	return nil
}
