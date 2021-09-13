package main

import (
	"log"
	myLog "lugo04/src/log"
	. "lugo04/src/pk1"
	"lugo04/src/pk1/pk2"
)

func main() {
	myLog.LogPrint()
	log.Printf("官方的log")
	//pk4.Func_test1()
	Func_test1()
	pk3.Func_test3()
}
