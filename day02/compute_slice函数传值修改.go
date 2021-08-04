package main

import "fmt"

func showSlice(s []int)  {
	fmt.Printf("[传入的切片为：%v]\n", s)
	s[2] = 30

}

func main(){
	a1 := []int{1,2,3}
	showSlice(a1)
	fmt.Printf("[函数执行后的切片为：%v]\n", a1)

	/*
	[传入的切片为：[1 2 3]]
	[函数执行后的切片为：[1 2 30]]
	*/
}
