package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// 方式一
	bytes1, _ := ioutil.ReadFile("go.mod")

	// 方式二
	file, _ := os.Open("go.mod")
	bytes2, _ := ioutil.ReadAll(file)
	file.Close()

	// 方法三
	file, _ = os.Open("go.mod")

	bo := bufio.NewReader(file)  // 缓存
	buf := make([]byte, 200)
	bo.Read(buf)

	fmt.Println(string(bytes1))
	fmt.Println(string(bytes2))
	fmt.Println(string(buf))
}

