package main

import (
	"errors"
	"fmt"
	"log/slog"
)

// Message 打招呼的消息
type Message string

// Greeter 迎宾员, 持有 Message
type Greeter struct {
	Message Message
}

// Event 活动, 持有 Greeter
type Event struct {
	Greeter Greeter
}

// NewMessage 提供 Message 的 Provider
func NewMessage(phrase string) Message {
	return Message(phrase)
}

// NewGreeter 提供 Greeter 的 Provider, 依赖 Message
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent 提供 Event 的 Provider, 依赖 Greeter
func NewEvent(g Greeter) (Event, error) {
	if g.Message == "" {
		return Event{}, errors.New("Message is empty")
	}
	return Event{Greeter: g}, nil
}

// Start 启动活动, 打印招呼语
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// Greet 迎宾员发出招呼
func (g Greeter) Greet() Message {
	return g.Message
}

// 依赖关系: Event -> Greeter -> Message
// 依赖的 provider 是 NewEvent, NewGreeter, NewMessage
func main() {
	e, err := InitializeEvent("hello wire!")
	if err != nil {
		slog.Error("Failed to initialize event", slog.Any("error", err))
		return
	}
	e.Start()

	ep, err := InitializeEventP("hello wire!")
	if err != nil {
		slog.Error("Failed to initialize event", slog.Any("error", err))
		return
	}
	ep.Start()
}
