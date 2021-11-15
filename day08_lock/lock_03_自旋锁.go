package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type spinLock int32

func (s1 *spinLock) Lock()  {

	// 自旋锁是CSA算法的某种实现之一
	for !atomic.CompareAndSwapInt32((*int32)(s1), 0 , 1){
		runtime.Gosched()
	}
}

func (s1 *spinLock) Unlock()  {
	atomic.StoreInt32((*int32)(s1), 0)
}

func NewSpinLock() sync.Locker {
	var lock spinLock
	return &lock
}
