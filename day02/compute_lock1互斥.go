package main

import (
	"log"
	"sync"
	"time"
)

// hcMuTex 它是一个互斥锁
var HcMutex sync.Mutex

func runMutex(id int)  {
	log.Printf("[任务id：%d][尝试获取锁]", id)
	HcMutex.Lock()
	log.Printf("[任务id：%d][get到了锁]", id)
	time.Sleep(5*time.Second)
	HcMutex.Unlock()
	log.Printf("[任务id：%d][干完了活，释放锁]", id)

}

func runHcLock()  {
	go runMutex(1)
	go runMutex(2)
	go runMutex(3)

}

func main(){
	runHcLock()
	time.Sleep(6*time.Minute)
}

/*
2021/08/05 09:10:05 [任务id：3][尝试获取锁]
2021/08/05 09:10:05 [任务id：3][get到了锁]
2021/08/05 09:10:05 [任务id：2][尝试获取锁]
2021/08/05 09:10:05 [任务id：1][尝试获取锁]

2021/08/05 09:10:10 [任务id：2][get到了锁]
2021/08/05 09:10:10 [任务id：3][干完了活，释放锁]

2021/08/05 09:10:15 [任务id：1][get到了锁]
2021/08/05 09:10:15 [任务id：2][干完了活，释放锁]

2021/08/05 09:10:20 [任务id：1][干完了活，释放锁]
*/