package main

import (
	"flag"
	"os/signal"
	"syscall"
	"context"
	"go-kit-service/lorem-service"
	"net"
	"google.golang.org/grpc"
	"go-kit-service/lorem-service/proto"
	"os"
	"fmt"
)

func main() {
	var (
		gRPCAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// init lorem service
	var svc lorem_service.Service
	svc = lorem_service.LoremService{}
	errChan := make(chan error)

	// creating Endpoints struct
	endpoints := lorem_service.Endpoints{
		LoremEndpoint: lorem_service.MakeLoremEndpoint(svc),
	}

	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}
		handler := lorem_service.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterLoremServer(gRPCServer, handler)
		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()
	fmt.Println("server start...")
	fmt.Println(<-errChan)
}
