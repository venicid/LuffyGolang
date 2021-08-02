
package main

import "fmt"

func main() {

	var f1,f2 float32
	var f3  float64
	f1 = 10000018
	f2 = 100000018
	f3 =  100000018
	f11 := f1 + 1
	f21 := f2 + 1
	f31 := f3 + 1

	fmt.Printf("[f1:%v,%T]\n", f1, f1)
	fmt.Printf("[f1:%v,%T]\n", f11, f11)
	fmt.Printf("[f1:%v,%T]\n", f2, f2)
	fmt.Printf("[f1:%v,%T]\n", f2, f21)
	fmt.Printf("[f1:%v,%T]\n", f3, f3)
	fmt.Printf("[f1:%v,%T]\n", f31, f31)

	fmt.Println(f1 == f11)  // false
	// f2 == f12 代表f2刚好达到了float32的精度上限
	fmt.Println(f2 == f21)  // true
	fmt.Println(f3 == f31)  // false

}