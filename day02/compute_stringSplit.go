package main

import (
	"fmt"
	"strings"
)

func main() {

	// 服务标识
	s1 := "inf.bigdata.kafka"
	s2 := "localhost:8080/api/v1/host/1"

	ss1:= strings.Split(s1, ".")
	ss2:= strings.SplitAfter(s1, ".")

	ps1 := strings.Split(s2, "/")
	ps2 := strings.SplitN(s2, "/", 2)

	fmt.Printf("[切割服务标识][]%v\n", ss1)
	fmt.Printf("[切割服务标识][SplitAfter]%v\n", ss2)
	fmt.Printf("[切割url][]%v\n", ps1)
	fmt.Printf("[切割url][SplitN]%v\n", ps2)

/*
   [切割服务标识][][inf bigdata kafka]
   [切割服务标识][SplitAfter][inf. bigdata. kafka]
   [切割url][][localhost:8080 api v1 host 1]
   [切割url][SplitN][localhost:8080 api/v1/host/1]
*/
}