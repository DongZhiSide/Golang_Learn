package main

import (
	"fmt"
	"time"
)

func main() {
	urls := []string{"a", "b", "c"}
	for _, v := range urls {
		go func() {
			// 捕获了同一个 v（range 的循环变量会被复用）
			fmt.Println(v)
		}()
	}
	time.Sleep(100 * time.Millisecond)
}
