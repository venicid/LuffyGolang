/*
固定包头解决“粘包”问题
其中，最重要的点是头部的长度以及字节序，2个字节便是2^16-1个字节的内容
如果不够，那就上4个字节 0 4 8 12

什么是字节序？
小端和大端：数据的存取和读取顺序
16进制数字：0x123456 占用了3个字节
如果协议用4个字节存储数据长度
12 34 56 00 高位在左边的，就叫大端
00 56 34 12 高位在右边的，就叫小段
*/

package main


import (
	"bufio"
	"bytes"
	"encoding/binary"
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

			peek, err := reader.Peek(4)  // 读取客户端前4个字节byte，存放的是数据长度
			if err != nil{
				if err!= io.EOF{
					log.Println(err)
				}else{
					break
				}
			}
			buffer := bytes.NewBuffer(peek)
			var length int32
			err = binary.Read(buffer , binary.BigEndian, &length) // 以大端的方式，将4个字节的byte，转为int32，放入length
			if err != nil {
				log.Println(err)
			}
			if int32(reader.Buffered()) < length + 4 {
				continue
			}

			// 读取真正的消息数据
			data := make([]byte, length + 4) // 在reader中读取length+4长度的字节
			_, err = reader.Read(data)
			if err != nil{
				continue
			}

			log.Println("received msg", len(data) , "bytes:", string(data[4:]))

			/*
				data, err := reader.ReadSlice('\n')
				if err != nil{
					if err!= io.EOF{
						log.Println(err)
					}else{
						break
					}
				}
				log.Println("received msg", len(data) , "bytes:", string(data))
			*/

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
