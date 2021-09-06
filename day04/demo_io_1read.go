package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	reader := strings.NewReader("xiaoyi 123dwd 123")
	// 每次读取4个字节
	p := make([]byte, 4)
	for {

		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				log.Printf("读完了:eof错误 :%d", n)
				break
			}
			log.Printf("其他错误:%v", err)
			os.Exit(2)
		}
		log.Printf("[读取到的字节数为:%d][内容:%v]", n, string(p[:n]))
		log.Printf("[读取到的字节数为:%d][内容:%v]", n, string(p))
	}

}

/*
2021/08/31 08:29:16 [读取到的字节数为:4][内容:yi 1]
2021/08/31 08:29:16 [读取到的字节数为:4][内容:23dw]
2021/08/31 08:29:16 [读取到的字节数为:4][内容:23dw]
2021/08/31 08:29:16 [读取到的字节数为:4][内容:d 12]
2021/08/31 08:29:16 [读取到的字节数为:4][内容:d 12]

// 倒数第二次  p为d 12
// 倒数第1次  读取到3，p为 3 12
2021/08/31 08:29:16 [读取到的字节数为:1][内容:3]
2021/08/31 08:29:16 [读取到的字节数为:1][内容:3 12]
2021/08/31 08:29:16 读完了:eof错误 :0
*/