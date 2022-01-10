package settings

import (
	"encoding/json"
	"fmt"
	"log"
	"lufflyagent/logger"
	"lufflyagent/utils"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type TransferConfig struct {
	Enabled  bool     `json: "enabled"`
	Addrs    []string `json: "addrs"`
	Interval int      `json: "interval"`
	Timeout  int      `json: "timeout"`
}

type HttpConfig struct {
	//Enabled bool   `json: "enabled"`
	Listen string `json: "listen"`
}

type CollectorConfig struct {
	IfacePrefix []string `json: "ifacePrefix"`
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

var LocalIp string

func InitLocalIp() {
	Ip := ""
	addr:= []string{"180.101.49.11:80"}
	for _, ip := range addr{
		conn, err := net.DialTimeout("tcp",ip,time.Second*5)
		if err != nil {
			log.Println("get local addr failed !")
		}else{
			Ip = strings.Split(conn.LocalAddr().String(),":")[0]
			fmt.Println(LocalIp)
			conn.Close()
			break
		}

	}
	if Ip != "" {
		LocalIp = Ip
	} else {
		logger.StartupFatal( "get local addr failed !")
	}
}

func IP() string {
	ip := Config().Ip
	if ip != "" {
		return ip
	}

	if len(LocalIp) > 0 {
		ip = LocalIp
	}
	return ip
}

func Hostname() (string, error) {
	hostname := Config().Hostname
	if hostname != "" {
		return hostname, nil
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("ERROR: os.Hostname() fail", err)
		logger.StartupFatal("ERROR: os.Hostname() fail", err)

	}
	return hostname, err
}

func LoadConfiguration() {
	var builder strings.Builder

	str, _ := os.Getwd()
	builder.Write([]byte(str))
	builder.WriteString("/cfg.json")
	cfg := builder.String()

	if !utils.IsExist(cfg) {
		fmt.Println("config file:", cfg, "cfg.json不存在，请运行命令 `mv cfg.example.json cfg.json`")
		logger.StartupFatal("cfg.json不存在，请运行命令 `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := utils.ToTrimString(cfg)
	if err != nil {
		fmt.Println("read config file:", cfg, "fail:", err)
		logger.StartupFatal("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		fmt.Println("parse config file:", cfg, "fail:", err)
		logger.StartupFatal("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	fmt.Println("read config file:", cfg, "successfully")
}
