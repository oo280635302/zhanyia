package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zhanyia/src/must"
	"zhanyia/src/program"
	pb "zhanyia/src/proto"
)

func main() {
	// 创建must组件实例
	must.Init()
	rand.Seed(time.Now().UnixNano())

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
