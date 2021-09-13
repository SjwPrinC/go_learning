package main

import (
	"context"
	"fmt"

	sjw_myrpc1 "github.com/sjw/grpctest/sjw.myrpc1"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:18888", grpc.WithInsecure())

	if err != nil {
		fmt.Println("connect failed")
		return
	}

	//创建grpc 客户端，就好比wcf的clientChannel
	c := sjw_myrpc1.NewGreetServiceClient(conn)

	re, err := c.Hello(context.Background(), &sjw_myrpc1.HelloRequest{Message: "winnie"})

	if err != nil {
		fmt.Println("call failed:", err.Error())
		return
	}

	fmt.Println("receive message: ", re.Reply)
}
