#  Go语言面向对象编程讲义
 - 六脉神剑

> 1. 结构体“类”的定义&实例化
> 2. 动态属性
> 3. 匿名结构体
> 4. 动态方法
> 5. 构造函数
> 6. 组合（继承）

## 结构体"类"的定义

Go语言可以通过自定义的方式形成新的类型，结构体就是这些类型中的一种复合类型，结构体是由零个或多个任意类型的值聚合成的实体，每个值都可以称为结构体的成员。

结构体成员也可以称为“字段”/“属性”，这些字段有以下特性：

1. 字段拥有自己的类型和值；

2. 字段名必须唯一；

3. 字段的类型也可以是结构体，甚至是字段所在结构体的类型。

4. 字段在Golang中只有实例化才有意义，也称“动态字段”

使用关键字 type 可以将各种基本类型定义为自定义类型，基本类型包括整型、字符串、布尔等。结构体是一种复合的基本类型，通过 type 定义为自定义类型后，使结构体更便于使用。

结构体的定义格式如下：

```
type 类型名 struct {
    字段1 字段1类型
    字段2 字段2类型
    …
}
```

对各个部分的说明：

类型名：标识自定义结构体的名称，在同一个包内不能重复。

struct{}：表示结构体类型，type 类型名 struct{}可以理解为将 struct{} 结构体定义为类型名的类型。

字段1、字段2……：表示结构体字段名，结构体中的字段名必须唯一。

字段1类型、字段2类型……：表示结构体各个字段的类型。

## 实例化

结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存，因此必须在定义结构体并实例化后才能使用结构体的字段。

实例化就是根据结构体定义的格式创建一份与格式一致的内存区域，结构体实例与实例间的内存是完全独立的。

### 基本的实例化形式
结构体本身是一种类型，可以像整型、字符串等类型一样，以 var 的方式声明结构体即可完成实例化。

基本实例化格式如下：

```
var ins T
```
其中，T 为结构体类型，ins 为结构体的实例。

### 创建指针类型的结构体
Go语言中，还可以使用 new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成指针类型的结构体。

使用 new 的格式如下：

```
ins := new(T)
```
其中：
T 为类型，可以是结构体、整型、字符串等。
ins：T 类型被实例化后保存到 ins 变量中，ins 的类型为 *T，属于指针。

### 取结构体的地址实例化

在Go语言中，对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作，取地址格式如下：

```
ins := &T{}
```

其中：
T 表示结构体类型。
ins 为结构体的实例，类型为 *T，是指针类型。

>>> 取地址实例化是最广泛的一种结构体实例化方式，可以使用函数封装上面的初始化过程，代码如下：

```
func newCommand(name string, varref *int, comment string) *Command {
    return &Command{
        Name:    name,
        Var:     varref,
        Comment: comment,
    }
}
cmd = newCommand(
    "version",
    &version,
    "show version",
)
```

## 初始化结构体的字段
结构体在实例化时可以直接对成员变量进行初始化，初始化有两种形式分别是以字段“键值对”形式和多个值的列表形式，键值对形式的初始化适合选择性填充字段较多的结构体，多个值的列表形式适合填充字段较少的结构体。

### 使用“键值对”初始化结构体
```
ins := 结构体类型名{
    字段1: 字段1的值,
    字段2: 字段2的值,
    …
}
```
下面是对各个部分的说明：

结构体类型：定义结构体时的类型名称。

字段1、字段2：结构体成员的字段名，结构体类型名的字段初始化列表中，字段名只能出现一次。

字段1的值、字段2的值：结构体成员字段的初始值。

### 使用多个值的列表初始化结构体
Go语言可以在“键值对”初始化的基础上忽略“键”，也就是说，可以使用多个值的列表初始化结构体的字段。
1)  多个值列表初始化结构体的书写格式
多个值使用逗号分隔初始化结构体，例如:

```
ins := 结构体类型名{
    字段1的值,
    字段2的值,
    …
}
```
使用这种格式初始化时，需要注意：

- `必须初始化结构体的所有字段`。

- 每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。

- 键值对与值列表的初始化形式不能混用。

## 初始化匿名结构体

匿名结构体没有类型名称，无须通过 type 关键字定义就可以直接使用。

1) 匿名结构体定义格式和初始化写法
匿名结构体的初始化写法由结构体定义和键值对初始化两部分组成，结构体定义时没有结构体类型名，只有字段和类型定义，如下格式所示：

```
ins := struct {
    // 匿名结构体字段定义
    字段1 字段类型1
    字段2 字段类型2
    …
}{
    // 字段值初始化
    初始化字段1: 字段1的值,
    初始化字段2: 字段2的值,
    …
}
```
键值对初始化部分是可选的，不初始化成员时，匿名结构体的格式变为

```
ins := struct {
    字段1 字段类型1
    字段2 字段类型2
    …
}
```
Tips: 如果只是临时使用struct一次，而不是多次使用，用匿名struct即可

## 构造函数
Go语言的类型或结构体没有构造函数的功能，但是可以使用结构体初始化的过程来模拟实现构造函数。

## 方法
在面向对象的语言中，类拥有的方法一般被理解为类可以做的事情。在Go语言中“方法”的概念与其他语言一致，只是Go语言建立一种“接收器”来强调方法的概念

### 面向过程实现方法
面向过程中没有“方法”概念，`只能通过结构体和函数`，由使用者使用函数参数和调用关系来形成接近“方法”的概念，代码如下：

```
type Bag struct {
    items []int
}

func Insert(b *Bag, itemid int) {
    // some things
}

```

### Go语言的结构体方法
```
type Bag struct {
    items []int
}
func (b *Bag) Insert(itemid int) {
    // some things
    }
```
"b *Bag" 即是“接收器”

`接收器的格式如下：`

```
func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

指针和非指针接收器的使用

在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器，`大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针`。


### 组合
在Go语言中，相比较于继承，组合更受青睐。`尽量使用内嵌结构体的方式`：
Go语言的结构体内嵌有如下特性。

1) 内嵌的结构体可以直接访问其成员变量  blackcat.name

2) 内嵌结构体的字段名是它的类型名

# 接口
## 为什么需要接口
下面是一个没有“接口”的代码场景例子：

```go
// 没有interface的时候
package main

import "fmt"

type Driver struct {
    Name string
}

type Benz struct{}

func (this *Benz) Drive(dr Driver) {
    fmt.Printf("%s drives Benz.\n", dr.Name)
}

func main() {
    me := Driver{Name: "gbe"}

    // 开Benz
    car := &Benz{}
    car.Drive(me)
}
```
输出

```
gbe drives Benz.
```

现在增加一辆Audi

```go
// 没有interface的时候
package main

import "fmt"

type Driver struct {
    Name string
}

type Benz struct{}

func (this *Benz) Drive(dr Driver) {
    fmt.Printf("%s drives Benz.\n", dr.Name)
}

// 添加一个奥迪车
type Audi struct {}

func (this *Audi) Drive(dr Driver) {
    fmt.Printf("%s drives Audi.\n", dr.Name)
}

func main() {
    me := Driver{Name: "Cyent"}

    // car为外部传递进来的值，假设传进来的是Audi
    get_car := "Audi"
    switch get_car {
    case "Benz":
        car := &Benz{}
        car.Drive(me)
    case "Audi":
        car := &Audi{}
        car.Drive(me)
    }
}
```

注意上面main函数里用了switch可以对传递进来的变量做判断，但这么写也不够好，`因为每增加一辆车，switch都要多加一个case`，那有没有无论增加多少辆车，main函数永远不用改代码的办法呢？

```go
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
```





## 面向对象综合复习
Go很好的采用了OOP设计模式，省略了其余的部分并提供了编写多态代码的更好方法。下面是更全面的综合复习示例：

```
// 以下是Go中OOP的快速视图。 从这三个结构开始：
type Animal struct {
	Name string
	mean bool
}

type Mouse struct {
	Basics Animal
	squeakStrength int
}

type Dog struct {
	Animal
	BarkStrength int
}

```

以上定义了一个基础结构体和两个特定于基础结构体的结构体。Animal结构体包含所有Animals共享的属性，另外两个结构体则特定于Mouse和Dog。

除了mean之外，所有成员属性（字段）都是公共的。 Animal结构中的mean字段以小写字母开头。 在Go中，变量，结构，字段，函数等的第一个字母的大小写确定了其可访问的规范。 `使用大写字母，它是公开的(public)，使用小写字母，它是私人的(private)。`

由于Go中没有继承，因此组合是唯一的选择。 Mouse结构有一个名为Basics的字段，其类型为Animal。 Dog结构使用未命名的struct（嵌入）作为Animal类型。 由你来决定哪个更适合。

要为Mouse和Dog创建成员函数（方法），语法如下：

```
func (dog *Dog) MakeNoise() {
	barkStrength := dog.BarkStrength

	if dog.mean == true {
		barkStrength = barkStrength * 5
	}

	for bark := 0; bark < barkStrength; bark++ {
		fmt.Printf("BARK ")
	}

	fmt.Println("")
}

func (mouse *Mouse) MakeNoise() {
	squeakStrength := mouse.squeakStrength

	if mouse.Basics.mean == true {
		squeakStrength = squeakStrength * 5
	}

	for squeak := 0; squeak < squeakStrength; squeak++ {
		fmt.Printf("SQUEAK ")
	}

	fmt.Println("")
}

```

在方法名称之前，我们指定了一个接收器，它是指向每种类型的指针。 现在Mouse和Dog都可以调用MakeNoise方法。

它们的MakeNoise方法都做同样的事情。 每个Animal都会根据它们的Mouse或Dog的MakeNoise方法打印一些字符，如果mean为true的话。它告诉我们如何访问引用的对象（值）。

现在已经介绍了封装，组合，访问规范和方法等， 剩下的就是如何创建多态行为。

接下来使用接口来创建多态行为：

```
type AnimalSounder interface {
    MakeNoise()
}
 
func MakeSomeNoise(animalSounder AnimalSounder) {
    animalSounder.MakeNoise()
}

```


当接口只包含一个方法时，Go中约定通常使用“er”后缀命名接口。

在Go中，通过方法实现接口的任何类型都表示接口类型。 在我们的例子中，Cat和Dog都使用指针接收器实现了AnimalSounder接口，因此被认为是AnimalSounder类型。

这意味着Mouse和Dog的指针都可以作为参数传递给MakeSomeNoise函数。 MakeSomeNoise函数通过AnimalSounder接口实现多态行为。

如果你想减少Mouse和Dog的MakeNoise方法中的代码重复，可以为Animal类型创建一个方法来处理它：

```
func (animal *Animal) PerformNoise(strength int, sound string) {
    if animal.mean == true {
        strength = strength * 5
    }

    for voice := 0; voice < strength; voice++ {
        fmt.Printf("%s ", sound)
    }

    fmt.Println("")
}

func (dog *Dog) MakeNoise() {
    dog.PerformNoise(dog.BarkStrength, "BARK")
}

func (mouse *Mouse) MakeNoise() {
    mouse.Basics.PerformNoise(mouse.squeakStrength, "SQUEAK")
}

```

现在，Animal类型有了一个包含调用MakeNoise业务逻辑的处理方法。这里的另一个好处是我们不需要将mean字段作为参数传递，因为它已经属于Animal类型。

下面是完整代码：

```go
package main

import "fmt"

type Animal struct {
	Name string
	mean bool
}

type Mouse struct {
	Basics         Animal
	squeakStrength int
}

type Dog struct {
	Animal
	BarkStrength int
}

type AnimalSounder interface {
	MakeNoise()
}

func MakeSomeNoise(animalSounder AnimalSounder) {
	animalSounder.MakeNoise()
}

func (animal *Animal) PerformNoise(strength int, sound string) {
	if animal.mean == true {
		strength = strength * 5
	}
	for voice := 0; voice < strength; voice++ {
		fmt.Printf("%s ", sound)
	}
	fmt.Println("")
}

func (dog *Dog) MakeNoise() {
	dog.PerformNoise(dog.BarkStrength, "BARK")
}

func (mouse *Mouse) MakeNoise() {
	mouse.Basics.PerformNoise(mouse.squeakStrength, "SQUEAK")
}


func main() {
	myDog := &Dog{
		Animal{
			"ALEX", // Name
			false,   // mean
		},
		2, // BarkStrength
	}
	myMouse := &Mouse{
		Basics: Animal{
			Name: "BAOER",
			mean: true,
		},
		squeakStrength: 3,
	}

	MakeSomeNoise(myDog)
	MakeSomeNoise(myMouse)
}
```
另外一种关于在结构中嵌入接口的实现分享

```
package main

import (
    "fmt"
)

type HornSounder interface {
    SoundHorn()
}

type Vehicle struct {
    List [2]HornSounder
}

type Car struct {
    Sound string
}

type Bike struct {
   Sound string
}

func main() {
    vehicle := new(Vehicle)
    vehicle.List[0] = &Car{"BEEP"}
    vehicle.List[1] = &Bike{"RING"}

    for _, hornSounder := range vehicle.List {
        hornSounder.SoundHorn()
    }
}

func (car *Car) SoundHorn() {
    fmt.Println(car.Sound)
}

func (bike *Bike) SoundHorn() {
    fmt.Println(bike.Sound)
}

func PressHorn(hornSounder HornSounder) {
    hornSounder.SoundHorn()
}
```
在这个示例中，Vehicle结构体维护了一个实现HornSounder接口的列表。 在main函数中创建了一个Vehicle实例，并为其成员变量List初始化了Car和Bike实例。 这种赋值是允许的，因为Car和Bike存在Sound。 然后使用一个简单的循环，使用hornSounder接口来SoundHorn。