package main

import "fmt"


type ByteSlice []byte

func (bs ByteSlice) Append(data []byte) []byte  {
	c:= append(bs, data...)
	return c

}

func main()  {

	c := ByteSlice{}
	c.Append([]byte("11"))
	a1 := c.Append([]byte("12"))
	fmt.Println(a1)

	/*
	[49 50]
	*/


	// 空接口实现复合类型元素添加
	var s [] interface{}

	s1 := append(s, "data")
	s2 := append(s1, 111)
	fmt.Println(s1)  // [data]
	fmt.Println(s2)  // [data 111]
}
