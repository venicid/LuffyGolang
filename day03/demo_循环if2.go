package main

import "fmt"

func main()  {
	m := map[string]string{
		"region": "bj",
		"idc":"世纪互联",
	}
	idc := m["idc"]
	if idc == "世纪互联"{
		fmt.Printf("地区： %v，机房：%v\n",idc,  m["region"])
	}

	if idc := m["idc"]; idc =="世纪互联"{
		fmt.Printf("地区： %v，机房：%v\n",idc,  m["region"])
	}

}


/*
地区： 世纪互联，机房：bj
地区： 世纪互联，机房：bj
*/