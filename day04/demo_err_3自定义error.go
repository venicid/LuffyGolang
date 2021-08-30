package main

import (
	"errors"
	"log"
)

type MyError struct {
	abc error
	msg string
}

func (myError *MyError) Error() string  {
	return myError.abc.Error() + myError.msg
}

func main() {
	theError := errors.New("原始错误")
	tempError := MyError{
		abc: theError,
		msg: "自定义错误",
	}

	errMsg := tempError.Error()
	log.Println(errMsg)
}

/*
2021/08/30 23:50:09 原始错误自定义错误

Error()源码
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
*/