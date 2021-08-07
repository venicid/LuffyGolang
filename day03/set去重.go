package main

import "fmt"

// 切片元素的去重
// 使用map的key
func main1(){
	s1 := []string{"abc", "def", "abc", "ok", "hello"}

	m := make(map[string]int)

	for i:= range s1{
		item := s1[i]
		_, exists := m[item]
		if exists{
			m[item] += 1
		}else{
			m[item] = 1
		}
	}
	fmt.Println(m)

	var res []string
	for k,_ := range m{
		res = append(res, k)
	}
	fmt.Println(res)
}

func main(){
	s1 := []string{"abc", "def", "abc", "ok", "hello"}

	m := make(map[string]struct{})

	for _,i:= range s1{
		m[i] = struct{}{}
	}
	fmt.Println(m)

	res := make([]string, 0)
	for k,_ := range m{
		res = append(res, k)
	}
	fmt.Println(res)

	/*
		map[abc:{} def:{} hello:{} ok:{}]
		[hello abc def ok]
	*/
}