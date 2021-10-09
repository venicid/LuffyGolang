# Go语言并发

### Agenda

- 并发基础
- Goroutine编程
- Go channel 
- channel通信
- channel select case的妙用
- 多路复用
- 综合实践&作业

## Go语言并发编程基础

进程/线程
进程是程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
线程是进程的一个执行实体，是 CPU 调度和分派的基本单位，它是比进程更小的能独立运行的基本单位。
一个进程可以创建和撤销多个线程，同一个进程中的多个线程之间可以并发执行。

并发/并行
`多线程程序在单核心的 cpu 上运行，称为并发`；`多线程程序在多核心的 cpu 上运行，称为并行`。
并发与并行并不相同，并发主要由切换时间片来实现“同时”运行，并行则是直接利用多核实现多线程的运行，Go程序可以设置使用核心数，以发挥多核计算机的能力。

协程/线程
协程：独立的栈空间，共享堆空间，调度由用户自己控制，本质上有点类似于用户级线程，这些用户级线程的调度也是自己实现的。
线程：一个线程上可以跑多个协程，协程是轻量级的线程。

## 同步&异步 
在计算机里，一切事物都是从同步阻塞开始，慢慢发展成异步非阻塞。
只要不想阻塞，就有异步的需求，因此就要实现并发，在计算机就由多进程、多线程来具体实现。

## 阻塞&非阻塞
阻塞与非阻塞是对同一个线程来说的，在某个时刻，线程要么处于阻塞，要么处于非阻塞。
阻塞是使用同步机制的结果，非阻塞则是使用异步机制的结果。

## 同步异步和阻塞非阻塞的组合
同步、异步，与阻塞、非阻塞，没有必然联系
组合例子：

- 同步阻塞: 打电话时候，其他任何事都做不了（比如不能同时看孩子)。效率低下 
- 同步非阻塞: 打电话时候，时不时看看孩子有没有危险。不停切换，效率低下 
- 异步阻塞: 发短信，在收到短信回复之前什么事都做不了（比如不能同时看孩子）。这种情形不常见 
- 异步非阻塞: 发短信，然后陪孩子玩，收到短信回复通知后再去看手机。效率最高 
编码应用，指的都是用户代码逻辑，而非底层的实现：

- 同步阻塞：最普遍最常见，符合思维，容易编写，容易调试 
- 同步非阻塞：自己的逻辑是同步，但由于使用了某个模块，而这个模块是实现异步的，所以立刻返回，于是从自己的代码逻辑角度看没有阻塞。

## cpu密集型和io密集型

- io密集型: `包含网络io和磁盘io等`，就是一个执行很久的操作不是由本机cpu进行操作而是发送一个请求然后由远程服务器操作，或者是发送一个读写磁盘的请求然后由磁盘硬件DMA进行操作。总之就是cpu发出一个请求，然后等待结果。如果一个业务的这种类型的操作比较多，就称这个业务为io密集型业务 
- cpu密集型: 一个操作由本机cpu进行运算。如果`一个业务的cpu运算量比较大`，而io操作比较少，则称这个业务为cpu密集型业务。


多核处理器可以实现cpu运算的并行，提高运算能力，这个时候用多进程或者多线程就适用于cpu密集型。而协程不适合，因为协程是用户态的技术，只能单核，因此并不能提高cpu并行运算能力，也因此只适合io密集型。多进程和多线程除了在多核时候适应cpu密集型，也可以适用io密集型，尤其是单核时候，只不过协程的切换开销比进程和线程都小。


# Go 并发

 goroutine 是一种非常轻量级的实现，可在单个进程里执行成千上万的并发任务，它是Go语言并发设计的核心。说到底 goroutine 其实就是线程，但是它比线程更小，十几个 goroutine 可能体现在底层就是五六个线程。

使用 go 关键字就可以创建 goroutine，将 go 声明放到一个需调用的函数之前，在相同地址空间调用运行这个函数，这样该函数执行时便会作为一个独立的并发线程，这种线程在Go语言中则被称为 goroutine。

goroutine 的用法如下：

```go
//go 关键字放在方法调用前新建一个 goroutine 并执行方法体
go GetThingDone(param1, param2);

//新建一个匿名方法并执行
go func(param1, param2) {
}(val1, val2)

//直接新建一个 goroutine 并在 goroutine 中执行代码块
go {
    //do someting...
}
```
并发有什么用？价值何在？下面先看以下几种场景。

1. 一方面我们需要灵敏响应的图形用户界面，一方面程序还需要执行大量的运算或者 IO 密集操作，而我们需要让界面响应与运算同时执行。

2. 当我们的 Web 服务器面对大量用户请求时，需要有更多的“Web 服务器工作单元”来分别响应用户。

3. 计算机的 CPU 从单内核（core）向多内核发展，而我们的程序都是串行的，计算机硬件的能力没有得到发挥。

4.  我们的程序因为 IO 操作被阻塞，整个程序处于停滞状态，其他 IO 无关的任务无法执行。

所以：

- 并发可以充分利用 CPU 核心的优势，提高程序的执行效率；
- 并发能充分利用 CPU 与其他硬件设备固有的异步性。

>> 所有 goroutine 在 main() 函数结束时会一同结束。



## Goroutine之间的通信
并发编程的难度在于协调，而协调就要通过交流，从这个角度看来，并发单元间的通信是最大的问题。

在工程上，`有两种最常见的并发通信模型：共享数据和消息。`

共享数据是指多个并发单元分别保持对同一个数据的引用，实现对该数据的共享。被共享的数据可能有多种形式，比如内存数据块、磁盘文件、网络数据等。在实际工程应用中最常见的无疑是内存了，也就是常说的共享内存。

## Go Channel

`channel 是Go语言在语言级别提供的 goroutine 间的通信方式`。我们可以使用 channel 在两个或多个 goroutine 之间传递消息。

一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定。如果对 Unix 管道有所了解的话，就不难理解 channel，可以将其认为是一种类型安全的管道。

定义一个 channel 时，也需要定义发送到 channel 的值的类型，注意，必须使用 make 创建 channel，代码如下所示：

```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```
通道的发送/接收都是使用特殊的操作符“<-”，如下所示：

```
// 创建通道
ch := make(chan int)
// 向通道发送数据0
ch <- 0
// 从通道取数据，并赋值给变量v
v := <- ch
```

>> “箭头”就是数据流的方向
注意接受时候<-和chan变量之间没有空格，虽然有空格也不会报错，但是ide会提示，因此还是依照规范不空格比较好

## Channel 通信

通道接收有如下特性：

① 通道的收发操作在不同的两个 goroutine 间进行。
    由于通道的数据在没有接收方处理时，`数据发送方会持续阻塞`，因此通道的接收必定在另外一个 goroutine 中进行。

② 接收将持续阻塞直到发送方发送数据。
     如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止。

③ 每次接收一个元素。
    通道一次只能接收一个数据元素。

通道的数据接收一共有以下 4 种写法：

```go
// 阻塞接收数据
data := <- ch
// 非阻塞接收，当未接收到数据时，data为通道类型的零值；ok：表示是否接收到数据
data,ok := <- ch  
// 接收任意数据，忽略接收的数据，通常会在并发同步的场景下使用
<- ch
// 循环接收
for data := range ch {}

```

### 单向通信

将一个 channel 变量传递到一个函数时，可以通过将其指定为单向 channel 变量，从而限制该函数中可以对此 channel 的操作，比如只能往这个 channel 中写入数据，或者只能从这个 channel 读取数据。

单向通道的声明格式：

```
// 单向 channel 变量的声明非常简单，只能写入数据的通道类型为chan<-，只能读取数据的通道类型为<-chan，格式如下：

// 支持读写
ch := make(chan int)
// 声明一个只能写入数据的通道类型, 并赋值为ch
var chSendOnly chan<- int = ch
//声明一个只能读取数据的通道类型, 并赋值为ch
var chRecvOnly <-chan int = ch

```

- 一个不能写入数据只能读取的通道是毫无意义的；
- 在某些情况下，定义单向通道有利于代码接口的严谨性；

例如下面的代码示例：

```go
func send(c chan<- int) {
    fmt.Printf("send: %T\n", c)
    c <- 1
}

func recv(c <-chan int) {
    fmt.Printf("recv: %T\n", c)
    fmt.Println(<-c)
}

func main() {
    c := make(chan int)
    fmt.Printf("%T\n", c)
    go send(c)
    go recv(c)
    time.Sleep(1 * time.Second)
}

/*
chan int
send: chan<- int
recv: <-chan int
1
*/
```

### 关闭Channel

关闭 channel 非常简单，直接使用Go语言内置的 close() 函数即可：
```
close(ch)
```

有且只有发送者可以关闭管道，接收者不能关闭管道
接收者可以通过v, ok := <-ch这种方式来测试管道是否关闭，若ok为false则表示管道已关闭

可以用for range来循环取出管道里的数据，range当遇到管道关闭时候就会自动结束循环，例子：

```go
func foo(c chan string) {
    c <- "a"
    c <- "b"
    close(c)
}

func main() {
    c := make(chan string)
    go foo(c)
    for i := range c {
        fmt.Println(i)
    }
}

/*
a
b
*/
```

>> 只有必须告知接收者没有数据要接收了才需要关闭管道，例如上面例子中需要告知range循环该结束了。如果不告知就会造成死锁

### 带缓冲channel

Go语言中无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。这种类型的通道要求发送 goroutine 和接收 goroutine 同时准备好，才能完成发送和接收操作。

Go语言中有缓冲的通道（buffered channel）是一种在被接收前能存储一个或者多个值的通道。这种类型的通道并不强制要求 goroutine 之间必须同时完成发送和接收。

这导致有缓冲的通道和无缓冲的通道之间的一个很大的不同：无缓冲的通道保证进行发送和接收的 goroutine 会在同一时间进行数据交换；有缓冲的通道没有这种保证。

在无缓冲通道的基础上，为通道增加一个有限大小的存储空间形成带缓冲通道。带缓冲通道在发送时无需等待接收方接收即可完成发送过程，并且不会发生阻塞，只有当存储空间满时才会发生阻塞。同理，如果缓冲通道中有数据，接收时将不会发生阻塞，直到通道中没有数据可读时，通道将会再度阻塞。

带缓冲通道在很多特性上和无缓冲通道是类似的。无缓冲通道可以看作是长度永远为 0 的带缓冲通道。因此根据这个特性，带缓冲通道在下面列举的情况下依然会发生阻塞：

- 带缓冲通道被填满时，尝试再次发送数据时发生阻塞。
- 带缓冲通道为空时，尝试接收数据时发生阻塞。

> 为什么Go语言对通道要限制长度而不提供无限长度的通道？

通道（channel）是在两个 goroutine 间通信的桥梁。使用 goroutine 的代码必然有一方提供数据，一方消费数据。当提供数据一方的数据供给速度大于消费方的数据处理速度时，如果通道不限制长度，那么内存将不断膨胀直到应用崩溃。因此，限制通道的长度有利于约束数据提供方的供给速度，供给数据量必须在消费方处理量+通道长度的范围内，才能正常地处理数据。

创建带缓冲通道:

1. 通道实例 := make(chan 通道类型, 缓冲大小)
2. 通道类型：和无缓冲通道用法一致，影响通道发送和接收的数据类型。
3. 缓冲大小：决定通道最多可以保存的元素数量。
4. 通道实例：被创建出的通道实例。

```
func main() {

    // 创建一个3个元素缓冲大小的整型通道
    ch := make(chan int, 3)

    // 查看当前通道的大小
    fmt.Println(len(ch))

    // 发送3个整型元素到通道
    ch <- 1
    ch <- 2
    ch <- 3

    // 查看当前通道的大小
    fmt.Println(len(ch))
}

```

### channel select语句
select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。通常的说，select就是用来监听和channel有关的IO操作，当 IO 操作发生时，触发相应的动作。	

```
select 语句的语法如下：

select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s);
    /* 你可以定义任意数量的 case */
    default : /* 可选 */
       statement(s);
}
```

---

案例1 ：
        如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行，否则的话，如果有default分支，则执行default分支语句，如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行。

```

	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		close(c)
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 3
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <-ch1:
		fmt.Printf("ch1 case...")
	case <-ch2:
		fmt.Printf("ch2 case...")
	default:
		fmt.Printf("default go...")
	}

```

与 switch 语句相比，select 有比较多的限制，其中最大的一条限制就是每个 case 语句里必须是一个 IO 操作，大致的结构如下：

```
select {
    case <-chan1:     
    // 如果chan1成功读到数据，则进行该case处理语句
       
   case chan2 <- 1:
         // 如果成功向chan2写入数据，则进行该case处理语句
    
   default:
      // 如果上面都没有成功，则进入default处理流程
     }
```

在一个 select 语句中，Go语言会按顺序从头至尾评估每一个发送和接收的语句。

如果其中的任意一语句可以继续执行（即没有被阻塞），那么就从那些可以执行的语句中任意选择一条来使用。

如果没有任意一条语句可以执行（即所有的通道都被阻塞），那么有如下两种可能的情况：
如果给出了 default 语句，那么就会执行 default 语句，同时程序的执行会从 select 语句后的语句中恢复；
如果没有 default 语句，那么 select 语句将被阻塞，直到至少有一个通信可以进行下去。

----

案例2 ：
       虽然 select 机制不是专门为超时而设计的，却能很方便的解决超时问题，因为 select 的特点是只要其中有一个 case 已经完成，程序就会继续往下执行，而不会考虑其他 case 的情况。

```
ch := make(chan int)
quit := make(chan bool)
//新开一个协程
go func() {
	for {
		select {
		case num := <-ch:
			fmt.Println("num = ", num)
		case <-time.After(3 * time.Second):
			fmt.Println("超时")
			quit <- true
		}
	}
}() //别忘了()
for i := 0; i < 5; i++ {
	ch <- i
	time.Sleep(time.Second)
}
<-quit
fmt.Println("程序结束")
```

### channel select 多路复用

多路复用是通信和网络中的一个专业术语。多路复用通常表示在一个信道上传输多路信号或数据流的过程和技术。

提示：

报话机同一时刻只能有一边进行收或者发的单边通信，报话机需要遵守的通信流程如下：


- 说话方在完成时需要补上一句“完毕”，随后放开通话按钮，从发送切换到接收状态，收听对方说话。
- 收听方在听到对方说“完毕”时，按下通话按钮，从接收切换到发送状态，开始说话。

电话可以在说话的同时听到对方说话，所以电话是一种多路复用的设备，一条通信线路上可以同时接收或者发送数据。同样的，网线、光纤也都是基于多路复用模式来设计的，网线、光纤不仅可支持同时收发数据，还支持多个人同时收发数据。

在使用通道时，想同时接收多个通道的数据是一件困难的事情。通道在接收数据时，如果没有数据可以接收将会发生阻塞。虽然可以使用如下模式进行遍历，但运行性能会非常差。

```
错误的示例：
for{
    // 尝试接收ch1通道
    data, ok := <-ch1
    // 尝试接收ch2通道
    data, ok := <-ch2
    // 接收后续通道
    …
}

```

```
正确的示例：

select{
    case 操作1:
        响应操作1
    case 操作2:
        响应操作2
    …
    default:
        没有操作情况
}

```

select 多路复用中可以接收的样式：

|  操作  | 语句示例  |
|  ----  | ----  |
| 接收任意数据  | case <- ch |
| 接收变量  | case d:= <-ch  |
| 发送数据  | case ch <- 100 |

- 操作1、操作2：包含通道收发语句
- 响应操作1、响应操作2：当操作发生时，会执行对应 case 的响应操作
- default：当没有任何操作时，默认执行 default 中的语句。


## 多核并行化

Go语言具有支持高并发的特性，可以很方便地实现多线程运算，充分利用多核心 cpu 的性能。

众所周知服务器的处理器大都是单核频率较低而核心数较多，对于支持高并发的程序语言，可以充分利用服务器的多核优势，从而降低单核压力，减少性能浪费。

```go

cpuNum := runtime.NumCPU() //获得当前设备的cpu核心数
fmt.Println("cpu核心数:", cpuNum)

runtime.GOMAXPROCS(cpuNum) //设置需要用到的cpu数量


```

## 大作业
实现一个“多并发获取URL的返回大小、响应时间”的程序,要求如下：

```
./fetch  https://baidu.com https://taobao.com
// 结果如下：
0.35s   304728  https://baidu.com
0.50s   111189  https://taobao.com

```



```go

/*
./fetch  https://baidu.com https://taobao.com
// 结果如下：
0.35s   304728  https://baidu.com
0.50s   111189  https://taobao.com

*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan <- string)  {
	start := time.Now()

	resp, err := http.Get(url)  // 这是用来请求的
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)  // 获取响应内容
	resp.Body.Close()  // don‘t leak resources
	if err != nil{
		ch <- fmt.Sprintf("读取错误%s:%v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

func main()  {
	ch := make(chan string)

	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}

	for range os.Args[1:]{
		fmt.Println(<-ch)
	}

}


/*
命令行运行
E:\golang\HelloGolang\day07_goroutine>go run fetch.go  https://baidu.com https://baidu.com
0.61  321190 https://baidu.com
1.00  320839 https://baidu.com

*/
```

