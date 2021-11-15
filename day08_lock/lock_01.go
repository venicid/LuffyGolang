package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	fmt.Printf("%+v\n", mutex)

	mutex.Lock()
	fmt.Printf("%+v\n", mutex)

	mutex.Unlock()
	fmt.Printf("%+v\n", mutex)
}

/*
{state:0 sema:0}
{state:1 sema:0}
{state:0 sema:0}
*/