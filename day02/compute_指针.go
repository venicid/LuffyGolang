package main

import "fmt"

func main() {

	var s1 = "hello"
	var s1p = &s1

	fmt.Printf("[%T %v]\n", s1, s1)  // string, hello
	fmt.Printf("[%T %v]\n", s1p, s1p) // *string 0xdadff
	fmt.Printf("[%T %v]\n", *s1p, *s1p) // string hello

	/*
	[string hello]
	[*string 0xc000040240]
	[string hello]
	*/
}

