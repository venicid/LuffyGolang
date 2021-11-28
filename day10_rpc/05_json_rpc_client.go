package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)


// 首先通过rpc.Dail 拨号rpc服务，然后

func main()  {

	client, err := jsonrpc.Dial("tcp", ":1235")
	if err != nil{
		log.Fatal("dialing :" ,err)
	}


	var reply int

	type Args struct {
		A, B int
	}
	args := Args{3,4}

	// 通过client.Call调用具体的rpc方法，在调用时，第一个参数用点号链接的rpc服务名和方法名
	// 第二个和第三个参数表示，我们定义rpc方法的两个参数
	err = client.Call("Arith.Multiply", args, &reply )
	if err != nil{
		log.Fatal("Call ", err)
	}
	fmt.Println(reply)

}


