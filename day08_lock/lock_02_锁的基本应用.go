package main

import (
	"fmt"
	"time"
)


type SafeCounter struct {
	v map[string] int // map并发不安全
}

func (c *SafeCounter) Inc(key string, id int)  {
	c.v[key]++
	fmt.Println(id)
}

func main()  {

	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i < 10; i++ {
		go c.Inc("key", i)
	}

	// main线程会提前关闭
	time.Sleep(time.Second)
}

/*
 map并发不安全

fatal error: concurrent map writes

goroutine 20 [running]:
runtime.throw(0x560008, 0x15)
*/