package settings

import (
	"day14LuffyAgent/logger"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func StopHandle()  {
	if _, err := os.Stat(Config().Pid); err == nil{
		// 判断pid文件是否存在
		data, err := ioutil.ReadFile(Config().Pid)
		if err != nil{
			fmt.Println("Not running")
			logger.StartupInfo("Not running")
			os.Exit(1)
		}
		logger.StartupInfo(string(data))

		// 校验pid文件是不是agent进程
		ProcessID, err := strconv.Atoi(string(data)) // 转为int类型
		if err != nil{
			fmt.Println("Unable to read and parse process id found in ", Config().Pid)
			logger.StartupInfo("获取进程ID出现异常")
			os.Exit(1)
		}

		// 删除pid文件
		process, err := os.FindProcess(ProcessID)
		if err != nil{
			fmt.Println("进程id不存在")
		}
		os.Remove(Config().Pid)

		// 停止agent进程
		p_info, _ := processInfo(ProcessID)
		if strings.Contains(p_info, "agentx_osx") && err == nil{
			logger.StartupInfo("正在停止Agent进程:", ProcessID)
			// kill process and exit immediately
			err = process.Kill()
			if err != nil{
				fmt.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
				//logger.FatalInfo("停止进程发生异常:", ProcessID, err)
				os.Exit(1)
			}
			logger.StartupInfo("agent已退出")
			os.Exit(0)
		}

	}else{
		fmt.Println("Not running")
		logger.StartupInfo("进程没有运行")
		os.Exit(1)
	}

}

// 根据pid解析进程信息
func processInfo(pid int) (string, error)  {
	p, err := process.NewProcess(int32(pid))
	if err != nil{
		msg := fmt.Sprintf("Cantnot read process info: %v", err)
		logger.StartupInfo(msg)
		return msg, err
	}
	if v, err := p.Cmdline();err == nil{
		msg := fmt.Sprintf("cmd+args:\t%v\n",v)
		logger.StartupInfo(msg)
		return msg, err
	}
	return "", nil
}