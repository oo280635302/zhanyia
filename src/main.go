package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tealeg/xlsx"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io/ioutil"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"zhanyia/src/common"
	"zhanyia/src/must"
	"zhanyia/src/program"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 创建must组件实例
	must.Init()
	mustComponent()
	fmt.Println("run start")

	program.Ingress()

	return
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

func curTime(unix int64) int64 {
	return ((unix-16*3600)/86400 + 1) * 86400
}

func httpReq() {

	request, err := http.NewRequest("GET", "https://rvakva.xiaokayun.cn/v1/allChannels", nil)
	if err != nil {
		fmt.Println(" http newRequest has err:", err)
		return
	}
	request.URL.RawQuery = "companyId=417&companyName=%E5%9B%9B%E5%B7%9D%E5%B0%8F%E5%92%96%E7%A7%91%E6%8A%80%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8&userId=1682&userName=shine&timestamp=1625628791"

	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	request.Header.Set("token", "eyJhbGciOiJIUzI1NiIsInJvbGVJZCI6IiIsInR5cCI6IkpXVCJ9.eyJhcHBLZXkiOiI0ODg0NDE5OTg5NTI0MzVkYTg5NTI4NjYzMmU4MmY0MCIsImV4cCI6MTYyNTg4Nzk4MywiaWF0IjoxNjI1NjI4NzgzLCJ1c2VySWQiOjE2ODJ9.Ay6cVhOtb0RMPQgT2ZPxZgSFaFlPMqQT3XUgPEKLv7w")

	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println(" http 请求失败 err:", err, resp)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(" io real 失败", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	res := make(map[string]interface{})
	if err = json.Unmarshal(body, &res); err != nil {
		fmt.Println(" 解析body 失败", err)
		return
	}
	fmt.Println(res)
	return
}

func getPublicToken() {
	resp, err := http.PostForm("http://117.172.236.74:30011"+"/v1/platform/login", url.Values{
		"appKey":      {"488441998952435da895286632e82f40"},
		"platformKey": {"73f1d74553e6c802070142e254c8f277"},
	})
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("1:", err, resp)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取公服Token GetToken 读取body err:", err)
		return
	}
	defer resp.Body.Close()
	m := make(map[string]interface{})
	json.Unmarshal(body, &m)
	fmt.Println(m)
}

func csGorm() {
	db, err := gorm.Open("mysql", "root:123@tcp(localhost:3306)/?charset=utf8mb4")
	if err != nil {
		fmt.Println("1", err)
		return
	}
	defer db.Close()
	db.LogMode(true)

}

func csMysql() {
	const (
		default_employ_fields = "`id`,`user_name`,`password`,`name`,`real_name`,`sex`,`id_card`,`phone`,`portrait_path`,`idcard_path`,`home_address`,`emergency`,`emergency_phone`," +
			"`licensing_time`,`drive_license_path`,`drive_license_type`,`introducer`,`remark`,`is_freezed`,`status`,`order_status`,`company_id`,`app_version`,`balance`,`pre_money`,`driver_type`," +
			"`job_time`,`device_no`,`device_type`,`full_body_path`,`grade`,`bank_name`,`bank_card_no`,`cash_person_name`,`idcard_back_path`,`created`,`updated`,`is_deleted`,`app_key`,`version`," +
			"`audit_type`,`score`,`device_info`,`device_version`,`device_net`,`device_net_type`,`reject`,`privilege_start_date`,`privilege_end_date`,`switch_privilege`," +
			"`start_freezed_time`,`end_freezed_time`,`qr_code`,`freeze_remark`,`privilege_nums`,`privilege_type`"
	)

	db, err := sql.Open("mysql", "xiaoka:Xiaoka520@tcp(rm-2ze0624x75gk25025fo.mysql.rds.aliyuncs.com:3306)/?charset=utf8&parseTime=true&loc=Local")
	//db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(600 * time.Second)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)
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

func csRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "r-2zeded68f61b3b04pd.redis.rds.aliyuncs.com:6379",
		Password: "YtYnpW9dbF1Y0i3j", // no password set
		DB:       0,                  // use default DB
	})

	result, err := client.HMGet("58ad60944a3745b6aea63212b531f0b3_info").Result()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	for _, v := range result {
		fmt.Println(v)
	}

	//r := redis.NewClient(&redis.Options{
	//	Addr: "localhost:6379",
	//	DB:   0, // use default DB
	//})
	//
	//t, _ := r.PTTL("123").Result()
	//fmt.Println(t == time.Millisecond*-2)
}

func csMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("conn : ", err)
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("PING ERROR:", err)
		return
	}
	cur, err := client.Database("db1").Collection("employ").Find(ctx, bson.D{
		{
			"$or", []interface{}{
				bson.D{
					{
						"id", bson.D{
							{"$lte", 1},
						},
					},
				},
				bson.D{
					{
						"id", bson.D{
							{"$gte", 3},
						},
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for cur.Next(ctx) {
		a := make(map[string]interface{}, 0)
		err = cur.Decode(&a)
		fmt.Println(err, a)
	}

}

func csXlsx() {
	file, err := xlsx.OpenFile("E://downfile/EmployTemplate_1612401111603.xlsx")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	s := file.Sheet["sheet1"]
	//fmt.Println(len(s.Rows))
	fmt.Println(s)
}
