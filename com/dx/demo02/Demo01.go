package main

import "fmt"

func main() {

	//4! = 4* 3*2*1 = 24
	//1! = 1
	fmt.Println("4!:", getJ(4))
	fmt.Println("1!:", getJ(1))
}

func getJ(i int) int {
	if i == 1 {
		return 1
	} else {
		return getJ(i-1) * i
	}
}
