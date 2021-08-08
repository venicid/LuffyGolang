package main

import (
	"fmt"
)

// if
func main()  {

	env := "fat"
	currEnv := "all"
	if env == currEnv{
		fmt.Println("环境选择相同，请继续")
	}else {
		fmt.Println("该环境不支持，请重新选择")
	}

}