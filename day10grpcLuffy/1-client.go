package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "main/01_SimpleRpc/proto"
)

func main()  {

	// 连接服务器
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("net.Connect err:%v", err)
	}

	defer conn.Close()

	// 建立grpc连接
	grpcClient := pb.NewSimpleClient(conn)
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data:"grpc",
	}

	// 调用我们的服务(Route方法)
	// 同时传入1个ConText,在需要时，可以让我们改变rpc的行为，比如超时/取消等
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil{
		log.Fatalf("call route err:%v", err)
	}
	log.Println(res)

}