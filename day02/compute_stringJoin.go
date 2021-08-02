package main

import (
	"fmt"
	"strings"
)

func main() {

	baseUrl := "http://localhost:8080/api/v1/query?"
	args := strings.Join([]string{"name=luffy", "id=33", "env=fat"}, "&")
	fullURl := baseUrl + args
	fmt.Println(fullURl)

	/*
	http://localhost:8080/api/v1/query?name=luffy&id=33&env=fat
	*/
}