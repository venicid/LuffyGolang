package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main()  {

	/*
		go build -o abc example.go
		echo "date" > a.txt
		./abc < a.txt
	*/
	cmd := exec.Command("sh")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil{
		fmt.Println("os.stderr", err)
		return
	}

}
