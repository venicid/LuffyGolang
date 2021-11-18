// 分隔符解决粘包

package main


import (
	"bufio"
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

		reader := bufio.NewReader(con)

		for {

			data, err := reader.ReadSlice('\n')
			if err != nil{
				if err!= io.EOF{
					log.Println(err)
				}else{
					break
				}
			}
			log.Println("received msg", len(data) , "bytes:", string(data))

			/*
				var data = make([]byte, 1024)
				n, err := con.Read(data)
				if err !=nil && err != io.EOF{
					log.Println(err)
				}
				if n>0{
					log.Println("received msg", n , "bytes:", string(data[:n]))
				}
			*/

		}
	}
}

func main()  {
	start()

}


/*
2021/11/18 09:00:59 received msg 124 bytes: 6[路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，]

2021/11/18 09:00:59 received msg 5 bytes: 6bbb

2021/11/18 09:00:59 received msg 5 bytes: 6ccc

2021/11/18 09:00:59 received msg 124 bytes: 7[路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，]

2021/11/18 09:00:59 received msg 5 bytes: 7bbb
*/