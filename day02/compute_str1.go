package main

import "fmt"

func main() {

	//var ch = 'abc'
	// .\compute.go:5:11: more than one character in rune literal
	var ch1 = 'a'
	fmt.Println(ch1)

	str1 := "你好呀，\t世界"
	fmt.Println(str1)

	str2 := `你好呀，\t世界`
	fmt.Println(str2)

	jsonStr := `
		{
		"region":"bj",
		"ids":[1,2,3]
		}
		`
	fmt.Println(jsonStr)

	// .\compute.go:26:39: syntax error: unexpected literal 2 at end of statement
	//promSql1 := "sum(rate(api_qps{code=~"2xx"}[1m])) * 100"
	//fmt.Println(promSql1)

	promSql2 := `sum(rate(api_qps{code=~"2xx"}[1m])) * 100`
	fmt.Println(promSql2)
}

