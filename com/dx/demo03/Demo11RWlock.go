package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	gli    int = 0
	wg8    sync.WaitGroup
	rwlock sync.RWMutex
	lock1  sync.Mutex
)

func main() {
	start := time.Now().Unix()
	wg8.Add(2000)
	for i := 0; i < 100; i++ {
		go write()
	}
	for i := 0; i < 100; i++ {
		go read()
	}

	wg8.Wait()
	end := time.Now().Unix()
	fmt.Println("消耗时间：", end-start)
}

func read() {
	for i := 0; i < 10; i++ {
		rwlock.RLock()
		//lock1.Lock()
		time.Sleep(time.Millisecond * 1)
		fmt.Println("gli:", gli)
		wg8.Done()
		rwlock.RUnlock()
		//lock1.Unlock()
	}
}

func write() {
	for i := 0; i < 10; i++ {
		rwlock.Lock()
		//lock1.Lock()
		gli++
		time.Sleep(time.Millisecond * 5)
		wg8.Done()
		rwlock.Unlock()
		//lock1.Unlock()
	}

}
