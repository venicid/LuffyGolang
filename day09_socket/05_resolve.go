/*
解析ip
go run 1.ip.go www.baidu.com
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {

	network := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", network)
	if err !=nil{
		fmt.Println("Err", err.Error())
		os.Exit(2)
	}
	fmt.Println("addr", addr)
	os.Exit(0)


}


/*
E:\golang\HelloGolang\day09_socket>go run 05_resolve.go www.baidu.com
addr 36.152.44.96

E:\golang\HelloGolang\day09_socket>go run 05_resolve.go www.mi.com
addr 36.150.14.1

*/