package settings

import (
	"day14LuffyAgent/logger"
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
	"os/exec"
	"strconv"
	"strings"
)
// Processes为所有当前正在运行的进程返回一段指向Process结构的指针。
func FindCheckProcess() ([] *process.Process, error) {
	return process.Processes()
}

// 返回当前运行进程的所有Name
func getAgentList() []string  {
	list_, _ := FindCheckProcess()
	agentNum := []string{}
	for _,v := range list_ {
		ss, _ := v.Name()
		if strings.Contains(os.Args[0], ss) {
			//if ss == "agentx_osx" {
			agentNum = append(agentNum, ss)
		}
	}
	return agentNum
}



// 保存pid文件
func savePID(pid int)  {
	file, err := os.Create(Config().Pid)
	if err != nil{
		fmt.Println("没有pid file：%v\n", err)
		//logger.FatalInfo("没有pid file：%v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))
	if err != nil{
		fmt.Printf("没有pid file： %v\n", err)
		os.Exit(1)
	}
	file.Sync()
}

func StartHandle()  {
	// 检查agent进程是否已经运行
	if _,err := os.Stat(Config().Pid); err == nil{
		fmt.Println("已经运行或pid文件已经存在")
		logger.StartupInfo("启动失败，已经运行或pid已经存在")
		os.Exit(1)
	}

	// 判断是否运行多个
	agentList := getAgentList()
	if len(agentList) > 1 {
		fmt.Println(len(agentList))
		fmt.Println("agent已启动，请不要重复运行")
		os.Exit(1)
	}

	// 启动agent，打印msg，保存pid文件
	fmt.Println("开始启动main")
	fmt.Println(os.Args)
	cmd := exec.Command(os.Args[0], "main")
	//cmd := exec.Command("./agentx_osx", "main")  // windows命令错误
	cmd.Start()

	logger.StartupInfo("配置初始化成功...")

	fmt.Println("版本：", GetVersion())
	fmt.Println("进程已启动 PID is :")
	//fmt.Println(cmd)
	//fmt.Println(cmd.Process)
	fmt.Println(cmd.Process.Pid)
	logger.StartupInfo("agent进程已经启动")
	logger.StartupInfo("IP: ", IP())

	savePID(cmd.Process.Pid)
	os.Exit(0)
}
