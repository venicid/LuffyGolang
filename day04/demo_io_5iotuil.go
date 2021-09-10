package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	// 读取文件
	bytes, err := ioutil.ReadFile("day04/go.mod")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", bytes)
	fmt.Printf("%v", string(bytes))


	// 写入文件
	fileName := "a.txt"
	err1 := ioutil.WriteFile(fileName, []byte("升职加薪"), 0644)
	fmt.Println(err1)

	// 读取目录下的文件元信息
	fs, _ := ioutil.ReadDir("day04")

	for _, f :=range  fs{
		log.Printf("[name:%v][size:%v][mode:%v][modTime:%v]",
			f.Name(),
			f.Size(),
			f.Mode(),
			f.ModTime(),
		)
	}
/*
   module day04

   go 1.16
   module day04

   go 1.16
*/
}

/*
2021/09/10 21:49:56 [name:01_函数.md][size:15014][mode:-rw-rw-rw-][modTime:2021-08-09 22:46:18.57 +0800 CST]
2021/09/10 21:49:56 [name:02结构体][size:0][mode:drwxrwxrwx][modTime:2021-08-10 23:48:44.2855053 +0800 CST]
2021/09/10 21:49:56 [name:03_面向对象和接口.assets][size:0][mode:drwxrwxrwx][modTime:2021-08-11 07:40:04.7111253 +0800 CST]
*/

