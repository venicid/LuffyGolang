package main

import (
	"fmt"
	"sort"
)

func main(){

	m1 := make(map[string]int)

	// 遍历赋值
	keys := make([]string, 0)
	for i:=0; i<10;i++{
		key := fmt.Sprintf("key_%d", i)
		m1[key] = i
		keys = append(keys, key)
	}
	fmt.Println(m1)


	for k := range m1{
		fmt.Printf("[key=%s]\n", k)
	}


	// range遍历取值
	fmt.Println("无序遍历")
	for k,v:= range m1{
		fmt.Printf("[%s=%d]\n", k,v)
	}

	// 有序key遍历
	fmt.Println("有序遍历")

	for _,key := range keys{
		fmt.Printf("[%s=%d]\n", key, m1[key])
	}

	// slice排序
	keys2 := []string{"c","a","b"}
	sort.Sort(sort.StringSlice(keys2))
	fmt.Println(keys2)

	/*
	[a b c]
	*/

}