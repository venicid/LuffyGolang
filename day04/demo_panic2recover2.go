package main

import "fmt"

func defer_func()  {



	defer func() {
		fmt.Println("2")
	}()


	defer func() {
		fmt.Println("3")
	}()

	defer func() {
		if err := recover();err!=nil{
			fmt.Printf("被recover捕获的panic内容是：%v\n", err)
		}
		fmt.Println("1")
	}()

	panic("我是panic")
	defer func() {
		fmt.Println("4")
	}()
}

func main()  {
	defer_func()
}

/*
3
2
被recover捕获的panic内容是：我是panic
1
*/