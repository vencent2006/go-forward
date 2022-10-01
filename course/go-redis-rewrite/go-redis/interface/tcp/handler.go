/**
 * @Author: vincent
 * @Description:
 * @File:  handler
 * @Version: 1.0.0
 * @Date: 2022/9/30 14:59
 */

package tcp

import (
	"context"
	"net"
)

// Handler 业务逻辑处理的接口
type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close() error
}
