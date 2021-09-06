package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

//  举例 a-z A-Z

type alphaReader struct{
	// 组合reader
	reader io.Reader
}

func (a *alphaReader) Read(p []byte) (int, error)  {
	// 调用io.Reader
	n, err := a.reader.Read(p)
	if err != nil{
		return n, err
	}

	// 自定义filter，过滤p的非字母
	buf := make([]byte, n)
	fmt.Println("初始的buf: ",buf)

	for i:=0; i<n; i++{
		if char :=alphaFilter(p[i]); char != 0{
			buf[i] = char
		}
	}

	// buf临时保存，copy到p
	fmt.Println("过滤后buf: ",buf)
	copy(p, buf)
	return n, nil

}

// 只保留字符串的a-Z
func alphaFilter(r byte) byte{
	if (r>'A' && r<'Z') || (r>='a' && r<='z'){
		return r
	}
	return 0
}

func main() {

	originReader := strings.NewReader("xiaohu 33d33 &()&^&*( rng mihu you")
	reader := alphaReader{
		reader: originReader,
	}

	p1 := make([] byte, 4)
	//p2 := make([]byte, 4)
	for {
		n1, err := reader.Read(p1)
		//_, err := originReader.Read(p2)  // 原生的Reader
		if err == io.EOF{
			break
		}

		log.Printf("[p1][内容:%v]", string(p1[:n1]))
		//log.Printf("[原生的][内容:%v]", string(p2[:n2]))
	}



}


/*
初始的buf:  [0 0 0 0]
过滤后buf:  [120 105 97 111]
初始的buf:  [0 0 0 0]
过滤后buf:  [104 117 0 0]
初始的buf:  [0 0 0 0]
过滤后buf:  [0 100 0 0]
初始的buf:  [0 0 0 0]
过滤后buf:  [0 0 0 0]
初始的buf:  [0 0 0 0]
过滤后buf:  [0 0 0 0]
初始的buf:  [0 0 0 0]
过滤后buf:  [0 0 114 110]
初始的buf:  [0 0 0 0]
过滤后buf:  [103 0 109 105]
初始的buf:  [0 0 0 0]
过滤后buf:  [104 117 0 121]
初始的buf:  [0 0]
过滤后buf:  [111 117]
2021/09/06 08:52:15 [p1][内容:xiao]
2021/09/06 08:52:15 [p1][内容:hu  ]
2021/09/06 08:52:15 [p1][内容: d  ]
2021/09/06 08:52:15 [p1][内容:    ]
2021/09/06 08:52:15 [p1][内容:    ]
2021/09/06 08:52:15 [p1][内容:  rn]
2021/09/06 08:52:15 [p1][内容:g mi]
2021/09/06 08:52:15 [p1][内容:hu y]
2021/09/06 08:52:15 [p1][内容:ou]

Process finished with the exit code 0


*/