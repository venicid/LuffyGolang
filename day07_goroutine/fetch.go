/*
./fetch  https://baidu.com https://taobao.com
// 结果如下：
0.35s   304728  https://baidu.com
0.50s   111189  https://taobao.com

*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan <- string)  {
	start := time.Now()

	resp, err := http.Get(url)  // 这是用来请求的
	if err != nil{
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)  // 获取响应内容
	resp.Body.Close()  // don‘t leak resources
	if err != nil{
		ch <- fmt.Sprintf("读取错误%s:%v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

func main()  {
	ch := make(chan string)

	for _, url := range os.Args[1:]{
		go fetch(url, ch)
	}

	for range os.Args[1:]{
		fmt.Println(<-ch)
	}

}


/*
命令行运行
E:\golang\HelloGolang\day07_goroutine>go run fetch.go  https://baidu.com https://baidu.com
0.61  321190 https://baidu.com
1.00  320839 https://baidu.com

*/