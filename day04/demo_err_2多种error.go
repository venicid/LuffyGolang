package main

import (
	"log"
	"os"
)

// 以os包举例 提供了 LinkError、PathError、SyscallError
func main() {

	f1, err := os.Stat("test.txt")
	if err != nil {
		switch err.(type) {
		case *os.PathError:
			log.Printf("PathError")
		case *os.LinkError:
			log.Printf("LinkError")
		case *os.SyscallError:
			log.Printf("SyscallError")
		}
	} else {
		log.Printf("f1:%v", f1)
	}
}

/*
2021/08/30 23:41:19 PathError
*/