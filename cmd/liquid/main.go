package main

import (
	"context"
	"liquid/internal/conf"
	"liquid/internal/server"

	"os"
	"sync"

	"golang.org/x/sync/errgroup"
)

type App struct {
	hs *server.HTTPServer
	gs *server.GRPCServer
}

func newApp(hs *server.HTTPServer, gs *server.GRPCServer) *App {
	return &App{
		hs,
		gs,
	}
}

func main() {
	app, err := InitApp(&conf.Data{})

	if err != nil {
		return
	}
	if err := app.Run(); err != nil {
		panic(err)
	}

}

func (a *App) Run() error {
	ctx := context.Background()
	eg, ctx := errgroup.WithContext(ctx)
	wg := sync.WaitGroup{}

	//启动服务
	eg.Go(func() error {
		<-ctx.Done() //等待关闭信号
		return nil   //关闭服务
	})
	wg.Add(1)
	eg.Go(func() error {
		wg.Done()
		return nil //启动服务
	})

	wg.Wait()
	//注册信号
	c := make(chan os.Signal, 1)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c:
			//接受系统信号，关闭服务
			return nil
		}
	})
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
