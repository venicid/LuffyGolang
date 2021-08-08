package main

import (
	"fmt"
	"log"
)

// 返回1
func f1() (res int) {
	defer func() {
		res++
	}()
	return 0
}

func f11() (res int) {
	res = 0
	func() {
		res++
	}()
	return

}

// 匿名不带参，返回res 5
func f2() (res int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// f2可以改造为
func f22() (res int) {
	t := 5
	res = t // 真正的res返回值 赋值
	func() {
		t = t + 5
	}()
	log.Printf("f22.t=%d",t)
	return
}

// 匿名带参数，返回res 0
func f3() (res int) {
	defer func(res int) {
		res = res + 5
	}(res)
	return res
}
func f33() (res int) {
	res = 0 // 赋值指令
	func(res int) {
		res = res + 5 //值传递，深拷贝， 操作的是副本
	}(res)

	return // 空的return
}

// f4返回 t 返回10
func f4() (t int) {

	defer func() {
		t = t * 10
	}()

	return 1
}

func f44() (t int) {
	t = 1 // 先给返回值赋值
	func() {
		t = t * 10 // 匿名函数，引用闭包的自由变量 t ,会直接更改t的值
	}()

	return
}

func main() {
	fmt.Println(f1())
	fmt.Println(f11())
	fmt.Println(f2())
	fmt.Println(f22())
	fmt.Println(f3())
	fmt.Println(f33())
	fmt.Println(f4())
	fmt.Println(f44())
}