package models

import "fmt"

// 数据采集后，格式定义，透穿给监控服务器
type MetricValue struct {
	Endpoint string `json:"endpoint"`
	Metric string `json:"metric"`
	Value interface{} `json:"value"`
	Step int64 `json:"step"`
	Type string `json:"type"`
	Tags string `json:"tags"`
	Timestamp int64 `json:"timestamp"`
}

func (metricValue *MetricValue) String() string {
	return fmt.Sprintf(
		"<EndPoint:%s, Metric: %s , Type: %s, Tags: %s, Step: %s, Time:%d, Value:%v>",
		metricValue.Endpoint,
		metricValue.Metric,
		metricValue.Type,
		metricValue.Tags,
		metricValue.Step,
		metricValue.Timestamp,
		metricValue.Value,
		)
}

