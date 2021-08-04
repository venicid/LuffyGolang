package main

import "fmt"

func main(){

	a1 := []int{10,20,30}
	// 遍历查询
	for index, value := range a1{
		fmt.Printf("[index: %d, value:%d]\n", index, value)
	}

	// 遍历修改值
	for index, _ := range a1{
		a1[index] += 100
	}
	fmt.Println(a1)

	/*
	[index: 0, value:10]
	[index: 1, value:20]
	[index: 2, value:30]
	[110 120 130]
	*/
}