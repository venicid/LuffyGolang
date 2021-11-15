/*
实现1个1000个自增整数的生成器

生成器模式
*/
package main

import "fmt"

func xrange() chan int {  // xrange用来生成自增的整数
	ch := make(chan int)
	go func() {
		for i := 0;; i++ {
			ch<-i  // 直到通道索要数据时，才把i添加进通道
		}
	}()
	return ch
}

func main()  {
	generator := xrange()

	for i := 0; i < 1000; i++ {  // 生成1000个自增的整数
		fmt.Println(<-generator)
	}

}
