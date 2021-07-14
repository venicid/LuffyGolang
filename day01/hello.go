package main

import (
    "bytes"
    "fmt"
    "strconv"
    "strings"
    "unicode/utf8"
)


func GetData()(int, int){
    return 100, 200
}

// 形参
func GetPerson(name string,age int)(string, int){
    return name, age
}

// 全局变量
var global int = 3

func main(){
    fmt.Println("Hello World!")
    fmt.Println("hello 可爱的帅哥美女")

    /*
    *  变量
    */
    // 声明变量
    var a int
    var b string
    var c float32
    var d func() bool

    fmt.Println(a,b,c,d)

    //var (
    //	aa int
    //	bb string
    //)
    //fmt.Println(aa, bb)

    var attack = 40
    fmt.Println(attack)  // 编译时自动推导类型
    // 判断数据的类型 type(attack)
    fmt.Printf("%T", attack)

    // 初始化变量
    var a1 int = 100
    var b1 int = 300
    fmt.Println(a1,b1)
    b1, a1 = a1, b1
    fmt.Println(a1,b1)


    var level = 1
    level2:=1
    /*
    静态语言，定义必须使用
    # command-line-arguments
    .\hello.go:35:5: level2 declared but not used
    */
    // var a3 := 300   // 短变量被var再次声明，error
    fmt.Println(level)
    fmt.Println(level2)

    //匿名变量
    f1, _ := GetData()
    fmt.Println(f1)

    f2, _ := GetData()
    fmt.Println(f2)

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

    /*
    *  浮点数
    */
    floatStr :=3.145
    fmt.Println(floatStr)  // 3.145
    fmt.Printf("%T\n", floatStr) // float64
    fmt.Printf("%.2f\n", floatStr)

    /* 布尔型
    */
    var aVar = 10
    fmt.Println(aVar == 5)
    fmt.Println(aVar == 10)
    fmt.Println(aVar != 5)
    fmt.Println(aVar != 10)
    i := 10
    b11 := 11
    if i==b11 {
        i = 1
    }
    fmt.Println(i)

    // && 与
    //|| 或
    // 与或非, &&优先级高于||
    s := "hello"
    fmt.Println(s != "" && s[0] == 'x') // false

    /*切片
    */
    // 切片是升级版的数组, 可动态扩张的数组
    // 切片，无长度限制，数组有
    // 相当于 python中的list
    var listA []int
    fmt.Println(listA)  // [ ]
    fmt.Printf("%T", listA)  // []int

    listB := []int{1,2,3}
    fmt.Println(listB) // []int[1 2 3]
    fmt.Println(len(listB))

    fmt.Println(listB[0]) // 1


    /*
    *  字符串
    */
    // 字符串定义
    str1 := "hello"
    str2 := "你好"
    fmt.Printf("%T\n", str1)
    fmt.Printf("%T\n", str2)

    // 字符串转义
    fmt.Println("\t",str1)
    fmt.Println("\t go大法好")
    fmt.Println(`\t go大法好`)  // \t go大法好

    // 字符的理解
    // 只要是字母和数字都是，unit8类型(也就是byte)
    // rune类型，代表一个utf8字符，处理 中文，日文
    // str = unit8 + rune
    var aa byte = 'a'
    fmt.Printf("%T\n",aa)  // uint8
    var bb rune = '你'
    fmt.Printf("%T\n",bb)  // int32

    // 字符的应用
    //如何计算字符串的长度
    str3 := "hello"
    str4 := "你好"
    fmt.Println(len(str3))  // 1个字母占1个字节
    fmt.Println(len(str4))  // 1个中文占3个字节，go从底层支持utf8
    fmt.Println(utf8.RuneCountInString(str4)) // 2

    // 字符的索引和遍历
    // 索引
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
    //  for range(推荐)
    // 中文只能用for range
    for _, s := range  str1{
        fmt.Printf("unicode: %c %d\n ", s, s)
    }
    for _, s := range  str6{
        fmt.Printf("unicode: %c %d\n ", s, s)
    }


    // 查找
    // 先去找index，再用切片 （不能直接找到某个值）
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


    // 修改字符串
    // string转换为--> byte ---> 修改 --> string()把ascii转换为string
    readyStr := "Hero never die"
    angleBytes := []byte(readyStr)
    for j:=5; j<=10; j++ {
        angleBytes[j] = ' '
    }
    fmt.Println(angleBytes)  // ascii值
    fmt.Println("str after modify: ",string(angleBytes))


    // 字符串拼接
    string1 := "美女"
    string2 := "很性感"
    var stringBuilder bytes.Buffer
    stringBuilder.WriteString(string1)
    stringBuilder.WriteString(string2)
    fmt.Println(stringBuilder.String())

    // 字符串格式化
    result := fmt.Sprintf(stringBuilder.String())
    fmt.Println(result)


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
    s11 := "你好,世界"
    fmt.Println(strings.Split(s11, ","))  // [你好 世界]

    // slice to str
    list_ := []string{"aaa", "bbb"}
    fmt.Println(strings.Join(list_, ",")) // aaa,bbb

    // 1. str := "语言Python大法好"
    // 输出："语言，Golang大法好"
    // - 修改Python为Golang
    // - 查找Golang的index
    // - index-1，插入，
    // - 根据，分割字符串
    // - 遍历字符串

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


    // 2. "语言Golang大法好,语言Golang大法好,语言Golang大法好,语言Golang大法好,语言Golang大法好，Golang”
    // 统计Golang的次数
    // - , 切片
    // - 遍历
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
    // 跳出循环，当index=-1时


    /*
    *  常量与枚举
    */
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


    /*
    * 指针入门
    */
    // 指针与变量
    var room int = 10  // room房间 里面放的 变量10
    var ptr = &room  // 门牌号px  指针  0xc00000a1f0

    fmt.Printf("%p\n", room)  // %!p(int=10)
    fmt.Printf("%p\n", &room)  // 变量的内存地址 0xc00000a1f0

    fmt.Printf("%T, %p\n", ptr, ptr)  // *int, 0xc00000a1f0

    fmt.Println("指针地址",ptr)   // &room
    fmt.Println("指针地址代表的值", *ptr)  // room

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



