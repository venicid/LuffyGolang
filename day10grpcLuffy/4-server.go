package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "main/04_BothSteamRpc/proto"
	"net"
	"strconv"
)

type StreamService struct {

}

func (s *StreamService) Conversations(srv pb.StreamService_ConversationsServer) error {
	n := 1
	for{
		req, err := srv.Recv()
		if err == io.EOF{
			return nil
		}
		if err != nil{
			return err
		}

		err = srv.Send(&pb.StreamResponse{
			Answer: "from stream server answer: the" +strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil{
			return err
		}
		n ++
		log.Printf("form stream client question: %s", req.Question)
	}

}

// Route实现Route方法
func (s *StreamService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}

	return &res, nil
}

func main()  {
	listener, err := net.Listen("tcp", ":8082")
	if err !=nil{
		log.Fatal("net listen err: %v", err)
	}
	log.Println(":8080 net.Listening...")

	grpcServer := grpc.NewServer()
	pb.RegisterStreamServiceServer(grpcServer, &StreamService{})

	err = grpcServer.Serve(listener)
	if err !=nil{
		log.Fatal("grpcServer.server err : %v", err)
	}
}


/*
2021/11/28 23:41:44 :8080 net.Listening...
2021/11/28 23:51:01 form stream client question: stream client rpc0
2021/11/28 23:51:01 form stream client question: stream client rpc1
2021/11/28 23:51:01 form stream client question: stream client rpc2
2021/11/28 23:51:01 form stream client question: stream client rpc3
2021/11/28 23:51:01 form stream client question: stream client rpc4
*/