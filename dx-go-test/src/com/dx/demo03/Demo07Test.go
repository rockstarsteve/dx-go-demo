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
type result struct {
	number int64
	sums   int64
}

var wg5 = sync.WaitGroup{}

var jobChan = make(chan *int64, 100)
var resultChan = make(chan *result, 100)

//开启一个goroutine循环生成int64类型的随机数，发送到jobChan
func jobs(jobChan chan<- *int64) {
	for {
		time.Sleep(time.Second)
		rand.Seed(time.Now().UnixNano())
		r := rand.Int63()
		jobChan <- &r
	}
}

//开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
func works(jobChan <-chan *int64, resultChan chan<- *result) {
	defer wg5.Done()
	for i := range jobChan {
		n := i
		var sum *int64
		sum = new(int64)
		var re *result = &result{}
		re.number = *i
		for {
			if *n > 0 {
				*sum = *sum + (*n % 10)
				*n = *n / 10
			} else {
				re.sums = *sum
				resultChan <- re
				break
			}
		}
	}
}

func main() {
	wg5.Add(25)
	go jobs(jobChan)
	for i := 0; i < 24; i++ {
		go works(jobChan, resultChan)
	}
	//主goroutine从resultChan取出结果并打印到终端输出
	for i := range resultChan {
		fmt.Println("i:", *i)
	}
	wg5.Wait()
}
