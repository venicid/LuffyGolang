package main

import "log"

func main()  {

	// 第1种 三段全的"
	for i:=0;i<8;i++{
		log.Printf("老子干了%v个小时了", i)
	}

	// 第2种 自增写在里面
	for i:=0; i<8;{
		log.Printf("老子干了%v个小时了", i)
		i ++
	}

	// 第3种 自增写在里面 ，初始化写在上面
	var j int
	for ;j<8;{
		log.Printf("老子干了%v个小时了", j)
		j ++
	}

}

