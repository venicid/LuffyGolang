/*
daisy chain 去创建很多个channel
然后这些channel首尾相连起来，组成单项链表
每个channel都在处理不同的子任务
*/

package main

import "fmt"

func f(left chan int, right chan int)  {
	// 这个函数把right的输出和left的输入联系起来
	left<- 1+ <-right
}

func main()  {
	const N = 10000
	ch := make(chan int)
	left := ch
	right := ch

	// 创建长度为n的菊花链
	for i := 0; i < N; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	// 在链的最右端输入1,那么最左端就会得到10001
	go func(c chan int) {
		c <- 1
	}(right)

	fmt.Println(<-ch)

}
