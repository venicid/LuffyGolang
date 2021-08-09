package main

import "fmt"

// 结构体定义
type Person struct {
	Name   string
	Age    int
	Labels map[string]string
}

type Leader struct {
	LeaderId int
	*Person  //结构体匿名 指针嵌入
}

func main() {

	l1 := Leader{}

	fmt.Println(l1)   // {0 <nil>}
	fmt.Printf(l1.Age)
	/*
	# command-line-arguments
	.\demo.go:22:15: cannot use l1.Person.Age (type int) as type string in argument to fmt.Printf
	*/




}