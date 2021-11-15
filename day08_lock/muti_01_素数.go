/*
素数过滤器
*/
package main

import "fmt"

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2;; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(in <- chan int, prime int) chan int{
	out := make(chan int)
	go func() {
		for{
			if i:= <-in; i%prime !=0{
				out <- i
			}
		}
	}()
	return out
}

func main()  {


	ch := GenerateNatural()  // 自然序列: 2,3,4,5  // 生产者
	//fmt.Printf("%d\n", ch)

	for i := 0; i < 100; i++ {
		prime := <- ch
		fmt.Printf("%v:%v\n", i+1, prime)
		ch = PrimeFilter(ch, prime)   // 消费者  生产者和消费者的关系 1:n
		fmt.Printf("%d\n", ch)
	}

}