package main

import "fmt"

func main() {

	var arr [5]string
	fmt.Println(arr)

	arr[1] = "golang"
	fmt.Println(arr)

	var arr1 = [...]int{1,2,3}
	fmt.Println(arr1)

	/*
	[    ]
	[ golang   ]
	[1 2 3]
	*/
}