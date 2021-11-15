/*

Connecting to  :3333
hello 10
hello 9
hello 8
hello 7
hello 6
Read
Send

Process finished with the exit code 0

*/

package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
)

//func handleWrite(conn net.Conn, done chan string)  {
func handleWrite(conn net.Conn, wg *sync.WaitGroup)  {
	defer wg.Done()

	for i := 10; i >0 ; i-- {
		_, e := conn.Write([]byte("hello "+ strconv.Itoa(i) + "\r\n"))
		if e !=nil{
			fmt.Println("Error to send msg because of ", e.Error())
			break
		}
	}
	//done <- "Send"
}

//func handleRead(conn net.Conn, done chan string)  {
func handleRead(conn net.Conn, wg *sync.WaitGroup)  {
	defer wg.Done()

	buf := make([]byte, 1024)
	requestLen, err := conn.Read(buf)
	if err !=nil{
		fmt.Println("Error to read msg because of", err.Error())
	}
	fmt.Println(string(buf[:requestLen-1]))
	//done <- "Read"
}

func main(){
	var host = flag.String("host", "", "host")
	var port = flag.String("port", "3333", "port")
	flag.Parse()

	conn, err := net.Dial("tcp", *host + ":" + *port)
	if err != nil{
		fmt.Println("Error Connect ", err)
	}
	defer conn.Close()
	fmt.Println("Connecting to ", *host + ":" + *port)

	/* channel方式
	done := make(chan string)
	go handleWrite(conn, done)
	go handleRead(conn, done)
	fmt.Println(<-done)
	fmt.Println(<-done)
	*/

	// waitGroup 方式
	var wg sync.WaitGroup
	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()


}