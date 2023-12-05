package tcp

/**
 * A echo server to test whether the server is functioning normally
 * echo服务器，用于测试服务器是否正常运行
 */

import (
	"bufio"
	"context"
	"go-redis/lib/logger"
	"go-redis/lib/sync/atomic"
	"go-redis/lib/sync/wait"
	"io"
	"net"
	"sync"
	"time"
)

// EchoHandler echos received line to client, using for test
type EchoHandler struct {
	// 记录客户端连接
	activeConn sync.Map
	// 布尔型的 closing，使用原子的布尔
	closing atomic.Boolean
}

// MakeHandler MakeEchoHandler creates EchoHandler 创建EchoHandler
func MakeHandler() *EchoHandler {
	return &EchoHandler{}
}

// EchoClient is client for EchoHandler, using for test
// EchoHandler 客户端，即代表客户端信息，用于测试
type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

// Close connection
func (c *EchoClient) Close() error {
	c.Waiting.WaitWithTimeout(10 * time.Second)
	c.Conn.Close()
	return nil
}

// Handle echos received line to client
func (h *EchoHandler) Handler(ctx context.Context, conn net.Conn) {
	if h.closing.Get() {
		// closing handler refuse new connection
		_ = conn.Close()
	}

	client := &EchoClient{
		Conn: conn,
	}
	// 将客户端存储进去
	h.activeConn.Store(client, struct{}{})
	reader := bufio.NewReader(conn)

	// 死循环，接收用户发送过来的报文
	for {
		// may occur: client EOF, client timeout, server early close
		// 用 '\n' 作为标记位，记录用户发送过来的信息
		msg, err := reader.ReadString('\n')
		if err != nil {
			// 收到客户端发送的结束符号，表示连接结束
			if err == io.EOF {
				logger.Info("connection close")
				h.activeConn.Delete(client)
			} else {
				logger.Warn(err)
			}
			return
		}
		client.Waiting.Add(1)
		b := []byte(msg)
		_, _ = conn.Write(b)
		client.Waiting.Done()
	}
}

// Close stops echo handler 实现关闭方法
func (h *EchoHandler) Close() error {
	logger.Info("handler shutting down...")
	// 将业务引擎状态关闭
	h.closing.Set(true)
	h.activeConn.Range(func(key interface{}, val interface{}) bool {
		client := key.(*EchoClient)
		// 将用户返回的异常抛出，然后将客户端关闭
		_ = client.Close()
		return true
	})
	return nil
}
