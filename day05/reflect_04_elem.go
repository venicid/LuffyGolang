package main

import (
"log"
"reflect"
)

func main() {


	var num = 3.14
	log.Printf("[num 原始值:%v]", num)

	typeNum := reflect.TypeOf(num)
	log.Println(typeNum)

	//赋值
	// 通过reflect.ValueOf获取num的value
	pointer := reflect.ValueOf(&num)
	log.Println(pointer)
	newValue := pointer.Elem()

	// 赋值必须是指针，不然会panic
	// panic: reflect: call of reflect.Value.Elem on float64 Value
	//pointer = reflect.ValueOf(num)
	//newValue = pointer.Elem()

	// 类型要对上
	newValue.SetFloat(10.1)
	log.Printf("[num 新值:%v]", num)
	// call of reflect.Value.SetInt on float64 Value


}

/*
2021/09/15 01:15:22 [num 原始值:3.14]
2021/09/15 01:15:22 float64
2021/09/15 01:15:22 0xc00000a098
2021/09/15 01:15:22 [num 新值:10.1]

*/