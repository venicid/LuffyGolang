
/*
echo协议 rfc 862协议
用go实现一个简单的tcp服务器和客户端

服务器只需要把收到的客户端请求数据发给客户端
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error)  {
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn)  {

	defer conn.Close()

	var buf [512] byte
	for{
		n, err:= conn.Read(buf[0:])
		if err != nil{
			return
		}
		fmt.Println(string(buf[0:]))
		_,err2 := conn.Write(buf[0:n])
		if err2 != nil{
			return
		}
	}
}

func main()  {
	service := ":12001"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for{
		conn, err := listener.Accept()
		if err != nil{
			return
		}

		go handleClient(conn)
	}

}