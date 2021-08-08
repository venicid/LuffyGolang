package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)


func main() {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	// 创建过期的30s缓存，每5秒清理一次
	c := cache.New(30*time.Second, 5*time.Second)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", 10*time.Second)

	res,ok := c.Get("foo")
	fmt.Println(res, ok)
	time.Sleep(10*time.Second)

	res,ok = c.Get("foo")
	fmt.Println(res, ok)

	/*
	   bar true
	   <nil> false

	*/
}