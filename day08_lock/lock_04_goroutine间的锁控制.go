package main

import (
	"fmt"
	"sync"
	"time"
)

type MyStruct struct {
	v   int
	mux sync.Mutex
}

func (s *MyStruct) Lock() {
	s.mux.Lock()
}

func (s *MyStruct) Unlock() {
	s.mux.Unlock()
}

func main() {
	s := MyStruct{v: 0}
	s.v = 1
	fmt.Printf("%+v\n", s)

	go s.Lock()
	time.Sleep(1 * time.Second)
	fmt.Printf("%+v\n", s)

	go s.Unlock()
	time.Sleep(1 * time.Second)
	fmt.Printf("%+v\n", s)
}