package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

func main(){
	m := sync.Map{}

	// 新增
	for i:=0; i<10;i++{
		key := fmt.Sprintf("key_%d",i)
		m.Store(key,i)
	}

	//删除
	m.Delete("key_8")

	// 修改
	m.Store("key_9",999)

	// 查询
	res, loaded := m.Load("key_09")
	if loaded{
		// 类型断言 res.(int)
		log.Printf("[key_09存在：%v 数字类型:%d]", res, res.(int))
	}

	// 遍历 return  false停止
	m.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(int)
		if strings.HasSuffix(k, "3"){
			log.Printf("不想要3")
			//return false  // 停止
			return true
		}else{
			log.Printf("[sync.map.Range][遍历][key=%s][ v=%d]", k,v)
			return true
		}
	})

	// LoadAndDelete先获取值，在删除
	s1, loaded := m.LoadAndDelete("key_7")
	log.Printf("key_7 LoadAndDelete %v\n",s1)


	s2, loaded := m.Load("key_7")
	log.Printf("key_7 LoadAndDelete:%v", s2)

	//LoadOrStore 先获取值，没有的话新增，有的话返回
	actual, loaded := m.LoadOrStore("key_88", 888)
	if loaded{
		log.Printf("key_88原来的值是:%v", actual)
	}else{
		log.Printf("key_88原来没有，实际是:%v", actual)
	}

	actual, loaded = m.LoadOrStore("key_3", 333)
	if loaded{
		log.Printf("key_3原来的值是:%v", actual)
	}else{
		log.Printf("key_3原来没有，实际是:%v", actual)
	}

	/*
	2021/08/07 11:57:44 [sync.map.Range][遍历][key=key_0][ v=0]
	2021/08/07 11:57:44 [sync.map.Range][遍历][key=key_2][ v=2]
	2021/08/07 11:57:44 [sync.map.Range][遍历][key=key_6][ v=6]
	2021/08/07 11:57:44 [sync.map.Range][遍历][key=key_1][ v=1]
	2021/08/07 11:57:44 不想要3
	2021/08/07 11:57:44 [sync.map.Range][遍历][key=key_4][ v=4]
	2021/08/07 11:57:44 [sync.map.Range][遍历][key=key_5][ v=5]
	2021/08/07 11:57:44 [sync.map.Range][遍历][key=key_7][ v=7]
	2021/08/07 11:57:44 [sync.map.Range][遍历][key=key_9][ v=999]

	2021/08/07 11:57:44 key_7 LoadAndDelete 7
	2021/08/07 11:57:44 key_7 LoadAndDelete:<nil>

	2021/08/07 11:57:44 key_88原来没有，实际是:888
	2021/08/07 11:57:44 key_3原来的值是:3


	*/
}