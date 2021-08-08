package main

import "log"


// 实现1+...+100
// 从低位加到高位
func addSmallToBig(num int)int{
	if num ==100{
		return num
	}
	return num + addSmallToBig(num+1)
}

// 从高位加到低位
func addBigToSmall(num int)int{
	if num ==1{
		return num
	}
	return num + addBigToSmall(num-1)
}


func main()  {

	log.Printf("从低位加到高位 ：%d", addSmallToBig(1))
	log.Printf("从高位加到低位 ：%d", addBigToSmall(100))

}

/*
2021/08/08 18:35:29 从低位加到高位 ：5050
2021/08/08 18:35:29 从高位加到低位 ：5050
*/