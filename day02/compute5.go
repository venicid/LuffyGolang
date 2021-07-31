
package main

import "fmt"

func main()  {

	var a int = 4
	var ptr *int
	fmt.Printf("[a变量类型为 %T]\n", a)
	fmt.Printf("[ptr变量类型为 %T]\n",ptr)


	ptr = &a
	fmt.Printf("[a变量的值为 %d]\n", a)
	fmt.Printf("[*ptr变量的值为 %d]\n", *ptr)
	fmt.Printf("[ptr变量的值为 %v]\n", ptr)
	fmt.Printf("[a的指针地址为 %p]\n", &a)

	/*
		[a变量类型为 int]
		[ptr变量类型为 *int]
		[a变量的值为 4]
		[*ptr变量的值为 4]
		[ptr变量的值为 0xc0000ac058]
		[a的指针地址为 0xc0000ac058]
	*/

}