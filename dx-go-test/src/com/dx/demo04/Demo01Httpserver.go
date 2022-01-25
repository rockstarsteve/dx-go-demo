package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/get", f1)
	http.ListenAndServe("127.0.0.1:8082", nil)
	fmt.Println("服务关闭。。。")
}

func f1(w http.ResponseWriter, r *http.Request) {

	fmt.Println("r:",r.URL)
	fmt.Println("r.URL.",r.URL.Query())
	fmt.Println("r:",r.Method)
	fmt.Println("r:",r.Body)


	//str := []byte("你好")
	f, err := ioutil.ReadFile("D:\\桌面文件\\保存考试信息.txt")
	if err != nil {
		fmt.Println("读取文件有误：", err)
		return
	}
	w.Write(f)
}
