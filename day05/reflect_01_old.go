package main

import "log"

func main()  {

	var s interface{} = "abc"
	switch s.(type) {
	case string:
		log.Println("是个string")
	case int:
		log.Println("是个string")
	case bool:
		log.Println("是个bool")
	}
}
   

/*
2021/09/14 23:33:03 是个string
*/
