package main

import (
	"com/dx/log"
)

func main() {

	//mylog := log.GetMyLog()
	mylog := log.MyLog{}
	mylog.Info("我记录下日志")
	mylog.Debug("我记录下日志")

}
