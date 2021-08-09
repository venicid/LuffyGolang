package main

import "fmt"

func mapMerge(m1 ...map[string]string) map[string]string {
	//m1 就是map的切片
	if len(m1) == 0 {
		return map[string]string{}
	}
	mRes := make(map[string]string)
	for _, m := range m1 {
		for k, v := range m {
			mRes[k] = v
		}

	}
	return mRes

}
func main() {

	m1 := map[string]string{
		"k1": "v1",
		"k2": "v2",
		"k3": "v3",
	}

	m2 := map[string]string{
		"k1": "v11",
		"k2": "v22",
		"k3": "v33",
	}
	m3 := map[string]string{
		"k11": "v11",
		"k22": "v22",
		"k33": "v33",
	}
	// 第一种 直接塞各个参数
	fmt.Println(mapMerge(m1, m2, m3))

	// 第二种 slice传进去
	s1 := make([]map[string]string, 0)
	s1 = append(s1, m1)
	s1 = append(s1, m2)
	fmt.Println(mapMerge(s1...))

}

/*
map[k1:v11 k11:v11 k2:v22 k22:v22 k3:v33 k33:v33]
map[k1:v11 k2:v22 k3:v33]
*/