package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unsafe"
	"zhanyia/src/common"
	"zhanyia/src/must"
	pb "zhanyia/src/proto"
)

type Stu struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Pop  []string `json:"pop"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 创建must组件实例
	must.Init()
	mustComponent()
	fmt.Println("run start")

	conn, exit := must.GetMongoDB()
	if conn == nil {
		return
	}
	defer exit()

	r, err := conn.Database("db1").Collection("employ").InsertOne(context.TODO(), map[string]interface{}{
		"id":    1,
		"name":  "罗天文",
		"phone": 13982551155,
	})
	fmt.Println(r, err)

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

func ToOffsetLimit(pg, lim int64, count int64) (offset, limit int64) {
	diff := int64(0)
	if count < pg*lim {
		diff = pg*lim - count
	}
	offset = (pg - 1) * lim
	limit = lim - diff
	if limit < 0 {
		limit = 0
	}
	return
}

func byte2string2(in [16]byte) string {
	tmp := make([]byte, 0)
	x := (*[3]uintptr)(unsafe.Pointer(&tmp))
	x[0] = uintptr(unsafe.Pointer(&in))
	x[1] = 16
	x[2] = 16
	return string(tmp)
}

type cs struct {
	Id    int64
	Key   string
	Value string
}

func csGorm() {
	db, err := gorm.Open("mysql", "root:123@tcp(localhost:3306)/?charset=utf8mb4")
	if err != nil {
		fmt.Println("1", err)
		return
	}
	defer db.Close()
	db.LogMode(true)
	arr := make([]*cs, 0)

	cc := []int{1, 2}

	bd := db.Table("cs.cs").Select("*").Where("`value` in (?)", cc).Find(&arr)
	if bd.Error != nil {
		fmt.Println(bd.Error)
		return
	}
	for _, v := range arr {
		fmt.Println(*v)
	}
}

func csMysql() {

	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(600 * time.Second)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
	stm, err := db.Prepare("insert into `cs`.`1E` (`key1`,`key2`,`value1`,`value2`) value (?,?,?,?)")
	if err != nil {
		fmt.Println("insert has err:", err)
		return
	}
	for i := 23004; i <= 100000; i++ {
		r, err := stm.Exec(i, i, i, i)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(r.LastInsertId())
	}
}

func csHttp() {
	str := fmt.Sprintf("phone=%s&message=%s&countryCode=%s&type=%s&signName=%s", "13982552218", "验证码：6674,请注意查收", "86", "2", "小咖科技")

	// 创建请求
	req, err := http.NewRequest("POST", "https://manage.rvaka.com/v1/sms/open/send", strings.NewReader(str))
	if err != nil {
		fmt.Println("根据code获取access http newRequest has err:", err)
		return
	}
	//req.URL.RawQuery = param.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "yJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZW1wIiwicGF5bG9hZCI6IjQ4ODQ0MTk5ODk1MjQzNWRhODk1Mjg2NjMyZTgyZjQwIiwiaXNzIjoi5Zub5bed5bCP5ZKW56eR5oqA5pyJ6ZmQ5YWs5Y-4IiwiaWF0IjoxNjAwOTk1NDgwLCJleHAiOjE2MDE2MDAyODB9.GSkegI02HMrZBByfjCxyIaz3LHthSvKP1EBId-Z2q_")

	// 发送请求
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("http 请求失败 err:", err, resp)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || len(body) == 0 {
		fmt.Println(" http 读取resp.body失败 err:", err)
		return
	}
	fmt.Println(string(body))
	res := make(map[string]interface{})

	json.Unmarshal(body, &res)
	fmt.Println(res)
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
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "r-2zeded68f61b3b04pd.redis.rds.aliyuncs.com:6379",
	//	Password: "YtYnpW9dbF1Y0i3j", // no password set
	//	DB:       0,                  // use default DB
	//})
	//
	//result := client.HSet("privacy1_server", "7e55fa97e5ea48ebb2fb8c4b17eab867", 1)
	//fmt.Print(result.String())

	r := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // use default DB
	})

	t, _ := r.PTTL("123").Result()
	fmt.Println(t == time.Millisecond*-2)
}
