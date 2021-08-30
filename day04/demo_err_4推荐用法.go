package main

import (
	"errors"
	"fmt"
)

func main() {

	err := errors.New("原始错误")
	myErr := fmt.Errorf("自定义错误:%w", err)
	fmt.Println(myErr)
}

/*
自定义错误:原始错误
*/

