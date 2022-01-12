package metrics

import (
	"day14LuffyAgent/models"
	"day14LuffyAgent/settings"
)

// 输出格式化,多个json数据组装为数组，上传
type FuncsAndInterval struct {
	Fs []func() []*models.MetricValue
	Interval int
}

var Mappers []FuncsAndInterval

// 映射关系，先初始化
func BuildMappers()  {
	interval := settings.Config().Transfer.Interval // 采集间隔时间 "interval": 60,

	Mappers = []FuncsAndInterval{
		FuncsAndInterval{
			Fs: []func() []*models.MetricValue{
				SysMetrics,  // 方法名称
			},
			Interval: interval,
		},
	}

}
