package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
使用goroutine和channel实现一个计算int64随机数各位数和的程序。
	1：开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	2：开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	3：主goroutine从resultChan取出结果并打印到终端输出
*/
var wg6 = sync.WaitGroup{}

//定义输入和输出结构体
type job struct {
	num int64
}
type resultData struct {
	job    *job
	result int64
}

var jobChans = make(chan *job, 100)
var resultDataChans = make(chan *resultData, 100)

//1：开启一个goroutine循环生成int64类型的随机数，发送到jobChan
func getRandomNumToJobChan(jobChan chan<- *job) {
	defer wg6.Done()
	for {
		//rand.Seed(time.Now().UnixNano())
		x := rand.Int63()
		newJob := &job{
			num: x,
		}
		jobChan <- newJob
		time.Sleep(time.Second)
	}
}

//2：开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
func getJobChanDataToResultChan(jobChan <-chan *job, resultChan chan<- *resultData) {
	defer wg6.Done()
	for {
		jobData := <-jobChan
		sum := int64(0)
		n := jobData.num
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &resultData{
			job:    jobData,
			result: sum,
		}
		resultChan <- newResult
	}
}

func main() {
	wg6.Add(1)
	go getRandomNumToJobChan(jobChans)
	wg6.Add(24)
	for i := 0; i < 24; i++ {
		go getJobChanDataToResultChan(jobChans, resultDataChans)
	}
	for dataChan := range resultDataChans {
		fmt.Println(dataChan.job, "      ", dataChan.result)
	}
	wg6.Wait()
}
