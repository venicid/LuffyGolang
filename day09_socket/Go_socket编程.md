# 套接字级编程
---

# Agenda 

1. 网络编程的基础方法，将涉及到主机和服务寻址，也会考虑到TCP和UDP

2. 如何使用GO的TCP和UDP相关的API来构建服务器和客户端。

3. 介绍了原生套接字来实现自己的协议

## TCP/IP协议栈

TCP是一个面向连接的协议，UDP（User Datagram Protocol，用户数据报协议）是一种无连接的协议。

```
--｜应用层｜ 应用层          OSI 5-7
         
      ^       ^
      |       |
     TCP      UDP          OSI 4
     ^         ^
     |         |
         IP                OSI 3
         ^
         |
   h/w interface           OSI 1-2
   
   
```
## IP数据包

IP层提供了无连接的不可靠的传输系统，任何数据包之间的关联必须依赖更高的层来提供。

IP层包头支持数据校验，在包头包括源地址和目的地址。

IP层通过路由连接到因特网，还负责将大数据包分解为更小的包，并传输到另一端后进行重组。

## UDP
UDP是无连接的，不可靠的。它包括IP数据报的内容和端口号的校验。在后面，我们会用它来构建一些客户端/服务器例子

## TCP
UDP是无连接的，不可靠的。它包括IP数据报的内容和端口号的校验。在后面，我们会用它来构建一些客户端/服务器例子。

## 互联网地址
要想使用一项服务，你必须先能找到它。互联网使用地址定位例如计算机的设备。这种寻址方案最初被设计出来只允许极少数的计算机连接上，使用32位无符号整形，拥有高达2^32个地址。这就是所谓的IPv4地址。近年来，连接（至少可以直接寻址）的设备的数量可能超过这个数字，所以在不久的某一天我们将切换到利用128位无符号整数，拥有高2^128个地址的IPv6寻址。这种转换最有可能被已经耗尽了所有的IPv4地址的新兴国家发达地区。

## IPv4地址
IP地址是一个32位整数构成。每个设备的网络接口都有一个地址。该地址通常使用’.’符号分割的4字节的十进制数，例如：”127.0.0.1” 或 “66.102.11.104”。

所有设备的IP地址，通常是由两部分组成：网段地址和网内地址。从前，网络地址和网内地址的分辨很简单，使用字节构建IP地址。

## IPv6地址
因特网的迅速发展大大超出了原来的预期。最初富余的32位地址解决方案已经接近用完。虽然有一些例如NAT地址输入这样不是很完美的解决方法，但最终我们将不得不切换到更广阔的地址空间。IPv6使用128位地址，即使表达同样的地址，字节数变得很麻烦，由’:’分隔的4位16进制组成。一个典型的例子如：2002:c0e8:82e7:0:0:0:c0e8:82e7。

要记住这些地址并不容易！DNS将变得更加重要。有一些技巧用来介绍一些地址，如省略一些零和重复的数字。例如：”localhost”地址是：0:0:0:0:0:0:0:1，可以缩短到::1。

## IP地址类型

### IP类型

“net”包定义了许多类型, 函数，方法用于Go网络编程。IP类型被定义为一个字节数组。

```
type IP []byte
```

有几个函数来处理一个IP类型的变量, 但是在实践中你很可能只用到其中的一些。例如, ParseIP(String)函数将获取逗号分隔的IPv4或者冒号分隔的IPv6地址

### IP掩码

为了处理掩码操作，有下面类型：

```
type IPMask []byte

```

下面这个函数用一个4字节的IPv4地址来创建一个掩码

```
func IPv4Mask(a, b, c, d byte) IPMask
```

另外, 这是一个IP的方法返回默认的掩码

```
func (ip IP) DefaultMask() IPMask
```

需要注意的是一个掩码的字符串形式是一个十六进制数，如掩码255.255.0.0为ffff0000。

一个掩码可以使用一个IP地址的方法，找到该IP地址的网络

```
func (ip IP) Mask(mask IPMask) IP
```

### IPAddr 类型

在net包的许多函数和方法会返回一个指向IPAddr的指针。
```
type IPAddr {
    IP IP
}
```

这种类型的主要用途是通过IP主机名执行DNS查找。

```
func ResolveIPAddr(net, addr string) (*IPAddr, os.Error)
```

### 主机查询

ResolveIPAddr函数将对某个主机名执行DNS查询，并返回一个简单的IP地址。然而，通常主机如果有多个网卡，则可以有多个IP地址。它们也可能有多个主机名，作为别名。

```
func LookupHost(name string) (cname string, addrs []string, err os.Error)
```
这些地址将会被归类为“canonical”主机名。如果你想找到的规范名称，使用 func LookupCNAME(name string) (cname string, err os.Error)

## 服务

服务运行在主机。它们通常长期存活，同时被设计成等待的请求和响应请求。通过各种方法向客户提供服务。互联网的世界基于TCP和UDP这两种通信方法提供许多这些服务，例如点对点, 远过程调用, 通信代理, 和许多其他建立在TCP和UDP之上的服务之上。

### 端口
在每台计算机上可能会提供多种服务，需要一个简单的方法对它们加以区分。TCP，UDP，SCTP或者其他协议使用端口号来加以区分。这里使用一个1到65,535的无符号整数，每个服务将这些端口号中的一个或多个相关联。

有很多“标准”的端口。Telnet服务通常使用端口号23的TCP协议。DNS使用端口号53的TCP或UDP协议。FTP使用端口21和20的命令，进行数据传输。HTTP通常使用端口80，但经常使用，端口8000，8080和8088，协议为TCP。X Window系统往往需要端口6000-6007，TCP和UDP协议。

在Unix系统中, /etc/services文件列出了常用的端口。Go语言有一个函数可以获取该文件。

```
func LookupPort(network, service string) (port int, err os.Error)
```

network是一个字符串例如”tcp”或”udp”, service也是一个字符串，如”telnet”或”domain”(DNS)。

### TCPAddr 类型

TCPAddr类型包含一个IP和一个port的结构：

```
type TCPAddr struct {
    IP   IP
    Port int
}
```

函数ResolveTCPAddr用来创建一个TCPAddr

```
func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error)
```

net是”tcp”, “tcp4”或”tcp6”其中之一，addr是一个字符串，由主机名或IP地址，以及”:”后跟随着端口号组成，例如： “www.google.com:80” 或 ‘127.0.0.1:22”。

## TCP套接字

当知道如何通过网络和端口ID查找一个服务时，然后呢？如果是一个客户端，你需要一个API，去连接到服务，然后将消息发送到该服务，并从服务读取回复。

如果是一个服务器，你需要能够绑定到一个端口，并监听它。当有消息到来，需要能够读取它并回复客户端。

net.TCPConn是允许在客户端和服务器之间的全双工通信的Go类型。两种主要方法是：

```
func (c *TCPConn) Write(b []byte) (n int, err os.Error)
func (c *TCPConn) Read(b []byte) (n int, err os.Error)
```

TCPConn被客户端和服务器用来读写消息。

### TCP客户端

一旦客户端已经建立TCP服务, 就可以和对方设备”通话”了. 如果成功，该调用返回一个用于通信的TCPConn。客户端和服务器通过它交换消息。通常情况下，客户端使用TCPConn写入请求到服务器, 并从TCPConn的读取响应。持续如此，直到任一（或两者）的两侧关闭连接。客户端使用该函数建立一个TCP连接。

```
func DialTCP(net string, laddr, raddr *TCPAddr) (c *TCPConn, err os.Error)
```

其中 laddr 是本地地址，通常设置为 nil 和 一个服务的远程地址raddr, net 是一个字符串，根据是否希望是一个 TCPv4 连接， TCPv6 连接来设置为” tcp4 ”, ” tcp6 “或” tcp “中的一个，当然你也可以不关心链接形式。


客户端可能发送的消息之一就是“HEAD”消息。这用来查询服务器的信息和文档信息。 服务器返回的信息，不返回文档本身。发送到服务器的请求可能是

```
"HEAD / HTTP/1.0\r\n\r\n"
```

服务器典型的响应可能是

```
HTTP/1.0 200 OK
Content-Length: 17931
Content-Type: text/html
Date: Mon, 23 Aug 2021 16:14:18 GMT
Etag: "54d97487-460b"
Server: bfe/1.0.8.18

```

## 控制TCP连接

### 超时

服务端会断开那些超时的客户端，通过下面方式实现：

```
func (c *TCPConn) SetTimeout(nsec int64) os.Error
```
套接字读写前。

### 存活状态
即使没有任何通信，一个客户端可能希望保持连接到服务器的状态。可以使用

```
func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error
```

### UDP数据报

UDP时间服务的客户端并不需要做很多的变化，仅仅改变...TCP...调用为...UDP...调用：

```
/* UDPDaytimeClient
 */
package main

import (
        "net"
        "os"
        "fmt"
)

func main() {
        if len(os.Args) != 2 {
                fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
                os.Exit(1)
        }
        service := os.Args[1]

        udpAddr, err := net.ResolveUDPAddr("up4", service)
        checkError(err)

        conn, err := net.DialUDP("udp", nil, udpAddr)
        checkError(err)

        _, err = conn.Write([]byte("anything"))
        checkError(err)

        var buf [512]byte
        n, err := conn.Read(buf[0:])
        checkError(err)

        fmt.Println(string(buf[0:n]))

        os.Exit(0)
}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
                os.Exit(1)
        }
}

```

UDP服务器端

```

/* UDPDaytimeServer
 */
package main

import (
        "fmt"
        "net"
        "os"
        "time"
)

func main() {

        service := ":1200"
        udpAddr, err := net.ResolveUDPAddr("up4", service)
        checkError(err)

        conn, err := net.ListenUDP("udp", udpAddr)
        checkError(err)

        for {
                handleClient(conn)
        }
}

func handleClient(conn *net.UDPConn) {

        var buf [512]byte

        _, addr, err := conn.ReadFromUDP(buf[0:])
        if err != nil {
                return
        }

        daytime := time.Now().String()

        conn.WriteToUDP([]byte(daytime), addr)
}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
                os.Exit(1)
        }
}
```

### Conn，PacketConn和Listener类型

迄今为止我们已经区分TCP和UDP API的不同，使用例子DialTCP和DialUDP分别返回一个TCPConn和 UDPConn。Conn类型是一个接口，TCPConn和UDPConn实现了该接口。在很大程度上可以通过该接口处理而不是用这两种类型。

现在可以使用一个简单的函数，而不是单独使用TCP和UDP的dial函数。

```
func Dial(net, laddr, raddr string) (c Conn, err os.Error)
```

net可以是”tcp”, “tcp4” (IPv4-only), “tcp6” (IPv6-only), “udp”, “udp4” (IPv4-only), “udp6” (IPv6-only), “ip”, “ip4” (IPv4-only)和”ip6” (IPv6-only)任何一种。它将返回一个实现了Conn接口的类型。注意此函数接受一个字符串而不是raddr地址参数，因此，使用此程序可避免的地址类型。

使用该函数需要对程序轻微的调整。例如, 前面的程序从一个Web页面获取HEAD信息可以被重新写为

```
/* IPGetHeadInfo
 */
package main

import (
        "bytes"
        "fmt"
        "io"
        "net"
        "os"
)

func main() {
        if len(os.Args) != 2 {
                fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
                os.Exit(1)
        }
        service := os.Args[1]

        conn, err := net.Dial("tcp", service)
        checkError(err)

        _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
        checkError(err)

        result, err := readFully(conn)
        checkError(err)

        fmt.Println(string(result))

        os.Exit(0)
}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
                os.Exit(1)
        }
}

func readFully(conn net.Conn) ([]byte, error) {
        defer conn.Close()

        result := bytes.NewBuffer(nil)
        var buf [512]byte
        for {
                n, err := conn.Read(buf[0:])
                result.Write(buf[0:n])
                if err != nil {
                        if err == io.EOF {
                                break
                        }
                        return nil, err
                }
        }
        return result.Bytes(), nil
}

```

使用该函数同样可以简化一个服务器的编写

```
func Listen(net, laddr string) (l Listener, err os.Error)
```

返回一个实现Listener接口的对象. 该接口有一个方法

```
func (l Listener) Accept() (c Conn, err os.Error)
```

这将允许构建一个服务器。使用它, 将使前面给出的多线程Echo服务器改变

```
/* ThreadedIPEchoServer
 */
package main

import (
        "fmt"
        "net"
        "os"
)

func main() {

        service := ":1200"
        listener, err := net.Listen("tcp", service)
        checkError(err)

        for {
                conn, err := listener.Accept()
                if err != nil {
                        continue
                }
                go handleClient(conn)
        }
}

func handleClient(conn net.Conn) {
        defer conn.Close()

        var buf [512]byte
        for {
                n, err := conn.Read(buf[0:])
                if err != nil {
                        return
                }
                _, err2 := conn.Write(buf[0:n])
                if err2 != nil {
                        return
                }
        }
}

func checkError(err error) {
        if err != nil {
                fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
                os.Exit(1)
        }
}

```