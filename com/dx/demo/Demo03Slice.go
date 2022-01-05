package main

import (
	"fmt"
	"reflect"
)

func main() {

	//切片不存值

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

	fmt.Println("array4:", arrray4, "长度：", len(arrray4), "容量是：", cap(arrray4))

	arrray41 := arrray2[1:]

	fmt.Println("array41:", arrray41)

	arrray5 := arrray2[:]

	fmt.Println("array5:", arrray5, "长度：", len(arrray5), "容量是：", cap(arrray5))

	fmt.Println("====================================================")

	var a1 = []int{1, 2, 3, 4, 5}
	fmt.Println("a1:", a1)
	fmt.Println("a1  type:", reflect.TypeOf(a1))

	var a2 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("a2:", a2)

	fmt.Println("a2  type:", reflect.TypeOf(a2))

	fmt.Println("====================================================")

	var b1 = []int{1, 2, 3, 4, 5}
	b2 := b1[:]

	b3 := b2[2:]
	b2 = b3

	fmt.Println("b3:", b3)

	fmt.Println("b1:", b1)

	b2 = append(b2[:1], b2[3:]...)

	fmt.Println("b1:", b1)

	fmt.Println("====================================================")

	var t = make([]string, 5, 10)

	fmt.Println("a:", t, "cap:", cap(t), "len:", len(t))

}
