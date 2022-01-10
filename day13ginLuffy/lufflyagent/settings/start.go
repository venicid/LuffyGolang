package settings

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"lufflyagent/logger"
	"os"
	"os/exec"
	"strconv"
)

func FindCheckProccess() ([]*process.Process, error) {
	return process.Processes()
}

func a() []string {
	list_, _ := FindCheckProccess()
	agentNums := []string{}
	for _, v := range list_ {
		ss, _ := v.Name()
		if ss == "agentx_osx" {
			agentNums = append(agentNums, ss)
		}
	}
	return agentNums
}

func savePID(pid int) {

	file, err := os.Create(Config().Pid)
	if err != nil {
		fmt.Printf("没有pid file : %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.WriteString(strconv.Itoa(pid))

	if err != nil {
		fmt.Printf("没有pid file : %v\n", err)
		os.Exit(1)
	}

	file.Sync()

}

func StartHandle() {
	// 检查agent进程是否已经运行.
	if _, err := os.Stat(Config().Pid); err == nil {
		fmt.Println("已经运行或pid文件已存在")
		logger.StartupInfo("启动失败，已经运行或pid文件已存在")
		os.Exit(1)
	}

	agentList := a()
	if len(agentList) > 1 {
		fmt.Println(len(agentList))
		os.Exit(1)
	}

	fmt.Println("开始启动main")
	fmt.Println(os.Args)
	cmd := exec.Command(os.Args[0], "main")
	cmd.Start()

	logger.StartupInfo("配置初始化成功...")
	fmt.Println("版本:", GetVersion())
	fmt.Println("进程已启动 PID is : ", cmd.Process.Pid)
	logger.StartupInfo("agent进程已启动")

	savePID(cmd.Process.Pid)
	os.Exit(0)
}
