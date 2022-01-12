package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	ret, err := ioutil.ReadFile("./Demo01.go")
	if err != nil {
		fmt.Println("读出错了！")
		return
	}
	fmt.Println("内容是：", string(ret))

}
