package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "main/02_ServerStreamRpc/proto"
)

// 服务端流式rpc
func listValue(){
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "stream server grpc",
	}
	// 调用我们的服务 ListValue方法
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("net connect err: %v", err)
	}

	defer conn.Close()

	grpcClient := pb.NewStreamServerClient(conn)
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil{
		log.Fatalf("call listStr err:%v", err)
	}

	for {
		// Recv方法接受服务端消息，默认每次Recv()最大消息长度为1024*1024*4 bytes （4m）
		res, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatalf("listen get stream err :%v", err)
		}

		// 打印返回值
		log.Println(res.StreamValue)
	}
}


// 简单模式
func route(){

	// 连接服务器
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("net.Connect err:%v", err)
	}

	defer conn.Close()

	// 建立grpc连接
	grpcClient := pb.NewStreamServerClient(conn)
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

	// 打印返回值
	log.Println(res)


}

func main()  {
	listValue()
	route()
}

/*

2021/11/26 09:24:11 stream server grpc0
2021/11/26 09:24:11 stream server grpc1
2021/11/26 09:24:11 stream server grpc2
2021/11/26 09:24:11 stream server grpc3
2021/11/26 09:24:11 stream server grpc4
2021/11/26 09:24:11 code:200  value:"hello grpc"

*/