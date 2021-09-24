package main

import "fmt"

func errorCatch(msg string)  {
	defer func() {
		// 捕获异常
		if info := recover(); info !=nil{
			fmt.Println("触发了宕机", info)
		}else{
			fmt.Println("程序运行正常")
		}
	}()

	panic(msg)
}

func main()  {

	errorCatch("list out of range")

}

/*
触发了宕机 list out of range
*/