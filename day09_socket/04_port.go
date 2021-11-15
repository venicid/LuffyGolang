/*
解析80端口
go run 1.ip.go 1.1.1.1 80
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {

	networkType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(networkType, service)
	if err !=nil{
		fmt.Println("Err", err.Error())
		os.Exit(2)
	}

	fmt.Println("port", port)
	os.Exit(0)
}

/*

E:\golang\HelloGolang\day09_socket>go run 04_port.go 1.1.1.1 80
port 80

E:\golang\HelloGolang\day09_socket>go run 04_port.go 1.1.1.1 80dafa
Err address 1.1.1.1: unknown network
exit status 2

*/