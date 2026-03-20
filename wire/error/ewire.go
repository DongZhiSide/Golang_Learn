//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// InitializeEvent 注入器: 声明需要的返回类型, Wire 会生成依赖注入代码
// phrase 作为参数传入, 会传递给 NewMessage
func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewMessage, NewGreeter, NewEvent)
	return Event{}, nil
}

// wire 为我们提供了 provider sets, 顾名思义,
// 它可以包含一组 providers.
// 使用 wire.NewSet 函数可以将多个 provider 添加到一个集合中.
var providerSet wire.ProviderSet = wire.NewSet(NewMessage, NewGreeter)

func InitializeEventP(phrase string) (Event, error) {
	wire.Build(NewEvent, providerSet)
	return Event{}, nil
}
