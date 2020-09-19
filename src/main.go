package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 创建must组件实例
	must.Init()
	mustComponent()
	fmt.Println("run start")
	//csRedis()
	csHttp()
	//common.UnmarshalPb2Url(&pb.ClearJoyImage{Width:123})
	//fmt.Println(program.LetterCombinations("23"))

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

func csHttp() {
	inter := make(map[string]interface{})
	str := "phone=1234444444&password=123&app_key=488441998952435da895286632e82f40&timeStamp=1597636234"
	req, err := http.NewRequest("PUT", "https://v5preapp.rvaka.cn/passenger/api/v1/changePassword", strings.NewReader(str))
	if err != nil {
		fmt.Println("修改工单 http newRequest has err:", err)
		return
	}

	req.Header.Set("Authorization", "76AC42E5DB74B88A7512437BBE85FA5BCE7AC890")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送请求
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("新增工单 http 请求失败 err:", err, str)
		return
	}
	defer resp.Body.Close()

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil || len(body) == 0 {
		fmt.Println("新增工单 读取resp.body失败 err:", e)
		return
	}
	json.Unmarshal(body, &inter)
	fmt.Println(inter)
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

func csRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "r-2zeded68f61b3b04pd.redis.rds.aliyuncs.com:6379",
		Password: "YtYnpW9dbF1Y0i3j", // no password set
		DB:       0,                  // use default DB
	})

	// 7e55fa97e5ea48ebb2fb8c4b17eab867 老兵
	result := client.HSet("privacy1_server", "7e55fa97e5ea48ebb2fb8c4b17eab867", 1)
	fmt.Print(result.String())
}
