package main

import (
	"context"
	"fmt"
	"net"

	sjw_myrpc1 "github.com/sjw/grpctest/sjw.myrpc1"
	"google.golang.org/grpc"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:18888")

	if err != nil {
		return
	}

	srv := grpc.NewServer()

	sjw_myrpc1.RegisterGreetServiceServer(srv, &greetServer{})

	fmt.Println("server run")
	srv.Serve(ln)
}

func init() {

}

type greetServer struct {
}

//实现grpc hello接口
func (gs *greetServer) Hello(ctx context.Context, in *sjw_myrpc1.HelloRequest) (*sjw_myrpc1.HelloReply, error) {
	fmt.Println("receive request:", in.Message)
	return &sjw_myrpc1.HelloReply{Reply: "hello :" + in.Message}, nil
}
