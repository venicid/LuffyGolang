package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "main/03_ClientStreamRpc/proto"
	"strconv"
)



var streamClient pb.StreamClientClient
func main()  {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("net.Connect err:%v", err)
	}

	defer conn.Close()
	streamClient = pb.NewStreamClientClient(conn)
	routeList()
	route()
}




func routeList(){
	stream, err := streamClient.RouteList(context.Background())
	if err != nil{
		log.Fatalf("call listStr err:%v", err)
	}

	for n := 0; n < 5; n++ {
		// 向流发送消息
		err := stream.Send(&pb.StreamRequest{
			StreamValue: "stream cleint prc" + strconv.Itoa(n),
		})
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("stream request err:%v", err)
		}
	}

}


// 简单模式
func route(){
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data:"grpc",
	}

	// 调用我们的服务(Route方法)
	// 同时传入1个ConText,在需要时，可以让我们改变rpc的行为，比如超时/取消等
	res, err := streamClient.Route(context.Background(), &req)
	if err != nil{
		log.Fatalf("call route err:%v", err)
	}

	// 打印返回值
	log.Println(res)
}

/*


client端
2021/11/27 17:56:59 code:200  value:"hello grpc"

*/