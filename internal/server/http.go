package server

import (
	"net/http"
	"liquid/internal/service"
)

type HTTPServer struct {
	hs http.Server
}

func NewHTTPServer(*service.GreeterService) *HTTPServer {
	srv := &HTTPServer{}
	srv.hs = http.Server{}
	
	//注册服务
	return srv
}
