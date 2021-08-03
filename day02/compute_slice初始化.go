package main

import "fmt"

func main(){

	var s1 []int
	fmt.Println(s1)

	s1 = append(s1, 1)
	s1 = append(s1, 2)
	fmt.Println(s1)

	var s2 = []int{12,34,5}
	fmt.Println(s2)

	/*
	[]
	[1 2]
	[12 34 5]
	*/
}