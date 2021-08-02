package main

import "fmt"

func main() {

	ch3 := "lu飞a育"

	// 下标遍历
	for i:=0; i< len(ch3); i++{
		fmt.Printf("[ascii:%c, %d]\n", ch3[i], ch3[i])
	}
	// for range遍历
	for _,i := range ch3{
		fmt.Printf("[unicode: %c, %d]\n", i, i )
	}

	/*
		[ascii:l, 108]
		[ascii:u, 117]
		[ascii:é, 233]
		[ascii:£, 163]
		[ascii:, 158]
		[ascii:a, 97]
		[ascii:è, 232]
		[ascii:, 130]
		[ascii:², 178]
		[unicode: l, 108]
		[unicode: u, 117]
		[unicode: 飞, 39134]
		[unicode: a, 97]
		[unicode: 育, 32946]
	*/


}

