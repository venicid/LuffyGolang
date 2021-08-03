package main

import "fmt"

func main(){

	arr1 := [4][2]int{{10,11},{20,21},{30,31},{40,41}}

	fmt.Println(arr1)
	fmt.Println(arr1[2][1])

	/*
		[
		[10 11]
		[20 21]
		[30 31]  // arr1[2][1]
		[40 41]
		]
	*/

}