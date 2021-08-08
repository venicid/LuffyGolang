package main

import (
	"fmt"
)

func Greeting() func(string) string{
	hi := "hello,"
	return func(name string) string{
		res := hi + name
		return res
	}
}


func Greeting1(name string) func() string{
	hi := "hello,"
	return func() string{
		res := hi + name
		return res
	}
}

func main()  {




	g1 := Greeting()
	g2:= Greeting()

	fmt.Println(g1("alex"))
	fmt.Println(g1("李逵"))
	fmt.Println(g1("宋江"))

	fmt.Println(g2("宋江"))
	fmt.Println(g2("李逵"))

	// 错误，不能指定name了
	f1 := Greeting1("alex")
	fmt.Println(f1())

	/*
	hello,alex
	hello,李逵
	hello,宋江

	hello,宋江
	hello,李逵

	hello,alex
	*/
}