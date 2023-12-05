package tcp

import (
	"context"
	"go-redis/interface/tcp"
	"go-redis/lib/logger"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Config 启动 Tcp Server 的结构体配置
type Config struct {
	Address string
}

// ListenAndServerWithSignal 监听服务，如果有异常返回信号
func ListenAndServerWithSignal(
	cfg *Config,
	handler tcp.Handler) error {

	closeChan := make(chan struct{})
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sig := <-sigChan
		switch sig {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeChan <- struct{}{}
		}
	}()
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
func ListenAndServer(
	listener net.Listener,
	handler tcp.Handler,
	// 防止用户在客户端未关闭就将进程杀死，感知系统信号
	closeChan <-chan struct{}) {

	go func() {
		// 实现对于系统系统的监听
		<-closeChan
		// 收到来自系统的信号
		logger.Info("shutting down")
		_ = listener.Close()
		_ = handler.Close()
	}()

	defer func() {
		// 在栈中对于 listener 和 handler 方法实现关闭逻辑
		_ = listener.Close()
		_ = handler.Close()
	}()
	// 从 background 方法获取到一个上下文对象
	ctx := context.Background()
	var waitDone sync.WaitGroup

	// 死循环接收连接错误
	for true {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		logger.Info("accepted link")
		go func() {
			defer func() {
				// 如果发生异常，会等待所有用户退出之后进行退出
				waitDone.Done()
			}()
			handler.Handler(ctx, conn)
		}()
	}
	waitDone.Wait()
}
