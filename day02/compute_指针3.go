package main

import "fmt"

func main() {

	i := 1
	var c int
	c = i

	fmt.Printf("[i的值: %d, 类型为%T, 地址为%p]\n", i, i, &i)
	fmt.Printf("[c的值: %d, 类型为%T, 地址为%p]\n", c, c, &c)

	/*
	[i的值: 1, 类型为int, 地址为0xc00000a098]
	[c的值: 1, 类型为int, 地址为0xc00000a0b0]
	*/
}

