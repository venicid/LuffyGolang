package main

import "fmt"

var A int
var B int = 10

//cc := 10   .\compute.go:8:1: syntax error: non-declaration statement outside function body
var cc = 10

func main()  {

	var a int
	var b string
	var c bool
	var d = 4
	fmt.Println(a,b,c,d)

	/*
	0  false 4

	*/

}
