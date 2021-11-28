package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	pb "main/01_SimpleRpc/proto"
	"net"
	"runtime"
	"time"
)


func main()  {

	listener, err := net.Listen("tcp", ":8083")
	if err != nil{
		log.Printf("net listener err : %v", err)
	}

	log.Println(":8083 net.listening...")

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleServer(grpcServer, &SimpleServer{})
	err = grpcServer.Serve(listener)
	if err != nil{
		log.Fatalf("%v", err)
	}
}

type SimpleServer struct {

}


func (s *SimpleServer) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error)  {
	data := make(chan *pb.SimpleResponse, 1)
	go handle(ctx, req, data)
	select {
		case res := <- data:
			return res, nil
		case <- ctx.Done():
			return nil, status.Errorf(codes.Canceled, "Client cancel")
	}

}

func handle(ctx context.Context, req *pb.SimpleRequest, data chan <- *pb.SimpleResponse)  {
	select {
		case <- ctx.Done():
			log.Println(ctx.Err())
			runtime.Goexit() // 超时后，退出该goroutine
		case <- time.After(4 * time.Second):  // 模拟耗时操作
			res := pb.SimpleResponse{
				Code:  200,
				Value: "hello "+ req.Data,
			}
			data <- &res
	}


}


/*
2021/11/29 00:30:00 :8083 net.listening...
2021/11/29 00:48:27 context deadline exceeded

*/