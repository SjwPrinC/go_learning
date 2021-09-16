package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sjw/consul/prodService"
	"google.golang.org/grpc"
)

func main() {
	NewMyProdService()

	ln, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("listen failed:", err.Error())
		return
	}

	server := grpc.NewServer()
	prodService.RegisterProductServiceServer(server, &MyProdService{})

	fmt.Println("server is running")
	server.Serve(ln)
}

type MyProdService struct {
}

func (mp *MyProdService) SayHello(ctx context.Context, request *prodService.HelloReq) (*prodService.HelloRly, error) {
	return &prodService.HelloRly{Greet: "hello: " + request.Name}, nil
}

func NewMyProdService() *MyProdService {
	return &MyProdService{}
}
