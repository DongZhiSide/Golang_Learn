//go:build wireinject
// +build wireinject

package main

import (
	"log/slog"

	"github.com/google/wire"
)

func InitializeServer(logger *slog.Logger) *Server {
	wire.Build(NewRepositoryA, NewRepositoryB, NewServiceA, NewServiceB,
		wire.Struct(new(Server), "*"))
	return &Server{}
}
