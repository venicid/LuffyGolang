/*
   Servers中的worker线程是按需生成的，并且工作处理完毕后就释放，
   在主进程中，有一个for循环， Accept函数一直去阻塞着循环的进行，一旦有新的请求过来，Accept就会生成一个连接，
然后呢主进程就创建一个子进程处理这个连接以及其他逻辑
 */

package main

import (
	"fmt"
	"net"
	"time"
)

func handler(c net.Conn , ch chan int)  {
	ch <- len(c.RemoteAddr().String())
	time.Sleep(10 * time.Millisecond)
	c.Write([]byte("ok"))
	c.Close()
}

func pool(ch chan int, n int) {
	wch := make(chan int)

	for i := 0; i <n; i++{
		go logger(wch)
	}
	for {
		wch <- <-ch
	}
}

func logger(ch chan int) {
	for {
		time.Sleep(1500 * time.Millisecond)
		fmt.Println(<-ch)
	}
}

func server(l net.Listener, ch chan int) {
	for {
		c, err := l.Accept()
		if err !=nil {
			continue
		}
		go handler(c, ch)
	}
}

func main(){
	l, err := net.Listen("tcp", ":5000" )
	if err != nil {
		panic(err)
	}
	ch := make(chan int)
	// go logger(ch)
	go pool(ch, 36)
	go server(l, ch)
	time.Sleep(10 * time.Second)
}


