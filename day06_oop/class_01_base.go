package main

import "fmt"

/*
class People:
	name = xxx
	age = 18

class Child(People):
	pass

Gof 设计模式
Mixins设计模式
*/


type People struct {  //首字母大写：公有，小写是私有
	name string
	age int
	child *People  // 第一种组合方式  people.child.name
}

type Cat struct {
	// 结构体的定义
	Name string  // 动态字段/属性的定义
	color string
}

func Eat1(cat *Cat)  {
	fmt.Println("吃饭")
}

func (cat *Cat) Eat()  {
	// c *Cat就是接收器，结构体方法
	fmt.Println("eat, 吃饭")
}


type BlackCat struct {
	Cat  // 第二种组合（继承） 内嵌结构体，blackcat.name 推荐用这种方式
}

func NewBlackCat(name string) *BlackCat  {
	// 构造函数
	cat := &BlackCat{}
	cat.Name = name
	cat.color = "black"
	return cat
}


func main()  {
	people := &People{}
	people.name = "alex"
	people.age = 18
	fmt.Println(people)

	initValue := &People{  // 实例化结构体的时候给予字段属性的默认值
		name: "boy",
		age: 18,
		child: &People{
			name: "zhangsan",
			age: 10,
		},
	}
	fmt.Println(initValue)


	// 黑猫，为什么需要构造函数
	cat := Cat{
		Name: "tom",
		color: "black",
	}
	blackCat := &BlackCat{
		cat,
	}
	blackCat.Name = "jack"
	blackCat.Name = "black"
	fmt.Println(blackCat)

	// 使用构造函数后
	blackCat1 := NewBlackCat("jerry")
	fmt.Println(blackCat1.Name, blackCat1.color)

	// 为什么需要给类绑定方法
	Eat1(&cat)

	blackCat1.Eat()
}
