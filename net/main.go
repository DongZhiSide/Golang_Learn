package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Listen() failed, err: ", err)
		os.Exit(1)
	}
	b := false
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() failed, err: ", err)
			continue
		}
		go func() {
			fmt.Println("Client connected: ", conn.RemoteAddr())
			time.Sleep(2 * time.Second)
			fmt.Println("write")
			conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("Content-type: text/html; charset=utf-8\r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("<h1>Hello, World!</h1><h1>Hello, World!</h1><h1>Hello, World!</h1><h1>Hello, World!</h1><h1>Hello, World!</h1>"))
			time.Sleep(2 * time.Second)
			if b {
				conn.Close()
				b = false
			}
			fmt.Println("close")
			b = true
		}()
	}
}
