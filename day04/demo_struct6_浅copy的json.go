package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Name     string
	Age      int
	Tags     map[string]string
	HouseId1 [2]int //数组是值类型
	HouseId2 []int  // 切片是引用类型
}

func main() {
	p1 := Person{
		Name:     "小乙",
		Age:      123,
		Tags:     map[string]string{"k1": "v1", "k2": "v2"},
		HouseId1: [2]int{100, 101},
		HouseId2: []int{200, 201},
	}

	var p2 Person

	data, _ := json.Marshal(p1)
	json.Unmarshal(data, &p2)

	// 修改两个值类型的字段
	p1.Age = 19
	p2.Name = "898"

	// 修改map  浅copy
	p1.Tags["k1"] = "v11"
	// 修改array 深copy
	p2.HouseId1[0] = 300
	// 修改切片  浅copy
	p1.HouseId2[1] = 301

	log.Printf("[p1的内存地址:%p ][value:%+v]", &p1, p1)
	log.Printf("[p2的内存地址:%p ][value:%+v]", &p2, p2)
}

/*
2021/08/10 08:14:12 [p1的内存地址:0xc00004e0a0 ][value:{Name:小乙 Age:19 Tags:map[k1:v11 k2:v2] HouseId1:[100 101] HouseId2:[200 301]}]
2021/08/10 08:14:12 [p2的内存地址:0xc00004e0f0 ][value:{Name:898 Age:123 Tags:map[k1:v1 k2:v2] HouseId1:[300 101] HouseId2:[200 201]}]
*/