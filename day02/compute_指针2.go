package main

import "fmt"

func main() {

	num := 1
	fmt.Printf("[初始化后的值， %d]\n", num)
	add(num)
	fmt.Printf("[调用完add函数的值， %d]\n", num)

	/*
	[初始化后的值， 1]
	[传入add函数参数的值， 1]
	[add函数增加后的值， 2]
	[调用完add函数的值， 1]
	*/

}

func add(num int)  {
	fmt.Printf("[传入add函数参数的值， %d]\n", num)
	num ++
	fmt.Printf("[add函数增加后的值， %d]\n", num)
}