package main

import "log"

type Person struct {
	Name string
	Age  int
}

func main() {

	p1 := Person{
		Name: "123",
		Age:  123,
	}
	p2 := p1

	p2.Age = 100
	p1.Name = "456"

	log.Printf("结构体中的字段都是值类型，那么就是深拷贝")
	log.Printf("[p1的内存地址:%p ][value:%+v]", &p1, p1)
	log.Printf("[p1的内存地址:%p ][value:%+v]", &p2, p2)
}