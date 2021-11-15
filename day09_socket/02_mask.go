
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
		os.Exit(1)
	}


	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	_,bits := mask.Size()
	fmt.Println(
		"this address is ", addr.String(),
		"Deafault mask length is", bits,
		"Mask is (hex)", mask.String(),
		"Network is", network.String(),
		)



}
