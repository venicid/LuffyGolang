package main

import "log"


// 返回多个匿名函数
func FGen(x, y int) (func() int, func(int) int){

	// 求和匿名
	f1 := func() int{
		return x+y
	}

	// (x+y) *z 的匿名函数
	avg := func(z int)int {
		return (x+y)/z
	}

	return f1, avg

}

func main()  {

	f1, f2 := FGen(1,2)
	res1 := f1()
	res2 := f2(3)
	log.Println(res1)
	log.Println(res2)

	/*
	2021/08/08 16:02:31 3
	2021/08/08 16:02:31 1

	*/

}