package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	fmt.Scanln(&str)

	fmt.Println("读取到的文件夹路径是：", str)
	str = "D:\\桌面文件"

	fileDir, err := os.OpenFile(str, os.O_RDONLY, os.ModeDir)
	if err !=nil {
		fmt.Println("您填写的目录存在问题！:",err)
		return
	}
	fmt.Println("文件名称是：",fileDir.Name())

}
