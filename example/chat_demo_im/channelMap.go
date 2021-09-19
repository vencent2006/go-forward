/**
 * @Author: vincent
 * @Description:
 * @File:  channelMap
 * @Version: 1.0.0
 * @Date: 2021/8/13 09:34
 */

package kim

type ChannelMap interface {
	Add(channel Channel)
	Remove(channelId string)
	Get(channelId string) (channel Channel, ok bool)
	All() []Channel
}
