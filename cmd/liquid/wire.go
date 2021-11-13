//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"
	"liquid/internal/conf"
	"liquid/internal/biz"
	"liquid/internal/data"
	"liquid/internal/server"
	"liquid/internal/service"
)

func InitApp(c *conf.Data) (*App,error) {
	wire.Build(newApp, server.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet)
	return &App{},nil //返回值没有实际意义，只需符合函数签名即可
}
