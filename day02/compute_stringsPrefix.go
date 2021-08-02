package main

import (
	"fmt"
	"strings"
)

func main() {

	// s1服务标识
	s1 := "inf.bigdata.kafka"

	// restful接口
	fmt.Println(strings.HasPrefix(s1, "inf"))
	fmt.Println(strings.HasSuffix(s1, "kafka"))
	fmt.Println(strings.HasSuffix(s1, ""))

	 /*
	 true
	 true
	 true
	 */
}