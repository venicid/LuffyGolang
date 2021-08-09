package main

import (
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age int
	Labels map[string]string
}


func main()  {
	// make初始化为0值
	p1 := new(Person)

	// make返回指针
	p1 = &Person{
		Age:  10,
		Name: "小乙",
		Labels: map[string]string{},
	}

	log.Printf(" %+v", p1)
	// 2021/08/09 23:21:10  &{Name:小乙 Age:10 Labels:map[]}



	// new返回引用，对应着初始化
	var p2 *Person= new(Person)
	p2.Name = "alex"
	p2.Age = 3

	var p3 Person = Person{
		Name: "李逵",
		Age:  20,
	}
	fmt.Println(p2, p3)
	// &{alex 3 map[]} {李逵 20 map[]}
}