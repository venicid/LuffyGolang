package main

import "fmt"

func main()  {

	defer func() {
		fmt.Println(1)
	}()

	defer func() {
		// 捕获异常
		if info := recover(); info !=nil{
			fmt.Println("触发了宕机", info)
		}else{
			fmt.Println("程序运行正常")
		}
	}()

	fmt.Println(2)
	fmt.Println(3)

	panic("程序错误，hahahaha")
	fmt.Println(4)

	defer func() {
		fmt.Println(5)
	}()

}

/*
2
3
触发了宕机 程序错误，hahahaha
1
*/