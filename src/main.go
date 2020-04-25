package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zhanyia/src/common"
	"zhanyia/src/must"
	"zhanyia/src/program"
	pb "zhanyia/src/proto"
)

func main() {
	fmt.Println(common.StackAdd("312312332132132132132132132132121421.412", "2321312321312.4111"))

}

func realMain() {
	rand.Seed(time.Now().UnixNano())

	// 创建must组件实例
	must.Init()
	mustComponent()

	common.LogDeBug("run start")
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

	// 重定向回控制台
	fmt.Println("bye bye")
}

// 必备组件
func mustComponent() {
	// 日志组件
	common.Log = common.AllGlobal["Log"].(*must.Log)
}

// 地图相关
func mapSpace() {
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

	// 将图谱转成二维数组
	a := &pb.ClearJoyImage{
		Width:  2,
		Height: 2,
		Body:   []int64{1, 3, 5, 1},
	}
	n := program.ImageToSqArray(a)
	program.PrintDoubleMap(n)
}
