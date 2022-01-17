package main

import (
	"fmt"
	LFCron "github.com/robfig/cron/v3"
	demolog "github.com/sirupsen/logrus"
	"os/exec"
	"time"
)

var XXCron *LFCron.Cron = LFCron.New(LFCron.WithParser(
	LFCron.NewParser(
		LFCron.Minute | LFCron.Hour | LFCron.Dom | LFCron.Month | LFCron.Dow ),
	))


func init()  {
	fmt.Println("test")
	XXCron.Start()
}

type TestJob struct {
	task_id string
	cmd string
}

func (testJob TestJob) Run()  {
	dt := time.Now().Format("2006-01-02 15:11:11") // ISO 8601标准
	fmt.Println("运行中...", dt)
	demolog.Println("JobId", testJob.task_id)

	// 通过go实现执行命令
	c1, err := exec.Command("bash", "-c", testJob.cmd).Output()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(c1))
}

func CronDemo() LFCron.EntryID {
	spec := "* * * * *"
	id, _ :=XXCron.AddJob(spec, &TestJob{
		"333",
		"df -h",
	})

	return id
}

func main()  {

	// 定时命令执行
	CronDemo()

	select {

	}
}
