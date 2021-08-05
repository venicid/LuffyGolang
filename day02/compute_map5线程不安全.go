package main

import "time"

func main(){

	c := make(map[int]int)

	// 匿名goroutine 循环写map
	go func() {
		for i:=0;i<10000;i++{
			c[i] = i
		}
	}()

	// 匿名goroutine 循环写map
	go func() {
		for i:=0;i<10000;i++{
			c[i] = i
		}
	}()
	

	time.Sleep(40*time.Minute)

	/*
		fatal error: concurrent map writes

		goroutine 5 [running]:
	*/

}