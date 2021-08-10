package main

import (
	"log"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	StudentId int
}

// 给Person结构体绑定一个SayHello
func (p Person) SayHello() {
	log.Printf("[Person.SayHello][name:%v]", p.Name)
}

func main() {
	p1 := Person{
		Name: "小乙",
		Age:  123,
	}

	s1 := Student{
		Person:    p1,
		StudentId: 99,
	}
	s1.SayHello()

	s2 := Student{
		Person:    Person{
			Name: "alex",
			Age : 33,
		},
		StudentId: 0,
	}
	s2.SayHello()
}

/*
2021/08/10 23:54:30 [Person.SayHello][name:小乙]
2021/08/10 23:54:30 [Person.SayHello][name:alex]
*/