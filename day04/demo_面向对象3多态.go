package main

import "log"

// 体现多态
// 告警通知的函数，根据不同的对象进行通知
//
type notifer interface {
	// 通知方法
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	log.Printf("[普通用户的通知][notify to user :%s]", u.name)
}

type admin struct {
	name string
	age  int
}

func (u *admin) notify() {
	log.Printf("[管理员的通知][notify to user :%s]", u.name)
}


// 多态的统一调用入口
func sendNotify(n notifer) {
	n.notify()
}

func main() {

	u1 := user{
		name:  "小乙",
		email: "xy@qq.com",
	}
	a1 := admin{
		name: "燕青",
		age:  18,
	}
	// 直接调用结构体绑定的方法
	log.Println("直接调用结构体绑定的方法")
	u1.notify()
	a1.notify()
	/*
	   2021/08/11 00:49:15 直接调用结构体绑定的方法
	   2021/08/11 00:49:15 [普通用户的通知][notify to user :小乙]
	   2021/08/11 00:49:15 [管理员的通知][notify to user :燕青]
	*/


	// 体现多态
	log.Println("体现多态")
	sendNotify(&u1)
	sendNotify(&a1)
	/*
	2021/08/11 00:49:15 体现多态
	2021/08/11 00:49:15 [普通用户的通知][notify to user :小乙]
	2021/08/11 00:49:15 [管理员的通知][notify to user :燕青]
	*/


	// 灵魂
	log.Println("多态灵魂承载器")
	ns := make([]notifer,0)
	ns = append(ns, &u1)
	ns = append(ns, &a1)

	for _,n :=range ns{
		n.notify()  // 不管n是什么对象，执行同样的方法
	}

}

/*
2021/08/11 00:53:04 直接调用结构体绑定的方法
2021/08/11 00:53:04 [普通用户的通知][notify to user :小乙]
2021/08/11 00:53:04 [管理员的通知][notify to user :燕青]
2021/08/11 00:53:04 体现多态
2021/08/11 00:53:04 [普通用户的通知][notify to user :小乙]
2021/08/11 00:53:04 [管理员的通知][notify to user :燕青]
2021/08/11 00:53:04 多态灵魂承载器
2021/08/11 00:53:04 [普通用户的通知][notify to user :小乙]
2021/08/11 00:53:04 [管理员的通知][notify to user :燕青]

Process finished with the exit code 0

*/