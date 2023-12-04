package tcp

import (
	"context"
	"go-redis/interface/tcp"
	"go-redis/lib/logger"
	"net"
)

// Config 启动 Tcp Server 的结构体配置
type Config struct {
	Address string
}

// ListenAndServerWithSignal 监听服务，如果有异常返回信号
func ListenAndServerWithSignal(cfg *Config,
	handler tcp.Handler) error {
	closeChan := make(chan struct{})
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info("start listen")

	// 将 listener 和 handler 传入  ListenAndServer,实现监听并且连接
	ListenAndServer(listener, handler, closeChan)

	return nil
}

// ListenAndServer  实现监听并且连接
func ListenAndServer(listener net.Listener,
	handler tcp.Handler,
	closeChan <-chan struct{}) {

	// 从 background 方法获取到一个上下文对象
	ctx := context.Background()

	for true {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		logger.Info("accepted link")
		go func() {
			handler.Handler(ctx, conn)
		}()
	}

}
