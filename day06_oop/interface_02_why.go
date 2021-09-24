/*
接口的使用场景
*/
package main

import (
	"errors"
	"fmt"
)

// 接口
type Car interface {
	Drive(dr Driver)
}

type Driver struct {
	Name string
}

type Benz struct{}

func (this *Benz) Drive(dr Driver) {
	fmt.Printf("%s drives Benz.\n", dr.Name)
}

type Audi struct {}

func (this *Audi) Drive(dr Driver) {
	fmt.Printf("%s drives Audi.\n", dr.Name)
}

func NewCar(c string)(Car, error)  {
	switch c {
	case  "Audi":
		return &Audi{}, nil
	case "Benz":
		return &Benz{}, nil
	default:
		return nil, errors.New("not support")

	}

}


func main() {
	me := Driver{Name: "gbe"}

	car, err := NewCar("Audi")
	if err != nil{
		fmt.Println(err)
	}else{
		car.Drive(me)
	}

}

/*
gbe drives Audi.
*/