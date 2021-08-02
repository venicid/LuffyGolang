package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main()  {

	var a int8
	var b int16
	var c int32
	var d int64

	fmt.Printf("[整型:%T, 取值范围是：%d~%d][字节大小：%d]\n", a, math.MinInt8, math.MaxInt8, unsafe.Sizeof(a))
	fmt.Printf("[整型:%T, 取值范围是：%d~%d][字节大小：%d]\n", b, math.MinInt16, math.MinInt16, unsafe.Sizeof(b))
	fmt.Printf("[整型:%T, 取值范围是：%d~%d][字节大小：%d]\n", c, math.MinInt32, math.MinInt32, unsafe.Sizeof(c))
	fmt.Printf("[整型:%T, 取值范围是：%d~%d][字节大小：%d]\n", d, math.MinInt64, math.MinInt64, unsafe.Sizeof(d))

/*
   [整型:int8, 取值范围是：-128~127][字节大小：1]
   [整型:int16, 取值范围是：-32768~-32768][字节大小：2]
   [整型:int32, 取值范围是：-2147483648~-2147483648][字节大小：4]
   [整型:int64, 取值范围是：-9223372036854775808~-9223372036854775808][字节大小：8]
*/

	var m uint16
	m = 499
	e := uint8(m)
	f := uint32(m)

	fmt.Println(m)
	fmt.Printf("[unit16==>uint8]: e=%d\n", e)
	fmt.Printf("[unit16==>uint32]: f=%d\n", f)

	/*
	   [unit16==>uint8]: e=243
	   [unit16==>uint32]: f=499
	*/

}