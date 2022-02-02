package main

import "fmt"

func main() {

	fmt.Println("")
	f111()
	f222()
	f333()

}

func f111() {
	fmt.Println("1:")
}
func f222() {
	defer func() {
		fmt.Println("a。。。")
	}()
	defer func() {
		fmt.Println("在退出系统时候进行资源释放。。。")
		err := recover()
		fmt.Println("恢复时候产生的错误是：", err)
		fmt.Println("不知道是否恢复了。。")
	}()
	defer func() {
		fmt.Println("b。。。")
	}()
	panic("出了不能恢复的错误")
	fmt.Println("2:")
}
func f333() {
	fmt.Println("3:")
}
