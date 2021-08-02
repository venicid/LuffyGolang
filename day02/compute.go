package main

import "fmt"

func main() {

	var arr [5]*int

	// 根据索引赋值
	arr[0] = new(int)
	arr[1] = new(int)
	//arr[1] = new(int)  // <nil>   panic: runtime error: invalid memory address or nil pointer dereference
	arr[2] = new(int)
	arr[3] = new(int)
	arr[4] = new(int)

	fmt.Println(arr)

	*arr[0] = 10
	*arr[1] = 2
	fmt.Println(arr)
	for i:=0; i<len(arr); i++{
		fmt.Printf("[索引：%d 值是:%d]\n", i, *arr[i])
	}

	/*
	[0xc00000a098 0xc00000a0b0 0xc00000a0b8 0xc00000a0c0 0xc00000a0c8]
	[0xc00000a098 0xc00000a0b0 0xc00000a0b8 0xc00000a0c0 0xc00000a0c8]
	[索引：0 值是:10]
	[索引：1 值是:2]
	[索引：2 值是:0]
	[索引：3 值是:0]
	[索引：4 值是:0]
	*/

}