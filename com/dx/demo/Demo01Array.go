package main

import "fmt"

func main() {

	var arrayb1 [2]bool

	var arrayb2 [3]bool
	arrayb2 = [3]bool{true, false, true}

	arrayb3 := []bool{true, false, true, false}

	arrayb4 := [...]bool{true, false, true, false, true}

	var arrayb5 = [5]bool{2: true, true, 4: true}

	fmt.Println("arrayb1: ", arrayb1)

	fmt.Println("arrayb2: ", arrayb2)

	fmt.Println("arrayb3: ", arrayb3)

	fmt.Println("arrayb4: ", arrayb4)

	fmt.Println("arrayb5: ", arrayb5)


}
