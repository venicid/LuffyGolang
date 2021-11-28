package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "main/03_ClientStreamRpc/proto"
	"net"
)

type SimpleService struct {

}

// Route实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}

	return &res, nil
}



func (s *SimpleService) RouteList(srv pb.StreamClient_RouteListServer) error {

	for{
		// 从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF{
			// 发送结果，并关闭
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
		}
		if err != nil {
			return err
		}

		log.Println(res.StreamValue)
	}
}



func main(){
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("net listen err: %v", err)
	}
	log.Println(":8080 net.Listening...")

	// 新建gRpc服务器实例
	grpcServer := grpc.NewServer()

	// 在grpc服务器注册我们的服务
	pb.RegisterStreamClientServer(grpcServer, &SimpleService{})

	// 需要用Server()方法以及我们的端口信息，去阻塞等待，直到进程被杀死
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("grpcServer.server err : %v", err)
	}
}


/*

server端
2021/11/27 17:46:39 :8080 net.Listening...
2021/11/27 17:56:59 stream cleint prc0
2021/11/27 17:56:59 stream cleint prc1
2021/11/27 17:56:59 stream cleint prc2
2021/11/27 17:56:59 stream cleint prc3
2021/11/27 17:56:59 stream cleint prc4

*/