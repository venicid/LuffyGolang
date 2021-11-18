package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"strconv"
	"time"
)

func main()  {

	conn, err := net.Dial("tcp", "127.0.0.1:8866")
	if err != nil{
		log.Fatal(err)
	}

	defer conn.Close()

	for i := 0; i < 10; i++ {
		var err error

		msg1 := strconv.Itoa(i) + "[路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，]\n"
		data ,err := encode(msg1)
		_, err = conn.Write(data)

		msg1 = strconv.Itoa(i) + "bbbbbb\n"
		data ,err = encode(msg1)
		_, err = conn.Write(data)

		msg1 = strconv.Itoa(i) + "ccccc\n"
		data ,err = encode(msg1)
		_, err = conn.Write(data)

		if err != nil{
			panic(err)
		}
	}

	/*
		for i := 0; i < 10; i++ {
			var err error
			_, err = conn.Write([]byte(strconv.Itoa(i) + "[路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，路飞学城，]\n"))
			_, err = conn.Write([]byte(strconv.Itoa(i) + "bbb\n"))
			_, err = conn.Write([]byte(strconv.Itoa(i) + "ccc\n"))
			if err != nil{
				panic(err)
			}
		}
	*/

	time.Sleep(time.Second * 1)
}

func encode(message string) ([]byte,error)  {
	var length = int32(len(message))  // 读取消息的长度，转换为int32类型（占用4个字节）

	var pkg = new(bytes.Buffer)
	err := binary.Write(pkg, binary.BigEndian, length) // 写入msg长度，以大端的方式存放
	if err != nil{
		return nil, err
	}

	err = binary.Write(pkg, binary.BigEndian, []byte(message))  // 写入消息实体
	if err != nil{
		return nil, err
	}

	return pkg.Bytes(), nil

}