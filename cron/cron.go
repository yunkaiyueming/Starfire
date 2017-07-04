package cron

import (
	"fmt"
	"time"

	"runtime"

	"github.com/astaxie/beego/logs"
)

func CreateCronJob(d time.Duration, f func() string, fNameInfo ...string) {
	Ticker := time.NewTicker(d)
	runLock := false

	var fName string
	if len(fNameInfo) > 0 {
		fName = fNameInfo[0]
	} else {
		pc, _, _, _ := runtime.Caller(1)
		fName = runtime.FuncForPC(pc).Name()
	}

	for {
		<-Ticker.C
		if !runLock {
			runLock = true

			logs.Alert(fmt.Sprintf("fname:%v,start_time:%s\n", fName, time.Now().Format("2006-01-02 15:04:05")))
			ret := f()
			logs.Alert(fmt.Sprintf("fanme:%v,end_time:%s,ret:%s\n", fName, time.Now().Format("2006-01-02 15:04:05"), ret))

			runLock = false
		}
	}
}
