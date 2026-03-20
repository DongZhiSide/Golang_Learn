//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// InitializeEvent 注入器: 声明需要的返回类型, Wire 会生成依赖注入代码
// phrase 作为参数传入, 会传递给 NewMessage
func InitializeEvent(phrase string) Event {
	wire.Build(NewMessage, NewGreeter, NewEvent)
	return Event{}
}
