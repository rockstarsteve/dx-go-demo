package main

import "fmt"

func main() {
	a := "hello你好啊1"
	//len是求字节的长度，汉子的字节是占3个byte
	fmt.Println(len(a))
	var b *int = new(int)
	fmt.Println("b:",*b)
	fmt.Println("b?:",b == nil)

}
