package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int = 23
	var b float64 = 23.344
	var c string = "23"
	var d bool = true
	var e string = "false"
	var f string = "23.34454"

	fmt.Println("------------------------------------------------", a, b, c, d, e)

	as := strconv.Itoa(a)
	fmt.Println("as", as)

	ci, _ := strconv.Atoi(c)
	fmt.Println("ci", ci)

	fs := strconv.FormatFloat(b, 'E', -1, 32)
	fmt.Println("fs", fs)

	ds := strconv.FormatBool(d)
	fmt.Println("ds", ds)

	eb, _ := strconv.ParseBool(e)
	fmt.Println("eb", eb)

	ef, _ := strconv.ParseFloat(f, 32)
	fmt.Println("ef", ef)

}
