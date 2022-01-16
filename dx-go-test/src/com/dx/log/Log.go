package log

import (
	"fmt"
	"time"
)

type MyLog struct {
}

func GetMyLog() MyLog {
	return MyLog{}
}

func (l MyLog) Info(info string) {
	fmt.Println(time.Now(), "   info日志   :", info)
}

func (l MyLog) Debug(info string) {
	fmt.Println(time.Now(), "   debug日志   :", info)
}
