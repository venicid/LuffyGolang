package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
交互式社交机器人
1、打印index
2、用户输入1 注册
	打印注册要求
	用户输入用户名密码
	保存成功
	跳出到步骤1
3、用户输入2 登录
	打印登录要求
	用户输入用户名密码
		读取文件，校验
		失败，请用户重新输入，失败超过3次，跳到步骤1
		成功，到主界面
4、用户输入3 退出
	是否退出
	Y 退出
	N 返回步骤1
**/
func main()  {


	indexStr := `
				寻找伴侣，挚爱一生
					1. Register
					2. Login
					3. Exit
				请选择:
				`
	registStr := `注册界面
					请输入用户名密码，用;号隔开，比如alex;33'
					`
	for {
		fmt.Println(indexStr)

		f := bufio.NewReader(os.Stdin)
		input, _ := f.ReadString('\n')  // \n停止输入
		fmt.Println(input)

		//inputSlice := strings.Split(input, "\n")
		//fmt.Println(inputSlice)
		fmt.Printf("%T\n", input)
		num,_ := strconv.Atoi(input)
		fmt.Println(num)
		if input == "1" {
			fmt.Println(registStr)
		}



	}




}