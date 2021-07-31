package main

import "fmt"


const (
	a = iota
	b
	c
	d
)

type MyType int

const (
	T1 MyType = iota
	T2
	T3
	T4
)

const (
	T11 MyType = iota
	T12
	_
	_
	T13
	T14
)

const (
	A = iota
	B = 5
	C
	D = iota
	E
	F
	G
)

const(
	A1 MyType = 1 << iota
	A2
	A3
	A4
)

const (
	_ = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	)


func main()  {

	fmt.Println(a,b,c,d)  // 1,2,3,4
	fmt.Println(T1, T2, T3, T4)  // 0 1 2 3
	fmt.Println(T11, T12, T13, T14)  // 0 1 4 5

	fmt.Println(A,B,C,D,E,F,G)  // 0 5 5 3 4 5 6
	fmt.Println(A1, A2,A3, A4)  // 1 2 4 8

	fmt.Println(1<< 10)  // 1024
	fmt.Println(KB,MB,GB,TB)  // 1024 1.048576e+06 1.073741824e+09 1.099511627776e+12

}
