package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
)

func main() {
	fmt.Println("绝对路径是：", getCurrentAbPathByCaller())
	//INFO 定义源读取文件，和追加写文件
	var srcFile, witeFile string
	srcFile = "D://test.txt"
	witeFile = "D://append.txt"

	//以创建形式打开追加的文件
	appendFileConten, err := os.OpenFile(witeFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer appendFileConten.Close()
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}
	//以创建形式打开要读取的文件
	srcFileConten, err := os.OpenFile(srcFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer srcFileConten.Close()
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}

	for {
		var bufs = [3]byte{}
		n, err := srcFileConten.Read(bufs[:])
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}
		if err != nil {
			fmt.Println("文件读错")
			return
		}

		fmt.Print(string(bufs[0:n]))
		appendFileConten.Write(bufs[0:n])
	}

	fmt.Println(witeFile)
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() (abPath string) {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return
}
