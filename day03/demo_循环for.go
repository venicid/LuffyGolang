package main

import (
	"log"
	"time"
)

func main()  {
	count := 8
	hours := 0
	for count>hours{
		log.Println("SB,我上班呢")
		hours += 1
		time.Sleep(1*time.Second)
	}
	log.Println("SB，老子下班了")
}

/*
2021/08/08 11:13:01 SB,我上班呢
2021/08/08 11:13:02 SB,我上班呢
2021/08/08 11:13:03 SB,我上班呢
2021/08/08 11:13:04 SB,我上班呢
2021/08/08 11:13:05 SB,我上班呢
2021/08/08 11:13:06 SB,我上班呢
2021/08/08 11:13:07 SB,我上班呢
2021/08/08 11:13:08 SB,我上班呢
2021/08/08 11:13:09 SB，老子下班了

*/