package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"zhanyia/src/common"
	"zhanyia/src/must"
)

func main() {
	// 创建must组件实例
	must.Init()
	common.AllGlobal["Mq"].(*must.Mq).BindReportQueue()

	a := "HelloWord"
	msg, err := json.Marshal(a)
	if err != nil {
		fmt.Println("main Marshal has err", err)
	}

	err = common.AllGlobal["Mq"].(*must.Mq).SendReport(msg)
	if err != nil {
		fmt.Println("main SendReport has err", err)
	}

	// 持久化
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		syscall.SIGINT,
		syscall.SIGILL,
		syscall.SIGFPE,
		syscall.SIGSEGV,
		syscall.SIGTERM,
		syscall.SIGABRT)
	<-signalChan
}
