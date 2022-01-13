package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(getCurrentDirectory())

	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE| os.O_WRONLY | os.O_RDONLY, 0644)
	if err !=nil {
		fmt.Println("出错误了！")
		return
	}

	fileObj.WriteString("hello 这是我用程序加的字符串！")

	fmt.Println("fileObje",fileObj)

	fileObj.Close()

}

/*
获取程序运行路径
*/
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
