package main

import "log"

// 变长参数 返回最小值
func min(a ...int)  int{
	if len(a) == 1{
		return a[0]
	}

	min := a[0]
	for _,v := range a{
		if v< min{
			min = v
		}
	}
	return min
}

func main() {

	x := min(4)
	log.Println(x)

	x1 := min(1, 7, 8, 3, 23, 9)
	log.Printf("[直接传多个参数]：%d", x1)

	// 数组传值 arr...
	s1 := []int{2, 3, 54, 10, 32, 5, 7, 34}
	x2 := min(s1...)
	log.Printf("[数组传参传]：%d", x2)
}
/*
2021/08/08 15:16:25 4
2021/08/08 15:16:25 [直接传多个参数]：1
2021/08/08 15:16:25 [数组传参传]：2

*/