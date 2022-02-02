package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go hello(i)
	}
	//等待
	wg.Wait()

}

func hello(i int) {
	defer wg.Done()
	rand.Seed(time.Now().Unix())
	var a1 int = rand.Intn(5)
	time.Sleep(time.Second * time.Duration(a1))
	fmt.Println("你好：", i)
}
