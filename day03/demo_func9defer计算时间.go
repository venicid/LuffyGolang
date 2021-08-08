package main

import (
	"log"
	"time"
)

func calcTime() {
	start := time.Now()
	time.Sleep(3 * time.Second)
	end := time.Since(start)
	log.Println(end)

	/*
	2021/08/08 18:49:16 3.001259487s
	*/
}

func calcTimeDefer()  {
	start :=time.Now()
	defer log.Printf("时间差：%v", time.Since(start))
	time.Sleep(3*time.Second)

	log.Printf("函数结束")

	/*
	2021/08/08 18:51:05 时间差：207ns
	*/
}

func calcTimeDeferFun()  {
	start := time.Now()
	log.Printf("开始时间为：%v", start)
	defer func() {
		log.Printf("开始调用defer")
		log.Printf("时间差：%v", time.Since(start))
		log.Printf("结束调用defer")
	}()
	time.Sleep(3*time.Second)
	log.Printf("函数结束")

}

func main() {
	calcTime()
	calcTimeDefer()
	calcTimeDeferFun()
}


/*
2021/08/08 19:05:04 3.004173012s
2021/08/08 19:05:07 函数结束
2021/08/08 19:05:07 时间差：285ns
2021/08/08 19:05:10 3.003485295s

*/