package main

import "fmt"

func main() {

	arrray1 := []string{"1", "2", "3"}
	fmt.Println("arrray1", arrray1)
	//fmt.Println("arrray1...:", arrray1...)

	arrray2 := []string{"001", "002", "003", "004", "005", "006", "007"}
	fmt.Println("arrray2", arrray2)

	arrray2 = append(arrray2, "008")
	fmt.Println("arrray2", arrray2)

	arrray2 = append(arrray2, arrray1...)
	fmt.Println("arrray2", arrray2)

	arrray4 := [7]string{}
	//arrray4 = append(arrray2, "008")
	fmt.Println("arrray4", arrray4)

}
