package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
使用goroutine和channel实现一个计算int64随机数各位数和的程序。
	开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	主goroutine从resultChan取出结果并打印到终端输出
*/
var wg5 = sync.WaitGroup{}
var jobChan = make(chan int64, 100)
var resultChan = make(chan int64, 100)

func jobs(jobChan chan int64) {
	rand.Seed(time.Now().Unix())
	r := rand.Int63()
	for {
		jobChan <- r
	}
}

func works(jobChan chan int64) (resultChan chan int64) {
	defer wg5.Done()
	for i := range jobChan {
		fmt.Println("i:", i)
	}
	return resultChan
}

func main() {
	wg5.Add(25)
	go jobs(jobChan)
	for i := 0; i < 24; i++ {
		go works(jobChan)
	}
	wg5.Wait()

}
