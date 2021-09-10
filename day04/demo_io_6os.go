package main

import (
	"log"
	"os"
)

func main() {

	// 创建文件，并写入string
	file, _ := os.Create("b.txt")
	for i:=0;i<5;i++{
		file.WriteString("WriteString\n")
		file.Write([]byte("Write\n"))
	}

	hn, _ := os.Hostname()
	log.Printf("主机名:%v", hn)
	log.Printf("进程pid:%v", os.Getpid())
	log.Printf("命令行参数:%v", os.Args)
	log.Printf("获取GOROOT 环境变量:%v", os.Getenv("GOROOT"))

	for _, v := range os.Environ() {
		log.Printf("环境变量 %v", v)
	}
	dir,_:=os.Getwd()
	log.Printf("当前目录:%v", dir)
	log.Println("创建单一config目录")
	os.Mkdir("config",0755)
	log.Println("创建层级config1/yaml/local目录")
	os.MkdirAll("config1/yaml/local",0755)

	//log.Printf("删除单一文件或目录%v",os.Remove("config"))
	//log.Printf("删除层级文件或目录%v",os.RemoveAll("config1"))
}


/*
2021/09/10 21:59:00 主机名:LAPTOP-P42T1S35
2021/09/10 21:59:00 进程pid:2376
2021/09/10 21:59:00 命令行参数:[C:\Users\hua'wei\AppData\Local\Temp\___2go_build_demo_go__1_.exe]
2021/09/10 21:59:00 获取GOROOT 环境变量:D:\program files\Go
2021/09/10 21:59:00 环境变量 =::=::\
2021/09/10 21:59:00 环境变量 ALLUSERSPROFILE=C:\ProgramData
2021/09/10 21:59:00 当前目录:E:\golang\HelloGolang
2021/09/10 21:59:00 创建单一config目录
2021/09/10 21:59:00 创建层级config1/yaml/local目录
*/