# map 使用

## 声明和初始化

### 使用var只声明

```go
// 只声明
var gMap map[string]string
```

### 使用var声明并初始化

```go
// 声明并初始化
var hMap = map[string][string]{"a":"b"}
```

### 使用make初始化



## 增删改查

### 增加和删除

### 修改

### 读取数据

两种手段

- 第1种，单变量 lang := m["lang"]，没有key就附1个value类型的0值
- 第2种，双变量 app2, exists := m["app2"] 可以根据exists判断true代表有key

### 增删改查举例

```go
package main

import "fmt"

// 只声明
var m1 map[string]string

// 声明并初始化
var m2 = map[string]string{"a":"b"}

func main(){

	m:=make(map[string]string)

	// 增加，如何批量增加呢？
	m["app"] = "taobao"
	m["lang"] = "golang"

	// 删除
	delete(m, "app")

	// 改
	m["lang"] = "python"

	// 查
	// 单变量形式
	lang := m["lang"]
	fmt.Println(lang)
	// 双变量形式
	lang, exists := m["lang"]
	if exists{
		fmt.Printf("[lang存在，值：%v]\n", lang)
	}
	if !exists{
		fmt.Printf("[lang不存在]\n")
		m["lang"] = "vue"
	}
	fmt.Println(m)

}
```



### 遍历map、按key的顺序遍历 

```go

package main

import "fmt"

func main(){

	m1 := make(map[string]int)

	// 遍历赋值
	keys := make([]string, 0)
	for i:=0; i<10;i++{
		key := fmt.Sprintf("key_%d", i)
		m1[key] = i
		keys = append(keys, key)
	}
	fmt.Println(m1)

	
	for k := range m1{
		fmt.Printf("[key=%s]\n", k)
	}


	// range遍历取值
	fmt.Println("无序遍历")
	for k,v:= range m1{
		fmt.Printf("[%s=%d]\n", k,v)
	}

	// 有序key遍历
	fmt.Println("有序遍历")
	for _,key := range keys{
		fmt.Printf("[%s=%d]\n", key, m1[key])
	}

    	
	// slice排序
	keys2 := []string{"c","a","b"}
	sort.Sort(sort.StringSlice(keys2))
	fmt.Println(keys2)
	
	/*
	[a b c]
	*/
    
}
```

## key的类型: float64可以作为key吗？

> 所有类型的都可以作为key吗？

- bool、int、string
- 特征是 支持 == 和 !== 比较
- float类型可以作为key，写入map时，会做math.Float64bits() 的转换，认为2.4=2.4000xxx1，看起来是同一个key

```go
package main

import "fmt"

func main(){

	m := make(map[float64] int)

	m[2.4] = 2
	fmt.Printf("k:%v, v:%d\n", "2.4000000000000001", m[2.4000000000000001])
	fmt.Println(m[2.4] == m[2.4000000000000001])
	
	/*
	k:2.4000000000000001, v:2
	true
	*/
}
```

## 二维map

- map的value是个map，每一层都要make

- value的类型：任意类型

```go
package main

import "fmt"

func main(){
	var doubleM  map[string]map[string]string
	// panic: assignment to entry in nil map

	doubleM = make(map[string]map[string]string)
	v1 := make(map[string]string)
	v1["k1"] = "v1"

	doubleM["m1"] = v1

	fmt.Println(doubleM)
	/*
	map[m1:map[k1:v1]]
	*/

}
```

## go原生map线程不安全

### fatal error :concurrent map read and map write

```go
package main

import "time"

func main(){

	c := make(map[int]int)

	// 匿名goroutine 循环map
	go func() {
		for i:=0;i<10000;i++{
			c[i] = i
		}
	}()

	// 匿名goroutine 循环读map
	go func() {
		for i:=0;i<10000;i++{
			_ = c[i]
		}
	}()

	time.Sleep(40*time.Minute)

	/*
	fatal error: concurrent map read and map write

	goroutine 6 [running]:
	*/
}
```

### fatal error: concurrent map writes

```go
package main

import "time"

func main(){

	c := make(map[int]int)

	// 匿名goroutine 循环写map
	go func() {
		for i:=0;i<10000;i++{
			c[i] = i
		}
	}()

	// 匿名goroutine 循环写map
	go func() {
		for i:=0;i<10000;i++{
			c[i] = i
		}
	}()
	

	time.Sleep(40*time.Minute)

	/*
		fatal error: concurrent map writes

		goroutine 5 [running]:
	*/

}
```

### 上述问题原因

- go原生的map线程不安全

### 解决方法之一 加锁

- 使用读写锁

### 解决方法之二 使用sync.map

- go 1.9引入的内置方法，并发线程安全的map
- sync.Map 将key和value 按照interface{}存储
- 查询出来后要类型断言 x.(int) x.(string)
- 遍历使用Range() 方法，需要传入一个匿名函数作为参数，匿名函数的参数为k,v interface{}，每次调用匿名函数将结果返回。

- 举例



### sync.map使用 总结



### sync.map 性能对比

- https://studygolang.com/articles/27515

- 性能对比结论

```shell script
只读场景：sync.map > rwmutex >> mutex
读写场景（边读边写）：rwmutex > mutex >> sync.map
读写场景（读80% 写20%）：sync.map > rwmutex > mutex
读写场景（读98% 写2%）：sync.map > rwmutex >> mutex
只写场景：sync.map >> mutex > rwmutex
```

- sync.Map使用场景的建议
  - 读多：给定的key-v只写一次，但是读了很多次，只增长的缓存场景
  - key不相交： 覆盖更新的场景比少

- 结构体复杂的case多不用sync.Map

###  分片锁 并发map github.com/orcaman/concurrent-map

- 基础用法

### 带过期时间的map

- 为什么要有过期时间
- map做缓存用的 垃圾堆积k1  k2 
- 希望缓存存活时间 5分钟，
- 将加锁的时间控制在最低，
- 耗时的操作在加锁外侧做


### 带过期时间的缓存 github.com/patrickmn/go-cache 



# map的实际应用 

![image](05_map.assets/map_cache.png)
![image](05_map.assets/map_cache2.png)
![image](05_map.assets/map_cache3.png)


# map的原理

![image](05_map.assets/map01.png)
![image](05_map.assets/map02.png)
![image](05_map.assets/hmap_bmap.png)

## map底层原理文章推荐

- https://zhuanlan.zhihu.com/p/66676224
- https://segmentfault.com/a/1190000039101378
- https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/

