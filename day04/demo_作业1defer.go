package main

import "log"

func f5() (s1 []int) {
	s1 = []int{1, 1}
	defer func() {
		s1[1] = 10
	}()

	return []int{3, 3}
}


func f55() (s1 [] int) {
	s1 = []int{1,1}

	s1 = []int{3, 3}   // 先给s1赋值

	func() {   // defer改为正常函数
		s1[1] = 10  //闭包 s1 [3,10]
	}()

	return   // return 为空
}



func main()  {
	log.Println(f5())
	log.Println(f55())
}