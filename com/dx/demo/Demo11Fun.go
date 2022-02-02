package main

import (
	"fmt"
	"reflect"
)

type Obj struct {
	a int
	b string
}

func main() {
	a := [...]int{1, 2, 3, 4}
	a[0] = 3

	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	var o1 Obj

	fmt.Println(reflect.TypeOf(o1))

}
