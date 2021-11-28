package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "main/04_BothSteamRpc/proto"
	"strconv"
)

var streamClient pb.StreamServiceClient

func route()  {

	req := pb.SimpleRequest{
		Data: "grpc",
	}
	res, _ := streamClient.Route(context.Background(), &req)
	log.Println(res.Value)
}

func conversations()  {
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err : %v", err)
	}
	for n := 0; n < 5; n++ {
		err := stream.Send(&pb.StreamRequest{
			Question: "stream client rpc" +strconv.Itoa(n),
		})
		if err != nil {
			log.Fatalf("stream request err :%v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err !=nil{
			log.Fatalf("Converstation  get stream err:%v", err)
		}
		log.Println(res.Answer)

	}

	// 最后关闭流
	err = stream.CloseSend()
	if err !=nil{
		log.Fatalf("Conversation close stream err: %v", err)
	}

}

func main()  {
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err!= nil{
		log.Fatalf("net.Connect err:%v", err)
	}

	defer conn.Close()

	streamClient = pb.NewStreamServiceClient(conn)
	conversations()
}


/*
2021/11/28 23:51:01 from stream server answer: the1 question is stream client rpc0
2021/11/28 23:51:01 from stream server answer: the2 question is stream client rpc1
2021/11/28 23:51:01 from stream server answer: the3 question is stream client rpc2
2021/11/28 23:51:01 from stream server answer: the4 question is stream client rpc3
2021/11/28 23:51:01 from stream server answer: the5 question is stream client rpc4
*/