package main

import (
	"fmt"
	"time"
)

func main() {
	// 告警规则
	want := `
[报警触发类型：%s]
[报警名称：%s]
[级别：%d]
[机器ip列表：%s]
[表达式：%s]
[报警次数：%d]
[报警时间：%s]
`

	alarmContent := fmt.Sprintf(
		want,
		"普罗米修斯",
		"支付接口qps大于1000",
		1,
		"1.1.1.1,2.2.2.2",
		`sum(rate(login_qps[1m]))>1000`,
		2,
		time.Unix(time.Now().Unix(),0).Format("2006-08-02 09:08:04"),
		)

	fmt.Println(alarmContent)


/*
   [报警触发类型：普罗米修斯]
   [报警名称：支付接口qps大于1000]
   [级别：1]
   [机器ip列表：1.1.1.1,2.2.2.2]
   [表达式：sum(rate(login_qps[1m]))>1000]
   [报警次数：2]
   [报警时间：2021-08-02 09:08:10]
*/
}