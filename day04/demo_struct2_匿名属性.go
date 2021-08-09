package main

import "fmt"

// 结构体定义
type Person struct {
	Name   string
	Age    int
	Labels map[string]string
}

type Student struct {
	StudentId int
	Person  // 匿名结构体，匿名属性，嵌套，继承
}


func main() {
	p1 := Person{
		Name:   "abc",
		Age:    18,
		Labels: nil,
	}
	s1 := Student{
		StudentId:    15276,
		Person: p1,
	}

	fmt.Printf("[Name:%v][Name:%v][Lables:%v]\n", p1.Name, p1.Age, p1.Labels)
	fmt.Printf("[StudentId:%v][Name:%v][Name:%v][Lables:%v]\n", s1.StudentId, s1.Name, s1.Age, s1.Labels)
	/*
	[Name:abc][Name:18][Lables:map[]]
	[StudentId:15276][Name:abc][Name:18][Lables:map[]]
	*/
}