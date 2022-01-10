package metrics

import (
	"lufflyagent/models"
	"lufflyagent/settings"
)

func SysMetrics() (L []*models.MetricValue) {
	L = append(L, models.GaugeValue("sysinfo.innerip", settings.IP()))
	return
}