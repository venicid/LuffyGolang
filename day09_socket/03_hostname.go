
/*
判断输入的参数是合法ip
go run 1.ip.go www.baidu.com
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	name := os.Args[1]
	addrs, err := net.LookupHost(name)
	if err !=nil{
		fmt.Println("error", err.Error())
		os.Exit(1)
	}
	fmt.Println(addrs)

	for _, s:= range addrs{
		fmt.Println(s)
	}
	os.Exit(0)

}
