package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// 带过期时间的map定时器
type Cache struct {
	sync.RWMutex
	mp map[string]*item
}

type item struct {
	value int  // 值
	ts int64   // 时间戳，item被创建出来的时间
}

func (c *Cache) Get(key string) *item  {
	c.RLock()
	defer c.RUnlock()
	return c.mp[key]
}

func (c *Cache)Set(key string, value *item)  {
	c.Lock()
	defer c.Unlock()
	c.mp[key] = value
}

// 获取cache中map的数量
func (c *Cache) CacheNum() int  {
	c.RLock()
	defer c.RUnlock()
	return len(c.mp)
}

// 清除过期的值
func (c *Cache) clean(timeDelta int64)  {
	// 每5秒执行一次
	for{
		now := time.Now().Unix()
		toDeleteKeys := make([]string, 0) // 待删除的key的切片

		// 先加读锁，把所有待删除的拿到
		c.RLock()
		for k,v:= range c.mp{
			// 时间比较
			// 认为这个k,v过期了,
			// 不直接删除，为了降低加锁时间，加入待删除的切片
			if now - v.ts > timeDelta{
				toDeleteKeys = append(toDeleteKeys, k)
			}
		}
		c.RUnlock()

		// 加写锁 删除,降低加写锁的时间
		c.Lock()
		for _, k := range toDeleteKeys{
			log.Printf("[删除过期数据][key:%s]", k)
			delete(c.mp, k)
		}
		c.Unlock()

		time.Sleep(2 * time.Second)
	}

}

func main() {

	c:=Cache{
		mp: make(map[string] *item),
	}

	// 让清理的任务异步执行
	// 每2秒运行一次，检查时间差大于30秒item 就删除
	go c.clean(30)


	// 设置缓存，30s后过期
	// 从mysql中读取到了数据，塞入缓存
	for i:=0;i<10;i++{

		now := time.Now().Unix()
		im := &item{
			value: i,
			ts: now,
		}

		key := fmt.Sprintf("key_%d", i)

		log.Printf("[设置缓存][item][key:%s][v:%v]", key, im)
		c.Set(key, im)
	}

	log.Printf("缓存中的数据量:%d", c.CacheNum())
	time.Sleep(36 * time.Second)
	log.Printf("缓存中的数据量:%d", c.CacheNum())

	// 更新缓存，30s后过期
	for i:=0; i<5;i++{
		now := time.Now().Unix()
		im := &item{
			value: i,
			ts: now,
		}

		key := fmt.Sprintf("key_%d", i)

		log.Printf("[更新缓存][item][key:%s][v:%v]", key, im)
		c.Set(key, im)
	}
	log.Printf("缓存中的数据量:%d", c.CacheNum())

	select{}

}

/*
2021/08/07 17:44:29 [设置缓存][item][key:key_0][v:&{0 1628329469}]
2021/08/07 17:44:29 [设置缓存][item][key:key_1][v:&{1 1628329469}]
2021/08/07 17:44:29 [设置缓存][item][key:key_2][v:&{2 1628329469}]
2021/08/07 17:44:29 [设置缓存][item][key:key_3][v:&{3 1628329469}]
2021/08/07 17:44:29 [设置缓存][item][key:key_4][v:&{4 1628329469}]
2021/08/07 17:44:29 [设置缓存][item][key:key_5][v:&{5 1628329469}]
2021/08/07 17:44:29 [设置缓存][item][key:key_6][v:&{6 1628329469}]
2021/08/07 17:44:29 [设置缓存][item][key:key_7][v:&{7 1628329469}]
2021/08/07 17:44:29 [设置缓存][item][key:key_8][v:&{8 1628329469}]
2021/08/07 17:44:29 [设置缓存][item][key:key_9][v:&{9 1628329469}]
2021/08/07 17:44:29 缓存中的数据量:10

2021/08/07 17:45:01 [删除过期数据][key:key_0]
2021/08/07 17:45:01 [删除过期数据][key:key_1]
2021/08/07 17:45:01 [删除过期数据][key:key_4]
2021/08/07 17:45:01 [删除过期数据][key:key_5]
2021/08/07 17:45:01 [删除过期数据][key:key_7]
2021/08/07 17:45:01 [删除过期数据][key:key_2]
2021/08/07 17:45:01 [删除过期数据][key:key_3]
2021/08/07 17:45:01 [删除过期数据][key:key_6]
2021/08/07 17:45:01 [删除过期数据][key:key_8]
2021/08/07 17:45:01 [删除过期数据][key:key_9]

2021/08/07 17:45:05 缓存中的数据量:0
2021/08/07 17:45:05 [更新缓存][item][key:key_0][v:&{0 1628329505}]
2021/08/07 17:45:05 [更新缓存][item][key:key_1][v:&{1 1628329505}]
2021/08/07 17:45:05 [更新缓存][item][key:key_2][v:&{2 1628329505}]
2021/08/07 17:45:05 [更新缓存][item][key:key_3][v:&{3 1628329505}]
2021/08/07 17:45:05 [更新缓存][item][key:key_4][v:&{4 1628329505}]
2021/08/07 17:45:05 缓存中的数据量:5

2021/08/07 17:45:37 [删除过期数据][key:key_2]
2021/08/07 17:45:37 [删除过期数据][key:key_0]
2021/08/07 17:45:37 [删除过期数据][key:key_1]
2021/08/07 17:45:37 [删除过期数据][key:key_3]
2021/08/07 17:45:37 [删除过期数据][key:key_4]

*/