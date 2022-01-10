package settings

import (
	"day14LuffyAgent/logger"
	"day14LuffyAgent/utils"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type TransferConfig struct {
	Enabled bool `json:"enabled"`
	Addrs []string `json:"addrs"`
	Interval int `json:"interval"`
	Timeout int `json:"timeout"`
}

type HttpConfig struct {
	Listen string `json:"listen"`
}

type CollectorConfig struct {
	IfacePrefix []string `json:"iface_prefix"`
}

type GlobalConfig struct {
	Debug       bool               `json:"debug"`
	Hostname    string             `json:"hostname"`
	Ip          string             `json:"ip"`
	Pid         string             `json:"pid"`
	Logfile     string             `json:"logfile"`
	Transfer    *TransferConfig    `json:"transfer"`
	Http        *HttpConfig        `json:"http"`
	Collector   *CollectorConfig   `json:"collector"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}
// 根据ping某个ip，获取当前的ip地址
var LocalIp string

func InitLocalIp()  {
	Ip := ""
	addr := []string{"180.101.49.11:80"}
	for _,ip := range addr{
		conn, err := net.DialTimeout("tcp", ip ,time.Second*5)
		if err != nil{
			fmt.Println("get local  addr failed")
			//logger.FatalInfo("get local  addr failed")
		}else{
			Ip = strings.Split(conn.LocalAddr().String(), ":")[0]
			logger.StartupInfo(LocalIp)
			conn.Close()
			break
		}


	}
	if Ip != ""{
		LocalIp = Ip
	}else{
		fmt.Println("get local  addr failed")
		//logger.FatalInfo("get local addr failed")
	}
	logger.StartupInfo("当前ip地址为：", LocalIp)

}

func IP() string {
	ip := Config().Ip
	if ip != ""{
		return ip
	}
	if len(LocalIp) >0 {
		ip = LocalIp
	}
	return ip
}

// 修改配置文件中的hostname
func Hostname() (string, error) {
	hostname := Config().Hostname
	if hostname != ""{
		return hostname, nil
	}

	hostname, err := os.Hostname()
	if err != nil{
		fmt.Println("get hostname error")
	}
	return hostname, nil

}

func LoadConfiguration()  {

	// 判断文件路径是否正确
	var builder strings.Builder
	str, _ := os.Getwd()
	builder.Write([]byte(str))
	builder.WriteString("/cfg.json")
	cfg := builder.String()

	if !utils.IsExist(cfg){
		fmt.Println("config file isnot exists")
		//logger.FatalInfo("config file isnot exists")
	}
	ConfigFile = cfg

	// 读取文件中的内容
	configContent, err := utils.ToTrimString(cfg)
	if err != nil{
		fmt.Println("read config files:", cfg, "fail:", err)
		//logger.FatalInfo("read config files:", cfg, "fail:", err)
	}

	// 序列化配置文件内容
	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err !=nil{
		fmt.Println("parse config file:", cfg, "fail:", err)
		//logger.FatalInfo("parse config file:", cfg, "fail:", err)
	}

	// 添加锁，防止共同修改
	lock.Lock()
	defer lock.Unlock()

	config = &c
	fmt.Println("read config file:",cfg, "successfully")
	logger.StartupInfo("read config file:",cfg, "successfully")
}

