package cron

import (
	"day14LuffyAgent/logger"
	"day14LuffyAgent/metrics"
	"day14LuffyAgent/models"
	"day14LuffyAgent/settings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func Collect()  {
	// 检查配置文件是否开启定时采集
	if !settings.Config().Transfer.Enabled{
		return
	}
	if len(settings.Config().Transfer.Addrs) == 0{
		return
	}

	// 根据配置的映射关系数组，进行采集
	for _, v := range metrics.Mappers {
		go collect(int64(v.Interval), v.Fs)
	}
}

// 数据采集
func collect(sec int64, fns []func() []*models.MetricValue)  {
	t := time.NewTicker(time.Second * time.Duration(sec)).C

	for{
		<- t

		hostname, err := settings.Hostname()
		if err != nil{
			continue
		}

		// 返回的数据结构体
		mvs := [] *models.MetricValue{}
		//ignoreMetrics := settings.Config().IgnoreMetrics

		for _, fn :=range fns{
			items := fn()  // 根据方法名称，调用
			if items == nil{
				continue
			}

			log.Println("items:", items)
			logger.StartupInfo(items)
			if len(items) == 0{
				continue
			}

			// 将函数返回的数据结果，保存
			for _, mv := range items{
				mvs = append(mvs, mv)
				// if b, ok := ignoreMetrics[mv.Metric]; ok && b {
				// 	continue
				// } else {
				// 	mvs = append(mvs, mv)
				// }
			}
		}

		// 增加时间戳
		now := time.Now().Unix()
		for j:=0; j < len(mvs); j ++{
			mvs[j].Step = sec
			mvs[j].Endpoint = hostname
			mvs[j].Timestamp = now
		}
		fmt.Println(mvs)

		// 将指定内容ss写到文件中
		out, err := json.Marshal(mvs)
		ss := string(out)
		err1 := ioutil.WriteFile("./output.txt", []byte(ss), 0666)
		log.Println(ss)
		if err1 != nil{
			fmt.Println("ioutil WriteFile error: ", err1)
		}

		// 生产情况下，可以直接发送到监控服务器
		//SendToTransfer(mvs)
	}
}

//func SendToTransfer(metrics []*models.MetricValue)  {
//	if len(metrics) == 0{
//		return
//	}
//
//	// debug模式，打印到控制台
//	debug := settings.Config().Debug
//	if debug{
//		log.Printf("=> <Total=%d> %v\n", len(mvs), metrics[0])
//	}
//
//	var resp models.TransFerResponse
//	rpc.SendMetrics(metrics, &resp)
//
//	if debug{
//		log.Println("<=", &resp)
//	}
//}


