package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	status2 "google.golang.org/grpc/status"
	"log"
	pb "main/01_SimpleRpc/proto"

	"time"
)

var grpcClient pb.SimpleClient

func route(ctx context.Context, deadlines time.Duration){
	// 设置3秒的超时时间
	clientDeadLine := time.Now().Add(time.Duration(deadlines * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadLine)
	defer cancel()

	// 发送结构体
	req := pb.SimpleRequest{Data: "grpc"}

	// 传入超时时间为3秒的ctx
	res, err := grpcClient.Route(ctx, &req)
	if err != nil{
		// 获取错误状态
		statu, ok := status2.FromError(err)
		if ok{
			// 判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded{
				log.Fatalln("Route timeout")
			}
		}
		log.Fatalf("call route err :%v", err)
	}
	log.Println(res.Value)

}

func main()  {

	conn, err := grpc.Dial(":8083", grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("err %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	grpcClient = pb.NewSimpleClient(conn)
	route(ctx, 3)


}


/*
2021/11/29 00:48:27 Route timeout

*/