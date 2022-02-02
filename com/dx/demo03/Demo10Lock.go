package main

import (
	"fmt"
	"sync"
)

//var mutex sync.Mutex

var (
	gi   int
	wg7  sync.WaitGroup
	lock sync.Mutex
)

func main() {
	wg7.Add(200000)
	go add()
	go add2()
	wg7.Wait()
	fmt.Println("gi:", gi)
}

func add() {
	for i := 0; i < 100000; i++ {
		lock.Lock()
		gi++
		if i == 100000-1 {
			t := 1 / 0
			fmt.Println("t", t)
		}
		lock.Unlock()
		wg7.Done()
	}
}

func add2() {
	for i := 0; i < 100000; i++ {
		lock.Lock()
		gi++
		lock.Unlock()
		wg7.Done()
	}
}
