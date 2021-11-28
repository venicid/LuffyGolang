package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "main/01_SimpleRpc/proto"
	"net"
)

// 定义我们的服务，并且实现Route方法
type SimpleServer struct {
}

// Route实现Route方法
func (s *SimpleServer) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}

	return &res, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("net listen err: %v", err)
	}
	log.Println(":8080 net.Listening...")

	// 新建gRpc服务器实例
	grpcServer := grpc.NewServer()
	// 在grpc服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleServer{})
	// 需要用Server()方法以及我们的端口信息，去阻塞等待，直到进程被杀死
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("grpcServer.server err : %v", err)
	}
}
