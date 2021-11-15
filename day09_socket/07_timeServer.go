package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func checkError1(err error)  {
	if err != nil{
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main()  {

	service := "127.0.0.1:12000"
	tcpAddr1, err := net.ResolveTCPAddr("tcp4", service)
	checkError1(err)

	listener, err := net.ListenTCP("tcp", tcpAddr1)
	checkError1(err)

	for{
		conn, err := listener.Accept()
		if err != nil{
			continue
		}

		dayTime := time.Now().String()
		conn.Write([]byte(dayTime))
		conn.Close()

	}

}

/*
telnet 127.0.0.1:12000

*/