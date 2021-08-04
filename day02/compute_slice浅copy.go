
package main

import "fmt"

func main(){

	a1 := []int{1,2,3}
	a2 := a1

	a1[1] = 10
	fmt.Println(a1, a2)
	a2[2] = 20
	fmt.Println(a1, a2)

	/*
		互相影响，浅copy
		[1 10 3] [1 10 3]
		[1 10 20] [1 10 20]
	*/
}