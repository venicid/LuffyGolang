package main

import (
	"log"
	"os"
	"reflect"
)

type Person struct {
	Name string
	Age int
}

func (p Person) ReflectCallFuncWithArgs(name string, age int)  {
	log.Printf("[调用的是带参数的方法][args.name:%v][args.age:%v][p.name:%v][p.age:%v]",
		name,
		age,
		p.Name,
		p.Age,
	)
}

func (p Person) ReflectCallFuncWithoutArgs()  {
	log.Printf("[调用的是带不参数的方法]")
}

func (p Person) ReflectCallFuncWithNoArgs1() {
	log.Printf("[调用的是带不参数的方法 ReflectCallFuncWithNoArgs1]")
}
func (p Person) ReflectCallFuncWithNoArgs2() {
	log.Printf("[调用的是带不参数的方法 ReflectCallFuncWithNoArgs2]")
}


func main() {

	input := os.Args[1]

	p := Person{
		Name: "alex",
		Age:  33,
	}

	// 1. 首先通过reflect.ValueOf 获取到反射类型对象
	value := reflect.ValueOf(p)
	log.Println(value)

	// 2. 调用的是带参数的方法 MethodByName先获取method对象
	funcName1 := value.MethodByName(input)
	// 造一些参数
	args1 := []reflect.Value{
		reflect.ValueOf("alex"),
		reflect.ValueOf(30),
	}
	//Call调用
	funcName1.Call(args1)

	//	3. 不带参数的方法调用,也得造个参数切片
	funcName := value.MethodByName(input)
	args := make([]reflect.Value, 0)
	funcName.Call(args)

}



