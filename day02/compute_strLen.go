package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	ch1 := "luffy"
	ch2 := "路飞教育"
	ch3 := "lu飞a育"

	fmt.Printf("[字符串：%v，字节大小or字符个数：%d，真实字符个数：%d]\n", ch1,len(ch1), utf8.RuneCountInString(ch1))
	fmt.Printf("[字符串：%v，字节大小or字符个数：%d，真实字符个数：%d]\n", ch2,len(ch2), utf8.RuneCountInString(ch2))
	fmt.Printf("[字符串：%v，字节大小or字符个数：%d，真实字符个数：%d]\n", ch3,len(ch3), utf8.RuneCountInString(ch3))

	/*
	[字符串：luffy，字节大小or字符个数：5，真实字符个数：5]
	[字符串：路飞教育，字节大小or字符个数：12，真实字符个数：4]
	[字符串：lu飞a育，字节大小or字符个数：9，真实字符个数：5]
	*/
}

