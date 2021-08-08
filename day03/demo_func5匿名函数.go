package main

import "fmt"

func main()  {

	// 不带参
	f := func() {
		fmt.Println("abc")
	}
	f()
	fmt.Printf("%T\n", f)

	// 带参数的
	f1 := func(a string) {
		fmt.Println(a)
	}
	f1("abcd")
	fmt.Printf("%T\n", f1)

	// 带返回值的
	f2 := func() string{
		return "abc"
	}

	res1 := f2()
	fmt.Println(res1)

}

/*
abc
func()

abcd
func(string)

abc


*/