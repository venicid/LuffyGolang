package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"log"
	"time"
)

/*
生产上用的 web缓存应用

维护用户信息的模块
在mysql中有1张user表

正常情况下，用orm，gorm，xorm去db中查询
查询qps很高，为了性能会加缓存
(更新不会太频繁)，说明在一定时间内，获取到旧数据也能容忍
*/

type user struct {
	Name string
	Email string
	Phone int64
}

var (
	DefaultInterval = time.Minute * 1
	UserCache = cache.New(DefaultInterval, DefaultInterval)
)

// http请求接口查询，mock模拟
func HttpGetUser(name string) user  {
	u := user{
		Name: name,
		Email: "qq.com",
		Phone: time.Now().Unix(),
	}
	return u
}

// 最外层调用函数
// 优先去本地缓存中查，有就返回
// 没有再去远端查询，远端用http请求表示
func GetUser(name string) user  {

	// 消耗 0.1cpu 0.1M内存 0.1秒返回
	res,found := UserCache.Get(name)
	if found{
		log.Printf("[本地缓存中找到了对应的用户][name：%v][value:%v]", name, res.(user))
		return res.(user)

	}else{
	// 消耗 1cpu 10M内存 3秒返回
		// 本地没有，但是从远端拿到了最新的数据
		// 更新本地缓存 ,我种树，其他人乘凉
		res := HttpGetUser(name)
		UserCache.Set(name, res, DefaultInterval)
		log.Printf("[本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：%v][value:%v]", name, res)
		return res
	}
}

// 查询方法
func queryUser()  {
	for i:=0 ;i<10;i++{
		userName  := fmt.Sprintf("user_name%d", i)
		GetUser(userName)
	}
}

func main()  {
	log.Printf("第1次query_user")
	queryUser()

	log.Printf("第2次query_user")
	queryUser()
	queryUser()

	time.Sleep(61*time.Second)
	log.Printf("第3次query_user")
	queryUser()
}


/*
2021/08/07 18:38:16 第1次query_user
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name0][value:{user_name0 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name1][value:{user_name1 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name2][value:{user_name2 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name3][value:{user_name3 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name4][value:{user_name4 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name5][value:{user_name5 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name6][value:{user_name6 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name7][value:{user_name7 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name8][value:{user_name8 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name9][value:{user_name9 qq.com 1628332696}]
2021/08/07 18:38:16 第2次query_user
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name0][value:{user_name0 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name1][value:{user_name1 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name2][value:{user_name2 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name3][value:{user_name3 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name4][value:{user_name4 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name5][value:{user_name5 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name6][value:{user_name6 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name7][value:{user_name7 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name8][value:{user_name8 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name9][value:{user_name9 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name0][value:{user_name0 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name1][value:{user_name1 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name2][value:{user_name2 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name3][value:{user_name3 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name4][value:{user_name4 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name5][value:{user_name5 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name6][value:{user_name6 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name7][value:{user_name7 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name8][value:{user_name8 qq.com 1628332696}]
2021/08/07 18:38:16 [本地缓存中找到了对应的用户][name：user_name9][value:{user_name9 qq.com 1628332696}]

2021/08/07 18:39:17 第3次query_user
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name0][value:{user_name0 qq.com 1628332757}]
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name1][value:{user_name1 qq.com 1628332757}]
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name2][value:{user_name2 qq.com 1628332757}]
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name3][value:{user_name3 qq.com 1628332757}]
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name4][value:{user_name4 qq.com 1628332757}]
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name5][value:{user_name5 qq.com 1628332757}]
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name6][value:{user_name6 qq.com 1628332757}]
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name7][value:{user_name7 qq.com 1628332757}]
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name8][value:{user_name8 qq.com 1628332757}]
2021/08/07 18:39:17 [本地缓存中没找到对应的用户，去远端查询获取到了，塞入缓存中][name：user_name9][value:{user_name9 qq.com 1628332757}]

*/