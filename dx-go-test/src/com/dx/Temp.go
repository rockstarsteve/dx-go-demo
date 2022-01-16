package main

import (
	"fmt"
	"reflect"
)

func main() {

	//类型断言

	var a interface{}

	a = 100

	fmt.Println("a:", a)

	fmt.Println(reflect.TypeOf(a))
	v, ok := a.(int)
	fmt.Println("T a:", v, ok)
	switch tt := a.(type) {
	case string:
		fmt.Println("tt", tt)
	case int:
		fmt.Println("tt", tt)
	}

}
