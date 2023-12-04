package tcp

import (
	"context"
	"net"
)

// Handler /*  Handler 接口
type Handler interface {
	Handler(ctx context.Context, conn net.Conn)
	Close() error
}
