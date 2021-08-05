package main

import "fmt"

// 只声明
var m1 map[string]string

// 声明并初始化
var m2 = map[string]string{"a":"b"}

func main(){

	m:=make(map[string]string)

	// 增加，如何批量增加呢？
	m["app"] = "taobao"
	m["lang"] = "golang"

	// 删除
	delete(m, "app")

	// 改
	m["lang"] = "python"

	// 查
	// 单变量形式
	lang := m["lang"]
	fmt.Println(lang)
	// 双变量形式
	lang, exists := m["lang"]
	if exists{
		fmt.Printf("[lang存在，值：%v]\n", lang)
	}
	if !exists{
		fmt.Printf("[lang不存在]\n")
		m["lang"] = "vue"
	}
	fmt.Println(m)





}