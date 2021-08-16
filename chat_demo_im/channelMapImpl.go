/**
 * @Author: vincent
 * @Description:
 * @File:  channelMapImpl
 * @Version: 1.0.0
 * @Date: 2021/8/13 09:52
 */

package kim

import (
	"go-examples/chat_demo_im/logger"
	"sync"
)

const (
	MODULE_CHANNELMAP_IMPL = "ChannelMapImpl"
)

type ChannelMapImpl struct {
	// todo: Optimization point
	channels *sync.Map
}

// Add
func (ch *ChannelMapImpl) Add(channel Channel) {
	// check
	if channel.ID() == "" {
		logger.WithFields(logger.Fields{
			"module": MODULE_CHANNELMAP_IMPL,
		}).Error("channel id is required")
	}

	// do
	ch.channels.Store(channel.ID(), channel)
}

// Remove
func (ch *ChannelMapImpl) Remove(channelId string) {
	ch.channels.Delete(channelId)
}

// Get
func (ch *ChannelMapImpl) Get(channelId string) (Channel, bool) {
	// check
	if channelId == "" {
		logger.WithFields(logger.Fields{
			"module": MODULE_CHANNELMAP_IMPL,
		}).Error("channel id is required")
	}

	// do
	val, ok := ch.channels.Load(channelId)
	if !ok {
		// 不存在
		return nil, false
	}
	return val.(Channel), true
}

// All
func (ch *ChannelMapImpl) All() []Channel {
	arr := make([]Channel, 0)
	ch.channels.Range(func(key, value interface{}) bool {
		arr = append(arr, value.(Channel))
		return true
	})
	return arr
}

func NewChannelMap(num int) ChannelMap {
	// todo: 这个num没有被使用，感觉应该是用来池化pool
	return &ChannelMapImpl{channels: new(sync.Map)}
}
