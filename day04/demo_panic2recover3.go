package main

import "fmt"

func main()  {

	defer func() {
		if err := recover(); err !=nil{
			fmt.Println("recover")  // 捕获recover，只捕获最后1个
			fmt.Println(err)
		}else{
			fmt.Println("哈哈")
		}
	}()

	defer func() {
		panic("defer内部触发的panic")   // 覆盖掉main中的异常 panic
	}()

	panic("defer触发的panic")
}

/*
recover
defer内部触发的panic

*/