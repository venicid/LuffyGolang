
/*
判断输入的参数是合法ip
go run 1.ip.go 1.1.1.1
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil{
		fmt.Println("非法的ip")
	}else{
		fmt.Println("this address is ", addr.String())
	}
}
