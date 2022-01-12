package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a int32 = 2
	fmt.Println("a", a)
	fmt.Println("a", reflect.TypeOf(a))
	var b = 3
	fmt.Println("b:", b)
	fmt.Println("a", reflect.TypeOf(b))

}
