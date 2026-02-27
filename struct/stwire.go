//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// InitializeEvent 注入器: 声明需要的返回类型, Wire 会生成依赖注入代码
// phrase 作为参数传入, 会传递给 NewMessage
func InitializeEvent(phrase string, code int) (Event, error) {
	// "Phrase" 是 Message 结构体中的字段名
	// "*" 表示注入所有字段
	wire.Build(NewEvent, NewGreeter, wire.Struct(new(Message), "Phrase"))
	return Event{}, nil
}
