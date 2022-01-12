package main

import (
	"fmt"
	"reflect"
)

func main() {
	//  &取地址、*根据地址取值   十进制：824633761944
	n := 29

	fmt.Println("&n:", &n)
	fmt.Println("&n:", reflect.TypeOf(&n))
	b := &n
	fmt.Println("&&n:", reflect.TypeOf(&b))
	c := &b
	fmt.Println("&&n:", reflect.TypeOf(&c))

	fmt.Println("*&n:", *&n)
	fmt.Println("&*&n::", reflect.TypeOf(*&n))

	fmt.Println("====================================================")

	//var s *int
	//*s = 0xc00000a098
	//
	//fmt.Println("*s:", s)

	var s = new(int)
	*s = 100
	fmt.Println("s", s)
	fmt.Println("s", *s)

}
