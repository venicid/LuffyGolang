package main

import (
	"fmt"
	"sync"
	"time"
)


type SafeCounter struct {
	v map[string] int // map并发不安全  -->  sync.Map
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string, id int)  {
	c.mux.Lock()   // 互斥锁
	c.v[key]++
	c.mux.Unlock()
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
0
9
7
2
1
3
4
5
6
8
*/