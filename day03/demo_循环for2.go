package main

import (
	"log"
	"time"
)

func main()  {
	for{
		log.Println("SB,我干活呢")
		time.Sleep(3*time.Second)
	}


}

/*
2021/08/08 11:14:10 SB,我干活呢
2021/08/08 11:14:13 SB,我干活呢
2021/08/08 11:14:16 SB,我干活呢

*/