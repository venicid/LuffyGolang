package models

import "strings"

// 返回mertic的json数据
func NewMetricValue(metric string, val interface{}, dataType string, tags ...string) *MetricValue  {
	mv := MetricValue{
		Metric:    metric,
		Value:     val,
		Type:      dataType,
	}

	size := len(tags)
	if size>0{
		mv.Tags = strings.Join(tags, ",")
	}

	return &mv

}

// GaugeValue进行组装数据
func GaugeValue(metric string, val interface{}, tags...string) *MetricValue  {
	return NewMetricValue(metric, val, "GAUGE", tags...)
}