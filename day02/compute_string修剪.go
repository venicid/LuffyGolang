package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	x := "@@@hello world@@@"

	fmt.Println(strings.Trim(x, "@"))
	fmt.Println(strings.TrimLeft(x, "@"))
	fmt.Println(strings.TrimRight(x, "@"))
	fmt.Println(strings.TrimSpace(x))
	fmt.Println(strings.TrimPrefix(x, "@@@"))
	fmt.Println(strings.TrimSuffix(x, "@@@"))

	f := func(r rune) bool {
		// r 如果是汉字返回true
		return unicode.Is(unicode.Han, r)
	}
	fmt.Println(strings.TrimFunc("你好，世界！ hello world", f))
	/*
	hello world
	hello world@@@
	@@@hello world
	@@@hello world@@@
	hello world@@@
	@@@hello world
	，世界！ hello world
	*/


	// 对比一下TrimLeft和TrimPrefix的区别
	x1 := "@@@hello@@@"
	fmt.Println("[TrimLeft: %v]", strings.TrimLeft(x1, "@"))
	fmt.Println("[TrimPrefix: %v]", strings.TrimPrefix(x, "@"))

	/*
	[TrimLeft: %v] hello@@@
	[TrimPrefix: %v] @@hello world@@@
	*/
}