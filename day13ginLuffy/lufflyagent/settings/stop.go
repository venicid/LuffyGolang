package settings

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"io/ioutil"
	"lufflyagent/logger"
	"os"
	"strconv"
	"strings"
)

func processInfo(pid int) (string, error) {
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		msg := fmt.Sprintf("Cannot read process info: %v", err)
		logger.StartupInfo(msg)
		return msg, err
	}
	if v, err := p.Cmdline(); err == nil {
		msg := fmt.Sprintf("cmd+args:\t%v\n", v)
		logger.StartupInfo(msg)
		return msg, err
	}
	return "", nil
}

func StopHandle() {
	//savePID(14014)
	//dd, err := getProcessRunningStatus(77937)
	//logger.StartupFatal(dd, err)

	if _, err := os.Stat(Config().Pid); err == nil {
		data, err := ioutil.ReadFile(Config().Pid)
		if err != nil {
			fmt.Println("Not running")
			logger.StartupDebug("Not running")
			os.Exit(1)
		}
		logger.StartupDebug(string(data))
		ProcessID, err := strconv.Atoi(string(data))
		if err != nil {
			fmt.Println("Unable to read and parse process id found in ", Config().Pid)
			logger.StartupDebug("获取进程ID出现异常")
			os.Exit(1)
		}
		process, err := os.FindProcess(ProcessID)
		if err != nil {
			fmt.Printf("Unable to find process ID [%v] with error %v \n", ProcessID, err)
			logger.StartupDebug("进程ID不存在")
			os.Exit(1)
		}
		// remove PID file
		os.Remove(Config().Pid)
		p_info, _ := processInfo(ProcessID)
		if strings.Contains(p_info, "agentx_osx") && err == nil {
			logger.StartupInfo("正在停止Agent进程:", ProcessID)
			// kill process and exit immediately
			err = process.Kill()
			if err != nil {
				fmt.Printf("Unable to kill process ID [%v] with error %v \n", ProcessID, err)
				logger.StartupDebug("停止进程发生异常:", ProcessID, err)
				os.Exit(1)
			}
			logger.StartupFatal()
		}
		logger.StartupInfo("agent已退出")
		os.Exit(0)
	} else {
		fmt.Println("Not running.")
		logger.StartupDebug("进程没有运行")
		os.Exit(1)
	}
}
