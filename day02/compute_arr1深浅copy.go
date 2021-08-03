package main

import "fmt"

func main() {

	arr1 := [2]int{1,2}
	var arr2 [2]int
	arr2 = arr1  // 深copy

	fmt.Printf("arr1[%v %p]\n", arr1, &arr1)
	fmt.Printf("arr2[%v %p]\n", arr2, &arr2)

	arr2[1] = 3   // 浅copy
	fmt.Printf("arr1[%v %p]\n", arr1, &arr1)
	fmt.Printf("arr2[%v %p]\n", arr2, &arr2)
	/*
	arr1[[1 2] 0xc00000a0b0]
	arr2[[1 2] 0xc00000a0c0]
	arr1[[1 2] 0xc00000a0b0]
	arr2[[1 20] 0xc00000a0c0]
	*/
}