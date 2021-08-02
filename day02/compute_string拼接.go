package main

import (
	"fmt"
	"strings"
)

func main() {

	ss := []string{
		"A",
		"说",
		"我要",
		"升职加薪",
	}

	var b strings.Builder
	for _,s := range ss{
		b.WriteString(s)
	}

	fmt.Println(b)
	fmt.Println(b.String())
	/*
	{0xc000153ef8 [65 232 175 180 230 136 145 232 166 129 229 141 135 232 129 140 229 138 160 232 150 170]}
	A说我要升职加薪
	*/
}