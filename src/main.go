package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zhanyia/src/must"
	"zhanyia/src/program"
)

func main() {
	// 创建must组件实例
	must.Init()
	rand.Seed(time.Now().UnixNano())
	//common.AllGlobal["Mq"].(*must.Mq).BindReportQueue()
	//
	//a := "HelloWord"
	//msg, err := json.Marshal(a)
	//if err != nil {
	//	fmt.Println("main Marshal has err", err)
	//}
	//
	//err = common.AllGlobal["Mq"].(*must.Mq).SendReport(msg)
	//if err != nil {
	//	fmt.Println("main SendReport has err", err)
	//}
	// 制作新的空白地图
	writeMap := program.MakeMap(5, 3)
	// 日志输出二维图
	program.PrintDoubleMap(writeMap)

	// 填充新的二维图
	program.FullMap(writeMap)
	// 日志输出二维图
	program.PrintDoubleMap(writeMap)

	// 降沉
	program.IconFall(writeMap)
	// 日志输出二维图
	program.PrintDoubleMap(writeMap)

	// 填充新的二维图
	program.FullMap(writeMap)
	// 日志输出二维图
	program.PrintDoubleMap(writeMap)

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
