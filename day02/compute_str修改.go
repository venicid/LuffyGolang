package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s1 := "localhost:8080"
	s11 := "localhost:8080你"
	fmt.Println(s1)
	fmt.Println(s11)

	// 强制类型转换 string to byte
	strByte := []byte(s1)
	// string to rune
	strByte1 := []rune(s11)

	// 下标修改
	strByte[len(s1)-1] = '1'
	fmt.Println("[ASCII]",strByte)

	strByte1[utf8.RuneCountInString(s11)-1] = '1'
	fmt.Println("[ASCII]", strByte1)

	// 强制类型转换 []byte to string
	s2 := string(strByte)
	fmt.Println(s2)
	// []rune to string
	s21 := string(strByte1)
	fmt.Println(s21)

	/*
	localhost:8080
	localhost:8080你
	[ASCII] [108 111 99 97 108 104 111 115 116 58 56 48 56 49]
	[ASCII] [108 111 99 97 108 104 111 115 116 58 56 48 56 48 49]
	localhost:8081
	localhost:80801
	*/
}
