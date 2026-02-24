package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	// 创建一个默认的cron对象
	c := cron.New()

	// 添加任务
	id, err := c.AddFunc("* * * * *", func() {
		fmt.Println("每分钟")
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = c.AddFunc("* * * * *", func() {
		fmt.Println("每分钟")
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	// 启动任务
	c.Start()

	// 获取任务
	ent := c.Entry(id)
	if !ent.Valid() {
		fmt.Println("not vaild")
		return
	}

	notice := time.Tick(time.Minute + time.Second*10)
	t := <-notice
	fmt.Println(t)

	// 移除任务
	c.Remove(id)

	// 停止全部任务
	ctx := c.Stop()
	allDone := ctx.Done()
	<-allDone
}
