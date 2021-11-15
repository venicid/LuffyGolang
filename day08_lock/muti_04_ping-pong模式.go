/*
用1个int表示ball（球），channel表table
假设两个goroutine代表2位运动员，标号为1和2
*/

package main

import (
	"fmt"
	"time"
)

func player(id string, table chan int)  {
	for  {
		ball := <- table
		fmt.Printf("%s got ball[%d]\n", id, ball)
		time.Sleep(time.Second)
		fmt.Printf("%s bonceback ball[%d]\n", id, ball)
		ball ++
		table <- ball
	}
}

func main()  {
	var Ball int
	table := make(chan int)
	go player("1", table)
	go player("2", table)
	go player("3", table)

	table <- Ball  // 首先把ball放到table上
	time.Sleep(5*time.Second)  // 5秒后结束比赛
	<- table  // 取回球


}
