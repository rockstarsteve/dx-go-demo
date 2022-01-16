package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	rand.Seed(time.Now().Unix())
	a1 := rand.Int()
	a2 := rand.Intn(10)
	fmt.Println("rand a1", a1)
	fmt.Println("rand a2", a2)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go hello(i)
	}
	//等待
	wg.Wait()

}

func hello(i int) {
	defer wg.Done()
	fmt.Println("你好：", i)
}
