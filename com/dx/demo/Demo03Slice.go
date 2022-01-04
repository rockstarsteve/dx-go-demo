package main

import "fmt"

func main() {

	var arrray []int

	var arrray2 []string

	fmt.Println("nil?:", arrray == nil)

	fmt.Println("arrray", cap(arrray))

	fmt.Println("arrray2", arrray2)

	arrray = []int{1, 2, 3}

	arrray2 = []string{"001", "002", "003", "004", "005", "006", "007"}

	fmt.Println("array size?:", len(arrray))

	arrray3 := arrray2[1:5]

	fmt.Println("array3:", arrray3)

	arrray4 := arrray2[:5]

	fmt.Println("array4:", arrray4, "长度：", len(arrray4),"容量是：",cap(arrray4))

	arrray41 := arrray2[1:]

	fmt.Println("array41:", arrray41)

	arrray5 := arrray2[:]

	fmt.Println("array5:", arrray5, "长度：", len(arrray5),"容量是：",cap(arrray5))

}
