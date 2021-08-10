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
	p2 := &p1
	// 等价于
	var p3 *Person
	p3 = &p1

	p1.Age = 100
	(*p2).Name = "456"

	log.Printf("结构体中的字段都是值类型，那么就是深拷贝")
	log.Printf("[p1的内存地址:%p ][value:%+v]\n", &p1, p1)
	log.Printf("[p2的内存地址:%p ][value:%+v]\n", p2, p2)
	log.Printf("[p3的内存地址:%p ][value:%+v]\n", p3, p3)
}

/*
2021/08/10 07:19:11 结构体中的字段都是值类型，那么就是深拷贝
2021/08/10 07:19:11 [p1的内存地址:0xc000098060 ][value:{Name:456 Age:100}]
2021/08/10 07:19:11 [p2的内存地址:0xc000098060 ][value:&{Name:456 Age:100}]
2021/08/10 07:19:11 [p3的内存地址:0xc000098060 ][value:&{Name:456 Age:100}]
*/