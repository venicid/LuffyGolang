package main

import (
	"day04/tt"
	"fmt"
)
// 外部包 结构体名和字段全大写
var a = tt.Test{X: 5}

// 同包 结构体名和字段可以小写
var b = Test2{x: 5}
var c = test3{x: 5}

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
