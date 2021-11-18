package main

import (
	"io"
	"log"
	"net"
)

func start(){
	listener, err := net.Listen("tcp", "127.0.0.1:8866")
	if err !=nil{
		log.Fatal(err)
	}

	defer listener.Close()

	for{
		con, err := listener.Accept()
		if err !=nil{
			log.Println(err)
			continue
		}

		defer con.Close()

		for {
			var data = make([]byte, 1024)
			n, err := con.Read(data)
			if err !=nil && err != io.EOF{
				log.Println(err)
			}

			if n>0{
				log.Println("received msg", n , "bytes:", string(data[:n]))
			}
		}
	}
}

func main()  {
	start()

}


/*
粘包 1

2021/11/18 08:46:16 received msg 35 bytes: 0aaa
0bbb
0ccc
1aaa
1bbb
1ccc
2aaa

2021/11/18 08:46:16 received msg 115 bytes: 2bbb
2ccc
3aaa
3bbb
3ccc
*/


/*
粘包2

7ccc
8[路飞学城，路飞学城，路飞学城，路飞学城，路飞学城�
2021/11/18 08:48:37 received msg 192 bytes: �路飞学城，路飞学城，路飞学城，]
8bbb
8ccc
9[路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，]
9bbb
*/


/*
访问百度

$ curl -v http://www.baidu.com
< Content-Length: 2381


*/