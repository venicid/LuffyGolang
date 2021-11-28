package main

import (
	"log"
	"net"
	"net/rpc"
)

// 我们先构造一个helloService类型方法，其中hello方法用于实现打印功能
type HelloService struct {

}

func (p *HelloService) Hello(request string, replay *string) error  {
	*replay = "hello:" + request
	return nil
}

func main()  {
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		return
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil{
		log.Fatal("ListenTCP ERROR", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept Error")
	}

	rpc.ServeConn(conn)

}