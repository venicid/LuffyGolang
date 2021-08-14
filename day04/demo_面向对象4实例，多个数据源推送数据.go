package main

import (
	"fmt"
	"log"
)

// 多个数据源推送数据和查询数据
// 定义接口interface
// query 查询数据
// push 写入数据
type DataSource interface {
	Push(data string)
	Query(name string) string
}


// 定义redis结构体和两个方法
type redis struct {
	Name string
	Addr string
}

func (r *redis) Push(data string)  {
	// 真实应该推入它消息队列
	log.Printf("[Pushdata][ds.name:%s][data:%s]", r.Name, data)
}

func (r *redis) Query(name string)string  {
	log.Printf("[Query.data][ds.name:%s][name:%s]", r.Name, name)
	return name +"-"+ r.Name
}

// 定义kafka结构体和两个方法
type kafka struct {
	Name string
	Addr string
}

func (k *kafka) Push(data string) {
	// 真实应该推入它消息队列
	log.Printf("[Pushdata][ds.name:%s][data:%s]", k.Name, data)
}

func (k *kafka) Query(name string) string  {
	log.Printf("[Query.data][ds.name:%s][name:%s]", k.Name, name)
	return name +"-"+ k.Name
}


// 灵魂容器
var DataSourceManager = make(map[string]DataSource)

// 注册方法
func register(name string, ds DataSource) {
	DataSourceManager[name] = ds
}


func main()  {

	r := redis{
		Name: "redis-6.0",
		Addr: "1.1",
	}
	k := kafka{
		Name: "kafka-2.11",
		Addr: "2.2",
	}
	// 将数据源注册到承载的容器中
	register("redis", &r)
	register("kafka", &k)

	// 模拟推送数据
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		for _, ds := range DataSourceManager {
			ds.Push(key)
		}
	}
	// 查询数据
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i)
		for _, ds := range DataSourceManager {
			log.Println(ds.Query(key))
		}
	}

}


/*
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_0]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_0]
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_1]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_1]
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_2]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_2]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_3]
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_3]
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_4]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_4]
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_5]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_5]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_6]
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_6]
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_7]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_7]
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_8]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_8]
2021/08/11 07:37:23 [Pushdata][ds.name:redis-6.0][data:key_9]
2021/08/11 07:37:23 [Pushdata][ds.name:kafka-2.11][data:key_9]
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_0]
2021/08/11 07:37:23 key_0-redis-6.0
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_0]
2021/08/11 07:37:23 key_0-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_1]
2021/08/11 07:37:23 key_1-redis-6.0
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_1]
2021/08/11 07:37:23 key_1-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_2]
2021/08/11 07:37:23 key_2-redis-6.0
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_2]
2021/08/11 07:37:23 key_2-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_3]
2021/08/11 07:37:23 key_3-redis-6.0
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_3]
2021/08/11 07:37:23 key_3-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_4]
2021/08/11 07:37:23 key_4-redis-6.0
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_4]
2021/08/11 07:37:23 key_4-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_5]
2021/08/11 07:37:23 key_5-redis-6.0
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_5]
2021/08/11 07:37:23 key_5-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_6]
2021/08/11 07:37:23 key_6-redis-6.0
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_6]
2021/08/11 07:37:23 key_6-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_7]
2021/08/11 07:37:23 key_7-redis-6.0
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_7]
2021/08/11 07:37:23 key_7-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_8]
2021/08/11 07:37:23 key_8-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_8]
2021/08/11 07:37:23 key_8-redis-6.0
2021/08/11 07:37:23 [Query.data][ds.name:kafka-2.11][name:key_9]
2021/08/11 07:37:23 key_9-kafka-2.11
2021/08/11 07:37:23 [Query.data][ds.name:redis-6.0][name:key_9]
2021/08/11 07:37:23 key_9-redis-6.0

Process finished with the exit code 0

*/