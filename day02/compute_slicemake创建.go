package main

import "fmt"

func main(){

	// 使用make初始化一个长度为0的slice
	s1 := make([]int, 0)
	s1 = append(s1, 1)
	s1 = append(s1, 2)
	s1 = append(s1, 3)
	fmt.Println(s1)  // [1 2 3]

	// 使用make初始化一个长度为5，容量为5的slice
	s2 := make([]int, 5, 5)
	//s2 := make([]int, 5, 10)  // 也是5个
	//s2 := make([]int, 5, 3)  // : len larger than cap in make([]int)
	fmt.Println(s2)  // [0 0 0 0 0]
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	s2 = append(s2, 3)
	fmt.Println(s2)  // [0 0 0 0 0 1 2 3]
}