/*
之前写一些代码时, 总是想要透传 logger,
用 slog.Default() 来获取或者是直接使用全局默认的 logger,
感觉又"不舒服", 会想到哪一天某个部分需要修改 logger 的配置,
但是其他部分又不需要修改 logger, 因此不免想到透传 logger 的方式.

但是透传 logger 的方式又感觉不是很好, 每次创建对象都要传递 logger,
要写很多重复的代码.

虽然 wire 可以解决这个问题, 但是感觉又不太好, 传来传去的没什么意义,
因为大部分情况下, 不会有修改 logger 的需求.

总的来说, 除非是某个部分的 logger 需要修改, 否则不建议透传 logger.

以下代码是一个使用 wire 透传 logger 的概念验证.
*/

package main

import (
	"log/slog"
)

type RepositoryA struct {
	logger *slog.Logger
}

func (r *RepositoryA) Start() {
	r.logger.Info("RepositoryA started")
}
func NewRepositoryA(logger *slog.Logger) *RepositoryA {
	return &RepositoryA{logger: logger}
}

type RepositoryB struct {
	logger *slog.Logger
}

func (r *RepositoryB) Start() {
	r.logger.Info("RepositoryB started")
}

func NewRepositoryB(logger *slog.Logger) *RepositoryB {
	return &RepositoryB{logger: logger}
}

type ServiceA struct {
	RepositoryA *RepositoryA
	logger      *slog.Logger
}

func (s *ServiceA) Start() {
	s.logger.Info("ServiceA started")
	s.RepositoryA.Start()
}

func NewServiceA(repositoryA *RepositoryA, logger *slog.Logger) *ServiceA {
	return &ServiceA{RepositoryA: repositoryA, logger: logger}
}

type ServiceB struct {
	RepositoryB *RepositoryB
	logger      *slog.Logger
}

func (s *ServiceB) Start() {
	s.logger.Info("ServiceB started")
	s.RepositoryB.Start()
}

func NewServiceB(repositoryB *RepositoryB, logger *slog.Logger) *ServiceB {
	return &ServiceB{RepositoryB: repositoryB, logger: logger}
}

type Server struct {
	ServiceA *ServiceA
	ServiceB *ServiceB
	logger   *slog.Logger
}

func (s *Server) Start() {
	s.logger.Info("Server started")
	s.ServiceA.Start()
	s.ServiceB.Start()
}

func main() {
	logger := slog.Default()
	server := InitializeServer(logger)
	server.Start()
}
