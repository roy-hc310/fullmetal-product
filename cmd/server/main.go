package main

import (
	"context"
	"fmt"
	"net"
	"os/signal"
	"syscall"

	"github.com/cloudwego/kitex/server"
	"github.com/roy-hc310/fullmetal-product/internal/bootstrap"
	"github.com/roy-hc310/fullmetal-product/pkg/config"
)

func main() {
	err := config.LoadGlobalEnv(".")
	if err != nil {
		fmt.Printf("Failed to load env: %v\n", err)
		return
	}

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("localhost:%s", config.GlobalEnv.RPCPort))
	if err != nil {
		fmt.Printf("Failed to resolve tcp addr: %v\n", err)
		return
	}

	svr := server.NewServer(server.WithServiceAddr(addr))

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	application, err := bootstrap.NewApplication(ctx, &config.GlobalEnv, &svr)
	if err != nil {
		fmt.Printf("Failed to bootstrap application: %v\n", err)
		return
	}

	go func() {
		err = svr.Run()
		if err != nil {
			fmt.Printf("Failed to run server: %v\n", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Shutting down server...")
	application.Shutdown(context.Background())
}
