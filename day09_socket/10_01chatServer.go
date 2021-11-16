/*
服务器端：
   1。接收来自客户端的连接请求并建立连接
   2。 所有的客户端连接会放进连接池，用于广播消息

客户端：
   连接服务器；
   向服务器发送消息
   接收服务器的广播消息

注意事项：
    某一个客户端断开连接后需要从连接池摘除，并不再接受广播消息
    某个客户端断开连接后不能影响服务器端或别的客户端连接
*/

/*
listening port 9080 ...
 map[127.0.0.1:11767:0xc000006038]
listening port 9080 ...
 connect from client 127.0.0.1:11767
127.0.0.1:11767say:1
connection is connected from  127.0.0.1:11767
*/

package main

import (
	"fmt"
	"log"
	"net"
)

// 监听port，是否有client连接到server
func start(port string)  {
	host := ":" + port

	// 解析
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil{
		log.Printf("reslove tcp addr failed: %v\n", err.Error())
		return
	}

	// 监听
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil{
		log.Printf("listen tcp port failed: %v%n", err.Error())
		return
	}

	// 建立连接池，用于广播消息
	conns := make(map[string]net.Conn)  // 1.1.1.1 : conn

	// 消息通道
	messageChan := make(chan string, 10)

	// 广播消息
	go broadMessage(&conns, messageChan)

	// 启动
	for{
		fmt.Printf("listening port %s ...\n ", port)
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("接受失败 %v\n", err)
			continue
		}

		// 把每个客户端连接扔入连接池
		conns[conn.RemoteAddr().String()] = conn
		fmt.Println(conns)

		// 处理消息
		go handler(conn, &conns, messageChan)
	}

}

// 向所有client广播
func broadMessage(conns *map[string] net.Conn, messageChan chan string)  {
	for{
		// 不断从通道中读取消息
		msg := <- messageChan
		fmt.Println(msg)

		// 向所有client发送消息
		for key, conn := range *conns {
			fmt.Println("connection is connected from ", key)
			_, err := conn.Write([]byte(msg))
			if err != nil{
				log.Printf("broad message to %s failed\n", key)
				delete(*conns, key)
			}
		}
	}

}

// 处理client发送到服务端的消息，将其扔到通道中
func handler(conn net.Conn, conns *map[string]net.Conn, messageChan chan  string)  {
	fmt.Println("connect from client", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for{
		length, err := conn.Read(buf)
		if err != nil{
			log.Printf("read client message failed: %v\n", conn)
			delete(*conns, conn.RemoteAddr().String())
			conn.Close()
			break
		}
		recvStr := string(buf[0:length])
		messageChan <- recvStr
	}
}



func main()  {
	port := "9080"
	start(port)

}