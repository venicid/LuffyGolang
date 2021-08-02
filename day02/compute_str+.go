package main

import "fmt"

func main() {

	s1:= "http://"
	s2 := "loaclhost:8080"
	s3 := s1 + s2
	fmt.Println(s3)

	s4 := "http://localhost:8080" +
		"/api/v1" +
		"/login"
	fmt.Println(s4)

	/*
	http://loaclhost:8080
	http://localhost:8080/api/v1/login
	*/
}

