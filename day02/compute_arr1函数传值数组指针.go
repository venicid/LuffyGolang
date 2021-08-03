package main

import (
	"fmt"
	"unsafe"
)


func bigArr(arr [1e6]int){
	fmt.Printf("[数组复制：大小%d字节]\n", unsafe.Sizeof(arr))
}

func bigArrPoint(arr *[1e6]int)  {
	fmt.Printf("[数组指针复制：大小%d字节]\n", unsafe.Sizeof(arr))
}

func main() {

	var arr [1e6]int
	bigArr(arr)
	bigArrPoint(&arr)

	/*
	[数组复制：大小8000000字节]
	[数组指针复制：大小8字节] // 指针是无符号的int64位，故是8字节
	// 64位CPU的地址总线可寻址范围 为 0 ~ 2^64-1，代表的存储单元编号的范围
	指针是uint64，故是8字节
	*/
}