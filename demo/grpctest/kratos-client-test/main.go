package main

import (
	"context"
	"fmt"

	v1 "github.com/sjw/grpctest/kratos-client-test/proto/v1"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("tcp", "127.0.0.1")

	if err != nil {
		fmt.Println("connection failerd: ", err.Error())
		return
	}

	cli, err := v1.NewGreeterClient(conn)

	re, err := cli.SayHello(context.Background(), &v1.HelloRequest{Name: "jerry"})

	fmt.Println("receive message : ", re.Reply)
}
