package main

import "fmt"

// 类型断言和判断

func main()  {

	var s interface{} = "abc"
	s1, ok := s.(string)
	fmt.Println(s1, ok)
	s2, ok := s.(int)
	fmt.Println(s2, ok)

	var i interface{} = 123
	i1, ok := i.(int)
	fmt.Println(i1, ok)
	i2, ok := i.(string)
	fmt.Println(i2, ok)


	// s.(type) 判断类型
	var f interface{} = false

	switch f.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("未知的值")
	}

}

/*
	abc true
	0 false
	123 true
	 false
	未知的值
*/