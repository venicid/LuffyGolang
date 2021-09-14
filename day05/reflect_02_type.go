package main

import (
	"log"
	"reflect"
)

func main() {

	var name string = "xiaoyi"
	// TypeOf 会返回目标的类型
	reflectType := reflect.TypeOf(name)
	// 返回值
	reflectValueOf := reflect.ValueOf(name)
	log.Printf("[typeof :%v]", reflectType)
	log.Printf("[ValueOf :%v]", reflectValueOf)
}

/*
2021/09/14 23:38:37 [typeof :string]
2021/09/14 23:38:37 [ValueOf :xiaoyi]
*/