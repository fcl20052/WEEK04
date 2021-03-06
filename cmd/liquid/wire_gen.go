// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"liquid/internal/biz"
	"liquid/internal/conf"
	"liquid/internal/data"
	"liquid/internal/server"
	"liquid/internal/service"
)

// Injectors from wire.go:

func InitApp(c *conf.Data) (*App, error) {
	dataData, err := data.NewData(c)
	if err != nil {
		return nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo)
	greeterService := service.NewGreeterService(greeterUsecase)
	httpServer := server.NewHTTPServer(greeterService)
	grpcServer := server.NewGRPCServer(greeterService)
	app := newApp(httpServer, grpcServer)
	return app, nil
}
