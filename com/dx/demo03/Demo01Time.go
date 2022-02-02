package main

import (
	"fmt"
	"time"
)

func main() {

	var t time.Time
	t = time.Now()

	fmt.Println("time:", t)

	fmt.Println("t.year :", t.Year())
	fmt.Println("t.month :", t.Month())
	fmt.Println("t.day :", t.Day())
	y, m, d := t.Date()
	fmt.Println("t.date :", y, m, d)

	fmt.Println("t.Unix :", t.Unix())
	fmt.Println("t.Unix :", t.UnixMilli())
	fmt.Println("t.Unix :", t.UnixMicro())
	fmt.Println("t.Unix :", t.UnixNano())
	str := t.Format("2006-01-02 03:04")
	fmt.Println("str:", str)
	fmt.Println("=================请等待5s钟==============================================================",time.Now())
	time.Sleep(time.Second * 5)
	t2, err := time.Parse("2006-1-2", "2022-01-03")
	if err != nil {
		fmt.Println("解析出错")
		return
	}
	fmt.Println("parse：", t2.Format("2006-01-02 03:04"))

	fmt.Println("================5s钟过去了==============================================================",time.Now())


}
