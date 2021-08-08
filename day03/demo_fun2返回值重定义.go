
package main

import "log"

func f1() (names []string, m map[string]int, num int) {
	m = make(map[string]int)
	m["k1"] = 2
	return
}
func main() {
	a,b,c :=f1()
	log.Println(a,b,c)
}

/*
2021/08/08 15:10:20 [] map[k1:2] 0
*/