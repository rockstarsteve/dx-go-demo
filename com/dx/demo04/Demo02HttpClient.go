package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {

	//post()

	post2()

}

func post() {
	urlValues := url.Values{}
	urlValues.Add("name", "zhaofan")
	urlValues.Add("age", "22")
	resp, _ := http.PostForm("http://www.baidu.com", urlValues)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func get() {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("错误内容是：", err)
		return
	}
	restbyte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("读取返回值为空：", err)
		return
	}
	fmt.Println("r:", string(restbyte))
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
func post2() {

	url := "http://127.0.0.1:8082/get"
	contentType := "application/json"
	data := time.Now().Format("2006010203")

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("string(result)", string(result))
}
