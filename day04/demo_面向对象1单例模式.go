package main

import (

	"log"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) changeAge1()  {
	p.Age += 10
	log.Printf("[单实例绑定方法][Person.ChangeAge1][p.Age:%v]", p.Age)
}

// [多实例] 多指针，单例模式
// 值在函数内部值被修改了，在外部没有被修改
func (p Person) changeAge2()  {
	p.Age += 10
	log.Printf("[单例模式][非指针型绑定][Person.ChangeAge2][p.Age:%v]", p.Age)
}

type Student struct {
	Person
	StudentId int
}


func main() {
	p1 := Person{
		Name: "小乙",
		Age:  100,
	}

	s1 := Student{
		Person:    p1,
		StudentId: 2339,
	}
	log.Println(s1.Age)
	s1.changeAge1()
	log.Println(s1.Age)

	log.Println(s1.Age)
	s1.changeAge2()
	log.Println(s1.Age)



}

/*
2021/08/11 00:05:20 100
2021/08/11 00:05:20 [单实例绑定方法][Person.ChangeAge1][p.Age:110]
2021/08/11 00:05:20 110
2021/08/11 00:05:20 110
2021/08/11 00:05:20 [单例模式][非指针型绑定][Person.ChangeAge2][p.Age:120]
2021/08/11 00:05:20 110
*/

/*
go 单实例，把值的地址传递进去，直接修改的是原数据

go 函数默认值传递，传递的是一份copy
*/
