package main

import "log"

func main()  {

	arr1 := []int{10,20,30,40}
	map1 := map[string] string{
		"k1":"v1",
		"k2":"v2",
		"k3":"v3",
		"k4":"v4",
	}


	for index := range arr1{
		log.Printf("[切片][range 遍历单变量 只有索引][index：%v]", index)
	}
	for index,val := range arr1{
		log.Printf("[切片][range 遍历双变量 索引和值][index：%v][value：%v]", index, val)
	}

	// 遍历单变量 只有索引
	for k:=range map1{
		log.Printf("[map][range 遍历单变量 只有key][key：%v]", k)
	}

	for k,v:= range map1{
		log.Printf("[map][range 遍历双变量 key,value][key：%v][value：%v]", k, v)
	}

}
