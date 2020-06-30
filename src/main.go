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
	"zhanyia/src/program"
	pb "zhanyia/src/proto"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AStudent struct {
	Key string
	Val int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 创建must组件实例
	must.Init()
	mustComponent()

	fmt.Println(program.NextGreaterElements([]int{1, 2, 2, 1}))

	fmt.Println("run start")
	time.Sleep(time.Second * 2)
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
	db, err := gorm.Open("mysql", "root:123@tcp(127.0.0.1:3306)/cs?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	}
	//db.LogMode(true)
	//a := &AStudent{}
	db.Exec("update `students` set `key` = `key`-"+fmt.Sprintf("%0.1f", 3.54)+" where val = ? ", 120)
	e := db.Table("students").Where("`key` = 1").Update("val", 120, "key", 123)
	fmt.Println("1", e.RowsAffected)
	fmt.Println(e.Error)
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
