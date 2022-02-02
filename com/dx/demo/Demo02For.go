package main

import "fmt"

func main() {

	var array = []int{2, 3, 4, 6, 2, 4}

	var total = 0
	fmt.Println(array)

	fmt.Println("count:", len(array))

	fmt.Println("total:", total)



	for i := 0; i < len(array); i++ {
		fmt.Println("当前数据是：", array[i])
	}

	for i, i2 := range array {
		fmt.Println("i:", i, "i2:", i2)
	}

	for _, i := range array {
		fmt.Println("i:", i)
	}

}
