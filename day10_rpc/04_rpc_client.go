package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 首先通过rpc.Dail 拨号rpc服务，然后

func main()  {

	client, err := rpc.Dial("tcp", ":1234")
	if err != nil{
		log.Fatal("dialing :" ,err)
	}

	var replay string

	// 通过client.Call调用具体的rpc方法，在调用时，第一个参数用点号链接的rpc服务名和方法名
	// 第二个和第三个参数表示，我们定义rpc方法的两个参数
	err = client.Call("HelloService.Hello", "hello", &replay)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(replay)

}


