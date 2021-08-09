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

// 函数外部定义
var p11 = Person{"jack", 33, map[string]string{}}

func main() {

	// 方法1：var声明
	var p Person
	fmt.Println(p)
	fmt.Printf(" %+v", p)
	/*
		{ 0 map[]}
		 {Name: Age:0 Labels:map[]}
	*/

	var p1 = Person{Name: "alex"}
	fmt.Println(p1)
	/*
	 {Name: Age:0 Labels:map[]}{alex 0 map[]}
	*/

	fmt.Println(p11)
	/*
	   {jack 33 map[]}
	*/

}
