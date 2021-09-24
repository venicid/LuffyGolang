package main

import (
	"errors"
	"fmt"
)

var errorByZero = errors.New("the intValue is zero")

func errorDemo(i int)(int , error)  {

	if i ==0{
		return 0, errorByZero
	}
	return i, nil
}

func main()  {

	res, err := errorDemo(1)
	if err!= nil{
		fmt.Println(err)
	}else{
		fmt.Println(res)
	}


}


