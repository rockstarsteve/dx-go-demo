package main

import (
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



	for  {
		var temp [128]byte
		n, err := filestr.Read(temp[:])

		if err == io.EOF {
			fmt.Println("刚读完")
			return
		}

		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("n:", n)
		fmt.Println("内容是:", string(temp[:n]))

		if n < 128 {
			return
		}

	}

}
