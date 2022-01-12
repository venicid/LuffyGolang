package metrics

import (
	"day14LuffyAgent/models"
	"day14LuffyAgent/settings"
)

func SysMetrics() (L []*models.MetricValue) {
	L = append(L, models.GaugeValue("sysinfo.innerip",settings.IP()))
	return
}