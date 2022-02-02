package main

import "fmt"

func main() {

	/**
	函数的返回分为两步：1：函数返回值赋值，2：返回返回值
	如果存在defer函数，那么会字返回返回值的时候进行执行
	*/

	f1()
	f2(1)
	fmt.Println(f3())
	fmt.Println(f4(1))
	fmt.Println(f5(1, 3))

	fundefer()
}

func f1() {
	fmt.Println("调用成功")
}
func f2(a int) {
	fmt.Println("调用成功f2", a)
}
func f3() int {
	return 1
}
func f4(a int) int {
	return 1
}
func f5(a int, b int) (r int) {
	r = a + b
	return
}

func fundefer() {
	fmt.Println("1............")
	defer fmt.Println("2...........")
	fmt.Println("3............")
	defer fmt.Println("4...........")
	defer fmt.Println("5...........")
}
