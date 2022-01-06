package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func main() {
	h := "汉"

	for _, t := range h {
		fmt.Println("h:", unicode.Is(unicode.Han, t))
	}
	y := "y"
	for _, t := range y {
		fmt.Println("y:", unicode.Is(unicode.Han, t))
	}


	d:=strings.Split("asdflkjsakldf  alsjfd sadflkj asdf sadflj"," ")
	fmt.Println(reflect.TypeOf(d))

	fmt.Println(d)

}
