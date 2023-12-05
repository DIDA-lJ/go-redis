package tcp

import (
	"context"
	"net"
)

// Handler /*  Handler 接口，用于处理 TCP 连接，可以使得在写 TCP 连接的时候忽视 TCP 连接
type Handler interface {
	Handler(ctx context.Context, conn net.Conn)
	Close() error
}
