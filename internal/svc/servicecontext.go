package svc

import (
	"github.com/lokichoggio/add/adder"
	"github.com/lokichoggio/check/checker"
	"github.com/lokichoggio/gateway/internal/config"

	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder     // 手动代码
	Checker checker.Checker // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),       // 手动代码
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)), // 手动代码
	}
}
