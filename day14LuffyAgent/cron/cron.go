package cron

import (
	Cron "github.com/robfig/cron/v3"
	"go.etcd.io/etcd/clientv3"
	"sync"
)

var(
	XCron *Cron.Cron = Cron.New(Cron.WithParser(
		Cron.NewParser(
			Cron.Minute | Cron.Hour | Cron.Dom | Cron.Month | Cron.Dow)))

	client *client3.Client
	initEtcd sync.Once
)

func init()  {
	XCron.Start()
}

func initScheduler() error {
	if !settings.Config()

}
