package main

import (
	"fmt"
	"sync"
	"time"
)

var wg3 = sync.WaitGroup{}

func recv(c chan int) {
	defer wg3.Done()
	fmt.Println("接收了。。。")
	time.Sleep(time.Second * 3)
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	wg3.Add(1)
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	fmt.Println("开启了。。。")
	time.Sleep(time.Second * 5)
	ch <- 10
	time.Sleep(time.Second * 5)
	fmt.Println("发送成功")
	wg3.Wait()
}
