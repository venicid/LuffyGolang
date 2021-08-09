package main

import (
	"fmt"
)

// 结构体定义
type Person struct {
	Name   string
	Age    int
	Labels map[string]string
}

// 不能使用短变量
p33 := Person{
Name:   "venicid",
Labels: map[string] string{},
Age:    11,
}

func main()  {


	// 方法2：短变量声明
	p3 := Person{
		Name:   "venicid",
		Labels: map[string] string{},
		Age:    11,
	}
	fmt.Printf("%v\n", p3)

	p4 := Person{"jack", 33, map[string]string{}}
	fmt.Printf("%v\n", p4)

	// 不能在函数外部声明，
	fmt.Printf("%v\n", p33)
	/*
		# command-line-arguments
		.\demo_struct1_定义初始化.go:17:1: syntax error: non-declaration statement outside 	function body
	*/
}
