package main

import "fmt"

func main() {

	var ch1 byte = 'a'
	var ch2 = 'a'
	var ch3 = '你'

	fmt.Printf("[字符chr1 指定byte类型 指定ASCII码：%c, id:%v, 实际类型：%T]\n", ch1, ch1, ch1)
	fmt.Printf("[字符chr2 没有显示指定byte类型 默认UTF-8编码：%c, id:%v, 实际类型：%T]\n", ch2, ch2, ch2)
	fmt.Printf("[字符chr3 中文:%c, id:%v, 实际类型：%T]\n", ch3, ch3, ch3)

	var ch4 = "你"  // 单引号与双引号不一致~ 坑
	fmt.Printf("[字符ch4 中文:%c, id:%v, 实际类型：%T]\n", ch4, ch4, ch4)

	/*
	[字符chr1 指定byte类型 指定ASCII码：a, id:97, 实际类型：uint8]
	[字符chr2 没有显示指定byte类型 默认UTF-8编码：a, id:97, 实际类型：int32]
	[字符chr3 中文:你, id:20320, 实际类型：int32]
	[字符ch4 中文:%!c(string=你), id:你, 实际类型：string]
	*/
}

