package main

import "fmt"

func main()  {

	var a = 21
	var c int

	c = a
	fmt.Println(c, a)
	c += a
	fmt.Println(c, a)
	c *= a
	fmt.Println(c, a)

	/**
	21 21
	42 21
	882 21
	**/
}