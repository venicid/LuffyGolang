package main

import "log"

func add1() int{
	sum := 0
	for i:=0;i<100;i++{
		sum += i
	}
	log.Printf("[普通的累加器]:%d", sum)
	return sum
}

func add2() func(int) int {
	// 自由变量
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func callCF() int{
	f := add2()

	sum := 0
	for i:=0; i<100;i++{
		sum = f(i)
	}
	log.Printf("[闭包的累加器]：%d", sum)
	return sum
}

func main()  {
	res1 := add1()
	log.Println(res1)

	// 闭包的
	res2 := callCF()
	log.Println(res2)
}

/*
2021/08/08 16:47:25 [普通的累加器]:4950
2021/08/08 16:47:25 [闭包的累加器]：4950
2021/08/08 16:47:25 4950
2021/08/08 16:47:25 4950
*/