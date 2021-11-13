package server

import (
	"google.golang.org/grpc"
	"liquid/internal/service"
)

type GRPCServer struct {
	Server *grpc.Server
}

func NewGRPCServer(*service.GreeterService) *GRPCServer {
	srv := &GRPCServer{}
	srv.Server = grpc.NewServer()
	
	//注册服务
	return srv
}
