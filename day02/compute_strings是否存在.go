package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("hello world", "world"))
	fmt.Println(strings.ContainsAny("hello world", "l o"))
	fmt.Println(strings.ContainsRune("你好", '好'))

	/*
	true
	true
	true
	*/
}