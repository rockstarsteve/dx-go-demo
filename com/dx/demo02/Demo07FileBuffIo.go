package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	filestr, err := os.Open("./Demo01.go")
	if err != nil {
		fmt.Println("filestr:", filestr)
		fmt.Println("filestr:", reflect.TypeOf(filestr))
		return
	}

	defer filestr.Close()

	reder := bufio.NewReader(filestr)

	for {
		line,err:=reder.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完或者有问题")
			break
		}
		if err!=nil{
			fmt.Println("文件读出错：",err)
			return
		}
		fmt.Println(line)
	}

}
