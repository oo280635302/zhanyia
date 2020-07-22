package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zhanyia/src/common"
	"zhanyia/src/must"
	pb "zhanyia/src/proto"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 创建must组件实例
	must.Init()
	mustComponent()
	fmt.Println("run start")
	//fmt.Println(program.GenerateTrees(1))

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

func goSqlOp() {
	//_, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/cs?charset=utf8")
	//if err != nil {
	//	fmt.Println("连接数据库失败", err)
	//	return
	//}
	db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/cs?charset=utf8")
	if err != nil {
		return
	}
	type Student struct {
		Key string
		val int64
	}
	a := &Student{}
	dErr := db.Table("a_students").Where("`key` = '222'").Find(a)
	fmt.Println(dErr.Error, dErr.Error == gorm.ErrRecordNotFound, a)
}

// 必备组件
func mustComponent() {
	// 日志组件
	common.Log = common.AllGlobal["Log"].(*must.Log)
}

// 地图相关
func mapSpace() {
	// 制作新的空白地图
	writeMap := common.MakeMap(5, 3)
	// 日志输出二维图
	common.PrintDoubleMap(writeMap)

	// 填充新的二维图
	common.FullMap(writeMap)
	// 日志输出二维图
	common.PrintDoubleMap(writeMap)

	// 降沉
	common.IconFall(writeMap)
	// 日志输出二维图
	common.PrintDoubleMap(writeMap)

	// 填充新的二维图
	common.FullMap(writeMap)
	// 日志输出二维图
	common.PrintDoubleMap(writeMap)

	// 将图谱转成二维数组
	a := &pb.ClearJoyImage{
		Width:  2,
		Height: 2,
		Body:   []int64{1, 3, 5, 1},
	}
	n := common.ImageToSqArray(a)
	common.PrintDoubleMap(n)
}
