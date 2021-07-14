# day01今日目标：用go能够写一些简单的脚本

## 学习go的基本知识



### 1、安装以及开发环境介绍

0. 安装

   Golang 1.16.3  下载地址： https://golang.google.cn/dl/

   IDE: Goland

1. 环境变量

   - GOPATH
   - GOROOT

2. 声明1个程序的入口

   ```go
   package main 
   
   fun main(){
    // 代码块
    fmt.Printf("%s\n", "hello world")
   }
   ```

3. 打印输出

   ```
   fmt.printIn()
   
   fmt.Printf()
   ```

   

### 2、变量&常量

#### 1. 变量声明

   标准格式

   > `var 变量名 变量类型`
   >
   > 变量声明以关键字var开头，变量类型后置，行尾无须分号

   举例子：声明1个整型类型的变量，包括保持整数数值

   > `var age int`

   判断变量的类型

   > var attack = 40
   > `fmt.Printf("%T", attack)`  // int  // 相当于python中的type(attack) 

   小demo

   ```go
package main

import "fmt"

func main() {
    // 声明变量
    var a int 
    var b string
    var c float32
    var d func() bool
    
    fmt.Println(a,b,c,d)
    
    var (
        aa int
        bb string 
        )
    fmt.Println(aa,bb)
    
    var attack = 40
    fmt.Println(attack)  // 编译时自动推导类型
    
    // 判断数据的类型 type(attack)
    fmt.Printf("%T", attack)  // int
}
   ```

   

#### 2. 初始化变量

标准格式

> `var 变量名 变量类型= 表达式`

举例：创建1个游戏角色，初始等级1级

> `var level int = 1`



**短变量声明**

> `var level int = 1`  ==>  `level := 1`

如果短变量被var重新声明，那么编辑器就会报错

```go
var a3 := 300   // 短变量被var再次声明，error
```



在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即使其他变量可能存在重复声明，编译器并不会报错：

```go
conn,err := net.Dial("tcp","127.0.0.1：8080")
conn2,err := net.Dial("tcp","127.0.0.1：8080")
```



小demo

```golang
package main

import "fmt"

func main() {
  
    // 初始化变量
    var a1 int = 100
    var b1 int = 200
    a1,b1 = b1,a1
    fmt.Println(a1,b1)
    
    a3 := 300
    // var a3 := 300   // 短变量被var再次声明，error
    fmt.Println(a3)

}
```



#### 3. 匿名变量_

使用`多重赋值`时，如果`不需要在左值中接受变量`，可以使用匿名变量.

匿名变量以“_”下划线表示

> 匿名变量不占用命名空间，也不会分配内存。匿名变量可以重复声明使用

小demo

```go
package main
import "fmt"
func main() {
    //匿名变量
    f1,_ :=GetData()
    fmt.Println(f1)
    
    // 已经声明的变量，不能够重复声明，匿名变量可以
    // no new variables on left side of :=
    // f1,_ := 300
    
    // 匿名变量不可以直接开头
    //  _ :=1
}

func GetData()(int, int){
    return 100, 200
}
```



#### 4. 作用域

>一个变量(常量，类型or函数)在程序中都有一定的作用范围，称之为作用域

Go(静态语言)会在编译时检查是否每个变量都被使用过
一旦出现未使用的变量，就会报编译错误。

```go
//变量声明必须使用
a2 := 300   //a2 declared and not used
```



根据变量定义的位置，划分为

- 全局变量。 函数外部
- 局部变量。函数内部
- 形式参数。函数定义中的 GetData(name,age)

```golang
package main

import "fmt"

// 全局变量
var global int = 3

func main() {
  
    // 局部变量
    var a int 
    fmt.Println(a)
  
  	fmt.Println(global)
  
    name,age := GetPerson("alex", 30)
  	fmt.Println(name,age)
}

// 形参
func GetData(name,age)(string, int){
    return name, age
}


```



#### 变量总结

- 变量的声明
- 变量的初始化
  - 短变量
- 匿名变量
- 作用域

```go
package main

import "fmt"

// 全局变量
var global int = 3


func main() {
    // 1)声明变量
    var a int 
    var b string
    var c float32
    var d func() bool
    
    fmt.Println(a,b,c,d)
    
    var (
        aa int
        bb string 
        )
    fmt.Println(aa,bb)
    
    var attack = 40
    fmt.Println(attack)  // 编译时自动推导类型
    
    // 判断数据的类型 type(attack)
    fmt.Printf("%T", attack)  // int
    
    //2)初始化变量
    var a1 int = 100
    var b1 int = 200
    a1,b1 = b1,a1
    fmt.Println(a1,b1)
    
    a3 := 300
    // var a3 := 300   // 短变量被var再次声明，error
    fmt.Println(a3)
    
    //3)匿名变量
    f1,_ :=GetData()
    fmt.Println(f1)
    
    // 已经声明的变量，不能够重复声明，匿名变量可以
    // no new variables on left side of :=
    // f1,_ := 300
    
    // 匿名变量不可以直接开头
    //  _ :=1
    
    // 4) 作用域
    // 局部变量
    var a4 int 
    fmt.Println(a4)
  
  	fmt.Println(global)
  	
  	name,age := GetPerson("alex", 30)
  	fmt.Println(name,age)
}


func GetData()(int, int){
    return 100, 200
}

// 形参
func GetPerson(name string,age int)(string, int){
    return name, age
}
```



### 3、整型、浮点型、布尔型、字符串型、切片的入门

#### 1. 整型

- 按长度分为 int8、int64、int32、int64
- 无符号整型 uint8、uint64、uint32、uint64

哪些情况下使用int或uint

> uint在硬件开发中使用 



位运算

bitmap算法：`节省内存`

> 251MB的内存，如何解析100GB的大文件？ 

#### 2.浮点型

Go语言支持两种浮点型数：

> 1. `float32` ： 最大范围约3.4e38
> 2. `float64` ：最大范围约1.8e308

打印浮点数时，可以使用fmt包配合动词”%f“

```go
   // 浮点数
    var floatStr float32;
    floatStr1 := 3.2
    fmt.Println(floatStr)
    fmt.Printf("%T\n", floatStr1)
    fmt.Printf("%.2f\n", floatStr1)
```

#### 3. 布尔型

在Go语言中，以bool类型进行声明：

> var 变量名 bool 

```go
   // 布尔型
    var aVar = 10
    fmt.Println(aVar == 5)
    fmt.Println(aVar == 10)
    fmt.Println(aVar != 5)
    fmt.Println(aVar != 10)
```

布尔型数据只有true和false，且不能参与任何计算以及类型转换

```go
    i := 10
    b := 11
    if i==b {
        i = 1
    }
    fmt.Println(i)
```

与或非

&&优先级高于||

```go
    // 与或非, &&优先级高于||
    s := "hello"
    fmt.Println(s != "" && s[0] == 'x') // false
```

#### 4. 切片

切片是升级版的`数组`，`没有长度限制`，相当于python中的list

切片是一个拥有相同类型元素的可变长度的序列。切片的声明方式如下：

> var name []T 

其中，T代表切片元素类型，可以是整型、浮点型、布尔型、切片、map、函数等等；
切片的元素可以使用[]进行访问，在方括号中提供切片的索引即可访问元素，
索引的范围从0开始，且不超过切片的最大容量。

```go
   // 切片
    // 升级版的数组，没有长度限制
    // 相当于 python中的list
    var listA []int
    fmt.Println(listA)  // [ ]
    fmt.Printf("%T", listA)  // []int
    
    listB := []int{1,2,3}
    fmt.Println(listB) // []int[1 2 3]
    fmt.Println(len(listB))
    
    fmt.Println(listB[0]) // 1
```



#### 5.字符串型

1、字符串的基本定义

go语言从底层支持utf-8

```go
    // 字符串定义
    str1 := "hello"
    str2 := "你好"
    fmt.Printf("%T\n", str1)
    fmt.Printf("%T\n", str2)
    fmt.Println(len(str1))  // 1个字母占1个字节
    fmt.Println(len(str2))  // 1个中文占3个字节，go从底层支持utf8
```

2、字符串转义

> \t 制表符
> \n 换行
> \\ \  反斜杠
> \\"
> \\'
> \r 回车符

如果使用``反引号，会被原样进行赋值和输出

```go
  // 字符串转义
    fmt.Println("\t",str1)
    fmt.Println("\t go大法好")
    fmt.Println(`\t go大法好`)  // \t go大法好
```



3、字符串的理解

字符串的每个元素都叫字符

Go语言的字符有以下2种:

> - `uint8类型`，也叫byte类型。代表了ASCII码的一个字符
> - `rune类型`，代表了一个utf8字符串。用于处理中文、日文

```go
    // 字符的理解
    // 只要是字母和数字都是，unit8类型(也就是byte)
    // rune类型，代表一个utf8字符，处理 中文，日文
    // str = unit8 + rune
    var aa byte = 'a'
    fmt.Printf("%T\n",aa)  // uint8
    var bb rune = '你'
    fmt.Printf("%T\n",bb)  // int32
```



4、字符的应用

如何计算字符串的长度

> ASCII字符使用`len()`函数
>
> Unicode字符串长度使用`utf8.RuneCountInString()`函数

```go
    // 字符的应用
    //如何计算字符串的长度
    str3 := "hello"
    str4 := "你好"
    fmt.Println(len(str3))  // 1个字母占1个字节
    fmt.Println(len(str4))  // 1个中文占3个字节，go从底层支持utf8
    fmt.Println(utf8.RuneCountInString(str4)) // 2
```



5、字符的索引和遍历

直接索引对rune类型无效，可以使用string方法转换

> string([]rune(str6)[0])

如何遍历字符串

> 字符串遍历用for range

```go
 // 字符的索引和遍历
    str5 := "hello"
    str6 := "你好,帅哥美女们"
    fmt.Println(str5[1])  // 101  ascii
    fmt.Println(str6[0])  // 228  ascii
    fmt.Printf("%c\n", str5[1])  // e
    fmt.Printf("%\n", str6[1])  // ½  直接用索引对rune类型是无效的
    fmt.Println(string([]rune(str6)[0]))  // 你
    fmt.Println(string([]rune(str6)[0:5]))  // 你好

    // 遍历
    for i :=0; i< len(str1); i++{
        fmt.Printf("ascii: %c %d\n", str1[i], str1[i])
    }
    for _, s := range  str1{
        fmt.Printf("unicode: %c %d\n ", s, s)
    }
    // 中文只能用 for range
    for _, s := range  str6{
        fmt.Printf("unicode: %c %d\n ", s, s)
    }
```



6、字符串的查找

如何获取字符串中的某一段字符

> strings.Index()： 正向搜索子字符串
>
> strings.LastIndex()：反向搜索子字符串
>
> `通过切片偏移`

```go
    // 查找
    tracer := "死神来了,死神bye bye" 
	
	// 正向搜索字符串
    comma := strings.Index(tracer, ",")
    fmt.Println(",所在的位置:",comma)
    fmt.Println(tracer[comma+1:])  // 死神bye bye

    add := strings.Index(tracer, "+")
    fmt.Println("+所在的位置:",add)  // +所在的位置: -1

    pos := strings.Index(tracer[comma:], "死神")
    fmt.Println("死神，所在的位置", pos) // 死神，所在的位置 1

    fmt.Println(comma, pos, tracer[comma+pos:])  // 12 1 死神bye bye
```



7、修改字符串

Golang语言的字符串是`不可变的`

修改字符串时，可以将字符串`转换为[]byte`进行修改

> []byte和string可以通过强制类型转换

```golang
    // 修改字符串
    // string转换为--> byte ---> 修改 --> string()把ascii转换为string
    readyStr := "Hero never die"
    angleBytes := []byte(readyStr)
    for j:=5; j<=10; j++ {
        angleBytes[j] = ' '
    }
    fmt.Println(angleBytes)  // ascii值
    fmt.Println("str after modify: ",string(angleBytes))
```



8、连接字符串

> 第1种方式：通过+连接
> 	缺点：性能低

> 第2种方式：`stringBuilder.WriteString()`
> 	缺点：节省内存分配，提高处理效率

```go
    // 字符串拼接
    string1 := "美女"
    string2 := "很性感"
    var stringBuilder bytes.Buffer
    stringBuilder.WriteString(string1)
    stringBuilder.WriteString(string2)
    fmt.Println(stringBuilder.String())
```



9、字符串的格式化

print版本：结果写到标准输出

Sprint：结果会以字符串形式返回

```go
// 字符串格式化
// println 直接输出
fmt.Println(",所在的位置:",comma)

// printf 根据格式输出
fmt.Printf("%T\n",aa)  // uint8
fmt.Printf("%c\n", str5[1])  // e

// Sprint 以字符串形式返回
result := fmt.Sprintf(stringBuilder.String())
fmt.Println(result)
```



10、字符串与其他数据类型的转换

1）整数 与 字符串

2）浮点数 与字符串

3）切片 与 字符串

```go
// 字符串与其他类型的转换
// str to int
newStr1 := "1"
intValue, _ := strconv.Atoi(newStr1)
fmt.Printf("%T,%d\n", intValue, intValue)  // int,1

// int to str
intValue2 := 1
strValue := strconv.Itoa(intValue2)
fmt.Printf("%T, %c, %s\n", strValue, strValue, strValue)

// str to float
string3 := "3.1415926"
f,_ := strconv.ParseFloat(string3, 32)
fmt.Printf("%T, %f\n", f, f)  // float64, 3.141593

// str to 切片slice
s := "你好,世界"
fmt.Println(strings.Split(s, ","))  // [你好 世界]

// slice to str
list_ := []string{"aaa", "bbb"}
fmt.Println(strings.Join(list_, ",")) // aaa,bbb
```



11、练习题

第一题：字符串替换

```
str := "语言Python大法好"
输出："语言，Golang大法好"、
```

思路

```go
// - 修改Python为Golang
// - 查找Golang的index
// - index-1，插入，
// - 根据，分割字符串
// - 遍历字符串
```

代码

```go
str8 := "语言Python大法好"
    target := "Python"
    newValue := "Golang"
    splitValue := ","

    index1 := strings.Index(str8, target)
    fmt.Println(index1)  // 6
    fmt.Println(len(target))  // 6

    fmt.Println(str8[:index1])
    fmt.Println(str8[index1+len(target):])

    startStr := str8[:index1]
    endStr := str8[index1+len(target):]

    var stringBuilder2  bytes.Buffer
    stringBuilder2.WriteString(startStr)
    stringBuilder2.WriteString(splitValue)
    stringBuilder2.WriteString(newValue)
    stringBuilder2.WriteString(endStr)
    resultStr := stringBuilder2.String()
    fmt.Println(resultStr)

    fmt.Println(strings.Split(resultStr, splitValue))
```



第二题：找出Golang的个数

```
"语言Golang大法好,语言Golang大法好,语言Golang大法好,语言Golang大法好,语言Golang大法好，Golang”
```

思路

```
// - , 切片
// - 遍历
```

代码

```go
    str9 := "语言Golang大法好,语言Golang大法好,语言Golang大法好,语言Golang大法好,语言Golang大法好，Golang,Golang"
    sliceValue := strings.Split(str9, ",")
    fmt.Println(sliceValue)

    count := 0
    for _, s := range sliceValue{
        fmt.Println(s)
        index := strings.Index(s, "Golang")
        fmt.Println(index)
        if index !=-1 {
            count += 1
        }
        index2 := strings.Index(s[index+len("Golang"):], "Golang")
        fmt.Println(index2)
        if index2 !=-1 {
            count += 1
        }
    }
    fmt.Println(count)
```

思考

跳出循环，当index = -1。当不能用，分割，怎么进行下去呢？



#### 6. 常量

常量就是恒定不变的值。一般不能修改

枚举可以用来定义不同的环境，DEV、FAT、PRE、PRO.

```go
    // 常量
    const number = 19
    const (
        title = "jack"
        salary = 2000
    )

    // 枚举
    const (
        left = iota
        top
        right
        bottom
    )
    fmt.Println(left,top,right,bottom)  // 0 1 2 3
```





### 5、指针的介绍以及使用场景

#### 什么是指针？

var a int = 10

一个变量的声明，将10赋值给了a。
我们把内存当做一个酒店，而每个房间就是一块内存。那么“a int = 10;”的实际含义如下：

> 去酒店订了`一间房间a`，`门牌号暂时用px表示`；`让 10 住进 px`；
> 其中门牌号就是 px，那么px 就是变量的地址

```go
    // 指针与变量
    var room int = 10  // room房间 里面放的 变量10
    var ptr = &room  // 门牌号px  指针  0xc00000a1f0

    fmt.Printf("%p\n", room)  // %!p(int=10)
    fmt.Printf("%p\n", &room)  // 变量的内存地址 0xc00000a1f0

    fmt.Printf("%T, %p\n", ptr, ptr)  // *int, 0xc00000a1f0

    fmt.Println("指针地址",ptr)   // &room
    fmt.Println("指针地址代表的值", *ptr)  // room
```

#### 为什么需要指针？

函数or方法传参数时，值类型可以是 string，bool，int

`高级的类型有map，struct，array/slice`等，如果`直接传值`，比如json为2GB大小，会导致`内存性能的问题`

#### 指针一般用在什么场景?

- 需要修改外部变量的值
- 有超级大的结构体，需要作为函数的参数（使用指针，可以节省内存开销）

#### Go开发可以不使用指针吗？

- 一般可以
- 不要随意在并发场景下使用，指针可以修改指向数据的值
- 像map，array/slice本身也是指针(相当于，python中的引用吗？)

#### 指针地址和指针类型

一个指针变量可以指向任何一个值的内存地址。
 当一个指针被定义后没有分配到任何变量时，它的`默认值为 nil`。指针变量通常缩写为` ptr`。



如何从一个变量中获取指针地址：

> `var ptx := &num`
> //& 可以看做是门口挂的花，也就是门牌号

如何从一个指针地址中获取其对应的实际值：

> `var num := *ptx`
>
> // *可以看做，箭头，指向



#### 如何使用指针修改值

> var a int = 10
> 传入指针&a  *int   ---> 修改指针指向的值 *a = 1000 --->  打印a

```go
package main

func main(){
    // 利用指针修改值
    var num = 10
    modifyFromPoint(num)
    fmt.Println("outside of func",num)

    var num2 = 22
    newModifyFromPoint(&num2)  // 传入指针
    fmt.Println("outside of func",num2)
}

func modifyFromPoint(num int)  {
    // 未使用指针
    num = 10000
    fmt.Println("inside of func:",num)
}

func newModifyFromPoint(ptr *int)  {
    // 使用指针
    *ptr = 1000   // 修改指针地址指向的值
    fmt.Println("inside of func:",*ptr)
}
```



### 大作业：开发《交互式社交机器人》

#### 要求

![image-20210714235903234](day01目标：用go能够写一些简单的脚本.assets/image-20210714235903234.png)

#### 如何启动程序

![image-20210714235920556](day01目标：用go能够写一些简单的脚本.assets/image-20210714235920556.png)

#### 需要的知识

1. 捕获标准输入，并转换为字符串

```go
f := bufio.NewReader(os.Stdin)
Input, _ := f.ReadString('\n')

```

2. 条件判断

```go
if 表达式 {
   // 逻辑式
}else if 表达式 {
   // 逻辑式
}else{
   // 逻辑式
}
```

3. 无限循环

```go
for {
  // 逻辑式
}
```



