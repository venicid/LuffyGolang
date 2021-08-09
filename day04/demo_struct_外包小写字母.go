package main

import (
	"day04/tt"
	"fmt"
)

var t1 = tt.Test1{X:5}

// 小写字母，外包不可见
//var t2 = tt.Test2{x: 6}
//var t3 = tt.test1{X: 6}
//var t4 = tt.test2{x: 6}
func main()  {

	fmt.Println(t1)
	//fmt.Println(t2)
	//fmt.Println(t3)
	//fmt.Println(t4)
	/*
	   .\demo.go:9:19: cannot refer to unexported field 'x' in struct literal of type tt.Test2
	   .\demo.go:10:10: cannot refer to unexported name tt.test
	   .\demo.go:11:10: cannot refer to unexported name tt.test
	*/
}