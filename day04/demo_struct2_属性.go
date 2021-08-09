package main

import "fmt"

// 结构体定义
type Person struct {
	Name   string
	Age    int
	Labels map[string]string
}


func main() {
	p1 := Person{
		Name:   "abc",
		Age:    18,
		Labels: nil,
	}

	fmt.Printf("[Name:%v][Name:%v][Lables:%v]\n", p1.Name, p1.Age, p1.Labels)
	p1.Age += 1
	fmt.Printf("[Name:%v][Name:%v][Lables:%v]", p1.Name, p1.Age, p1.Labels)
	/*
		[Name:abc][Name:18][Lables:map[]]
		[Name:abc][Name:19][Lables:map[]]
	*/
}