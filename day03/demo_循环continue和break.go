package main

import (
	"fmt"
	"log"
	"strings"
)

func main()  {

	m1 := make(map[string] string)

	for i:=0;i<20;i++{
		key := fmt.Sprintf("key_%v", i)
		value := fmt.Sprintf("value_%v", i)
		m1[key] = value
	}

	for k,v:=range m1{
		if strings.HasSuffix(v,"3"){
			log.Println("value存在3就continue")
			continue
		}
		if k =="key_18"{
			log.Println("遇到8就退出")
			break
		}

		log.Println(k,v)
	}


/*
   2021/08/08 11:49:19 key_2 value_2
   2021/08/08 11:49:19 value存在3就continue
   2021/08/08 11:49:19 key_7 value_7
   2021/08/08 11:49:19 key_19 value_19
   2021/08/08 11:49:19 遇到8就退出
*/


}