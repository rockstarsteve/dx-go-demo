package main

import (
	"fmt"
)

func main() {
	fmt.Println("结果是：", 1/10)
	fmt.Println("结果是：", 1%10)

	var a int = 12
	var sum int = 0

	for {
		if a > 0 {
			sum = sum + (a % 10)
			a = a / 10

		} else {
			fmt.Println("a:", a)
			fmt.Println("sum:", sum)
			break
		}
	}

	//for {
	//	time.Sleep(time.Second)
	//	rand.Seed(time.Now().UnixMilli())
	//	r := rand.Int63()
	//	fmt.Println("rrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr:", r)
	//}

}
