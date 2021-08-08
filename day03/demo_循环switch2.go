package main

import (
	"fmt"
	"log"
)

func lang(lang string) string {
	var res string

	love := "golang"
	python2 := "python"
	python3 := "python"

	switch lang {

	// case允许两个常量
	case "java", "go":
		res = "go"

	// case不允许两个常量相同
	//case "c++", "c++":
	case "c++", "c":
		res =  "c"

	case love:
		res = love

	//  fallthrough,相当于继续往下走
	case python2:
		fmt.Println("python")
		fallthrough

	case python3:
		fmt.Println("python")
		fallthrough

	default:
		res = "unknow"
	}

	return res
}

func main()  {

	res := lang("fat")
	log.Println(res)

	res1 := lang("go")
	log.Println(res1)

	res3 := lang("python")
	log.Println(res3)
}

/*
2021/08/08 11:09:10 unknow
2021/08/08 11:09:10 go
python
python
2021/08/08 11:09:10 unknow
*/


