package main

import "fmt"

func main() {
	strArry := [...]string{"小张", "小王", "小李", "小花", "小刘"}

	f1(strArry[0:2])

}

func f1(strs []string) {
	for i, i2 := range strs {
		fmt.Println(i, i2)
	}
}
