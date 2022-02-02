package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func main() {
	//INFO GMP模型
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go f1(i)
	}
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go f2(i)
	}
	wg2.Wait()
}

func f1(i int) {
	time.Sleep(time.Second * 5)
	defer wg2.Done()
	fmt.Println("你好", i)
}

func f2(i int) {
	time.Sleep(time.Second * 4)
	defer wg2.Done()
	fmt.Println("你好丫丫丫丫丫丫丫丫丫丫丫丫丫丫丫丫丫丫丫", i)
}
