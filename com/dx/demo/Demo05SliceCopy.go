package main

import "fmt"

func main() {

	arrray1 := []string{}

	fmt.Println("array1 info :", cap(arrray1), "ararry  len", len(arrray1))

	arrray1 = make([]string, 4)

	fmt.Println("arrray1 is nil?", arrray1 == nil)

	fmt.Println("arrray1", arrray1)
	//fmt.Println("arrray1...:", arrray1...)

	arrray2 := []string{"001", "002", "003", "004", "005", "006", "007"}
	fmt.Println("arrray2", arrray2)

	copy(arrray1, arrray2)

	fmt.Println("array1:", arrray1)
	fmt.Println("array2:", arrray2)

}
