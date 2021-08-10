package main

import "log"

type Person struct {
	Name string
	Age  int
}

func main() {

	p1 := new(Person)

	p1.Name = "小乙"
	p1.Age = 123

	p2 := p1

	//log.Printf("结构体中的字段都是值类型，使用&赋值给另外一个，就是浅拷贝")

	p1.Age = 19
	p2.Name = "898"
	log.Printf("[p1的内存地址:%p ][value:%+v]", p1, p1)
	log.Printf("[p2的内存地址:%p ][value:%+v]", p2, p2)

}

/*
2021/08/10 07:18:19 [p1的内存地址:0xc000004078 ][value:&{Name:898 Age:19}]
2021/08/10 07:18:19 [p2的内存地址:0xc000004078 ][value:&{Name:898 Age:19}]
*/