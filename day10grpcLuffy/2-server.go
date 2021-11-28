package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "main/02_ServerStreamRpc/proto"
	"net"
	"strconv"
	"time"
)

// 定义我们的服务，并且实现Route方法
type StreamServer struct {
}

// 实现ListValue方法
func (s *StreamServer) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer)  error {
	//for n := 0; n < 5; n++ {
	for n := 0; n < 15; n++ {

		// 向流中发送下次，默认每次send消息， 最大长度是 math.MathInt32 bytes
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil{
			return err
		}

		log.Println(n)
		time.Sleep(time.Second)
	}



	return nil
}

// Route实现Route方法
func (s *StreamServer) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
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
	pb.RegisterStreamServerServer(grpcServer, &StreamServer{})

	// 需要用Server()方法以及我们的端口信息，去阻塞等待，直到进程被杀死
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("grpcServer.server err : %v", err)
	}
}
