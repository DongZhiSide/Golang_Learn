package main

import (
	"context"
	"fmt"
	"golanglearn/writing-style/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	server.Entry(engine)

	mainServer := &http.Server{Addr: ":8080", Handler: engine}

	go func() {
		err := mainServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// 信号注册
	sigs := make(chan os.Signal, 1)
	// 注册需要监听的信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("等待信号, Ctrl+C 或 kill ...")
	// 阻塞等待信号
	sig := <-sigs
	fmt.Println("收到信号:", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // 处理 操作完成先于超时就需要"上下文取消" 的情况

	// 关闭
	err := mainServer.Shutdown(ctx)
	if err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}
}
