package main

import (
	myLog "day05/src/log"
	. "day05/src/pk1"
	"day05/src/pk1/pk2"
	"log"
)

func main() {
	myLog.LogPrint()
	log.Printf("官方的log")
	//pk4.Func_test1()
	Func_test1()
	pk3.Func_test3()
}

/*
我自定义的log
func_test1
func_test2
func_test3
func_test3
2021/09/16 22:36:01 官方的log
*/