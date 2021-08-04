package main

import "fmt"

func main(){

	a1 := []int{1,2,3,4}
	fmt.Printf("a1[len: %d, cap: %d]\n", len(a1), cap(a1))
	a2 := append(a1, 5)
	fmt.Printf("a2[len: %d, cap: %d]\n", len(a2), cap(a2))

	a3 := make([]int,1000)
	fmt.Printf("a3[len: %d, cap: %d]\n", len(a3), cap(a3))

	a4 := append(a2, a3...)
	fmt.Printf("a4[len: %d, cap: %d]\n", len(a4), cap(a4))

	a6 :=append(a4, a3...)
	fmt.Printf("a6[len: %d, cap: %d]\n", len(a6), cap(a6))
	/*
	a1[len: 4, cap: 4]
	a2[len: 5, cap: 8]
	a3[len: 1000, cap: 1000]
	a4[len: 1005, cap: 1024]
	a6[len: 2005, cap: 2560]
	*/
}