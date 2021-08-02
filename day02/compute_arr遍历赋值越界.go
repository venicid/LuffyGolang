package main

import "fmt"

func main() {

	var arr1 [10] int

	//根据索引赋值
	for i:=0; i<10; i++{
		arr1[i] = i
	}

	//根据索引查询数据
	for i:=0; i<len(arr1); i++{
		fmt.Println(arr1[i])
	}

	// 越界
	// .\compute.go:19:18: invalid array index 20 (out of bounds for 10-element array)
	//fmt.Println(arr1[20])

	// 骗过编译器
	// panic: runtime error: index out of range [20] with length 10
	a := 20
	fmt.Println(arr1[a])
}