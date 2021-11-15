package main

/*

Fanout，扇出模式，就是工作模式就一个管道分发任务，然后由多个goroutines去执行
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ch <- chan int, wg *sync.WaitGroup) {
	for {
		task, ok := <-ch
		if !ok {
			return
		}
		time.Sleep(20 * time.Millisecond)
		fmt.Println("启动task", task)
	}
	defer wg.Done()
}

func pool(wg *sync.WaitGroup, workers, tasks int){
	ch := make(chan int)

	for i :=0; i < workers; i++ {
		time.Sleep(1 * time.Millisecond)
		// spaw出很多worker线程
		go worker(ch, wg)
	}

	for i:=0; i <tasks; i++{
		time.Sleep(10 * time.Millisecond)
		// 开始分发任务，被激活的workers开始工作
		ch <-i
	}
	close(ch)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36,36)
	wg.Wait()
}


 */