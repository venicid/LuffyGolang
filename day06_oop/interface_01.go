package main

import "fmt"

type MyStruct struct {
	X int
	Y int
}

func (a *MyStruct) add()  int{
	return a.X + a.Y
}

type Adder interface {
	add() int
}


func main()  {

		ms := MyStruct{
			X: 3,
			Y: 4,
		}
		fmt.Println(ms.add())

		var f Adder
		f = &ms
		fmt.Println(f.add())
}
