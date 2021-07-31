package main

import "fmt"

func main() {

	var a = 21
	var b = 10
	var c int
	fmt.Printf("[初始化的值：a=%d,b=%d,c=%d]\n", a, b, c)
	// 加法
	c = a + b
	fmt.Printf("[加法操作：a+b = %d]\n", c)
	// 减法
	c = a - b
	fmt.Printf("[减法操作：a-b=%d]\n", c)
	// 乘法
	c = a * b
	fmt.Printf("[乘法操作 a*b=%d]\n", c)
	// 除法
	c = a / b
	fmt.Printf("[除法操作 a/b=%d]\n", c)
	// 取余
	c = a % b
	fmt.Printf("[取余操作 a%%b=%d]\n", c)
	// 自增
	c++
	fmt.Printf("[取余操作 c++ =%d]\n", c)
	// 自减
	c--
	fmt.Printf("[取余操作 c--=%d]\n", c)
	/**
	[初始化的值：a=21,b=10,c=0]
	[加法操作：a+b = 31]
	[减法操作：a-b=11]
	[乘法操作 a*b=210]
	[除法操作 a/b=2]
	[取余操作 a%b=1]
	[取余操作 c++ =2]
	[取余操作 c--=1]
	**/
}
