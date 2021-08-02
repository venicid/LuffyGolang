package main

import "fmt"

var QUANJUBL string = "我是全局变量"

func T1() string {
	str1 := "函数中的字符串变量"
	fmt.Println(str1)
	fmt.Println(QUANJUBL)
	return str1
}

func main()  {
	// 试图引用在函数里面定义的变量
	// .\compute.go:12:14: undefined: str1
	//fmt.Println(str1)

	for i := 0; i < 10 ; i++{
		fmt.Println(i)
		fmt.Println(QUANJUBL)
	}

	// 试图引用在for里面定义的变量
	// .\compute.go:19:14: undefined: i
	//fmt.Println(i)

	if str := T1(); str == "" {
		fmt.Println("[函数返回为空]")
		fmt.Println(QUANJUBL)
	}
	str := T1()
	if str == ""{}

	// 试图引用在if里面定义的变量
	// .\compute.go:28:14: undefined: str
	//fmt.Println(str)

}
