/*
Listening on :3333
收到信息127.0.0.1:14152 -> 127.0.0.1:3333
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

func handleNewRequest(conn net.Conn)  {
	defer conn.Close()
	for{
		io.Copy(conn, conn)
	}

}

func main()  {

	var host = flag.String("host", "", "host")
	var port = flag.String("port", "3333", "port")
	flag.Parse()

	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host +":" + *port)
	if err != nil{
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on "+ *host + ":" + *port)

	for{
		conn, err := l.Accept()
		if err != nil{
			fmt.Println("Error accepting:", err)
		}
		fmt.Printf("收到信息%s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		// 每次处理连接，创建1个goroutine
		go handleNewRequest(conn)
	}
}