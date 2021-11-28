package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)




type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, replay *int) error  {
	*replay = args.A * args.B
	return nil
}

func main()  {
	arith := new(Arith)
	err := rpc.Register(arith)
	if err != nil {
		return
	}

	listener, err := net.Listen("tcp", ":1235")
	if err != nil{
		log.Fatal("ListenTCP ERROR", err)
	}


	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept Error")
	}
	//rpc.ServeConn(jsonrpc.NewServerCodec(conn))
	jsonrpc.ServeConn(conn)
}