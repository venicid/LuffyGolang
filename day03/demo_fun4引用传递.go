package main

import (
	"log"
	"time"
)

// 值类型 值传递
func add1(num int) {
	log.Printf("[值传递][传入的参数值为:%d]", num)
	num++
	log.Printf("[值传递][add1计算后的值为:%d]", num)
}

// 值类型 引用传递
func add2(num *int) {
	log.Printf("[引用传递][传入的参数值为:%d]", *num)
	*num++
	log.Printf("[引用传递][add2计算后的值为:%d]", *num)
}

// 引用类型 引用传递
func mod(s1 []int, m1 map[string]string){

	log.Printf("[引用传递][传入的参数为:%v %v]", s1, m1)
	s1[0] = 100
	m1["a"] = "a2"
	log.Printf("[引用传递][函数内部处理完的值为:%v %v]", s1, m1)
}



func main()  {

	num := 1
	log.Printf("[局部遍历的值:%d]", num)

	add1(num)
	time.Sleep(1 * time.Second)
	log.Printf("[局部遍历的值:%d]", num)

	add2(&num)
	time.Sleep(1 * time.Second)
	log.Printf("[局部遍历的值:%d]", num)

	s1 := []int{1,2,3}
	m1 := map[string]string{"a1":"b1","a2":"b2"}
	mod(s1, m1)
	log.Printf("[引用传递][函数外部的值为:%v %v]", s1, m1)
}

/*
2021/08/08 15:42:24 [局部遍历的值:1]
2021/08/08 15:42:24 [值传递][传入的参数值为:1]
2021/08/08 15:42:24 [值传递][add1计算后的值为:2]

2021/08/08 15:42:25 [局部遍历的值:1]
2021/08/08 15:42:25 [引用传递][传入的参数值为:1]
2021/08/08 15:42:25 [引用传递][add2计算后的值为:2]
2021/08/08 15:42:26 [局部遍历的值:2]

2021/08/08 15:42:26 [引用传递][传入的参数为:[1 2 3] map[a1:b1 a2:b2]]
2021/08/08 15:42:26 [引用传递][函数内部处理完的值为:[100 2 3] map[a:a2 a1:b1 a2:b2]]
2021/08/08 15:42:49 [引用传递][函数外部的值为:[100 2 3] map[a:a2 a1:b1 a2:b2]]

*/