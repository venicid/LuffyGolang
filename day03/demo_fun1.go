package main

import "log"

func max(n1 int, n2 int) int{

	if n1>n2{
		return n1
	}
	return n2
}

func main() {
	res := max(34,4)
	log.Println(res)

	res1 := max(33,111)
	log.Println(res1)

}

/*
2021/08/08 15:07:52 34
2021/08/08 15:07:52 111

*/

