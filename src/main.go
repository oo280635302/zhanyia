package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
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
	//fmt.Println(program.LetterCombinations("23"))
	cs()
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

func cs() {
	inter := make(map[string]interface{})

	// 创建请求
	req, err := http.NewRequest("GET", "http://manage-test.rvaka.cn/v1/manage/open/word/order", strings.NewReader("page=1&size=1"))
	if err != nil {
		fmt.Println("获取工单详情 http newRequest has err:", err)
		return
	}
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZW1wIiwicGF5bG9hZCI6IjQ4ODQ0MTk5ODk1MjQzNWRhODk1Mjg2NjMyZTgyZjQwIiwiaXNzIjoi5Zub5bed5bCP5ZKW56eR5oqA5pyJ6ZmQ5YWs5Y-4IiwiaWF0IjoxNTk1OTMwMTM5LCJleHAiOjE1OTY1MzQ5Mzl9.cwR0gHkLyD6WWkVYvTPu7Mop6VgrxK32wnXmNwdA9As")

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("do has err:", err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || len(body) == 0 {
		fmt.Println("获取工单列表 http 读取resp.body失败 err:", err)
		return
	}
	err = json.Unmarshal(body, &inter)
	if err != nil {
		fmt.Println("unmarl err:", err)
		return
	}
	fmt.Println(inter)
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
