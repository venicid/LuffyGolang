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

type Teacher struct {
	TeacherId int
	p Person  //命名结构体
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

	// 结构体匿名嵌入
	t1 := Teacher{
		TeacherId: 999,
		p:         p1,
	}


	fmt.Printf("[匿名结构体可以直接访问属性名][StudentId:%v][Name:%v][Name:%v]\n", s1.StudentId, s1.Name, s1.Age)
	fmt.Printf("[匿名结构体可以加嵌入的结构体名称访问继承的属性名][Name:%v][age:%v]\n", s1.Person.Name, s1.Person.Age)

	fmt.Printf("[命名嵌入，访问继承的属性必须加上嵌入的字段名][Name:%v][age:%v]\n", t1.p.Name, t1.p.Age)
	/*
	[匿名结构体可以直接访问属性名][StudentId:15276][Name:abc][Name:18]
	[匿名结构体可以加嵌入的结构体名称访问继承的属性名][Name:abc][age:18]
	[命名嵌入，访问继承的属性必须加上嵌入的字段名][Name:abc][age:18]
	*/
}