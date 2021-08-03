package main

import "fmt"

func main() {

	arr1 := [3]*string{new(string),new(string), new(string)}
	*arr1[0] = "k1"
	*arr1[1] = "k2"
	*arr1[2] = "k3"

	var arr2 [3]*string
	arr2 = arr1

	fmt.Println(arr1)
	fmt.Println(arr2)

	for i:=0; i<len(arr1); i++{
		fmt.Printf("[arr1: %d, %v, %v]\n", i, *arr1[i], &arr2[i])
		fmt.Printf("[arr2: %d, %v, %v]\n", i, *arr2[i], &arr2[i])
	}

	fmt.Printf("arr1 [%v, %p]\n", arr1, &arr1)
	fmt.Printf("arr2 [%v, %p]\n", arr2, &arr2)
	/*
	[0xc000040240 0xc000040250 0xc000040260]
	[0xc000040240 0xc000040250 0xc000040260]
	[arr1: 0, k1, 0xc000004078]
	[arr2: 0, k1, 0xc000004078]
	[arr1: 1, k2, 0xc000004080]
	[arr2: 1, k2, 0xc000004080]
	[arr1: 2, k3, 0xc000004088]
	[arr2: 2, k3, 0xc000004088]

	arr1 [[0xc000040240 0xc000040250 0xc000040260], 0xc000004078]
	arr2 [[0xc000040240 0xc000040250 0xc000040260], 0xc000004090]
	*/

}