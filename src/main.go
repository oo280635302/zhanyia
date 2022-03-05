package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"os/signal"
	"strconv"
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

	s := time.Now().UnixNano()
	program.Ingress()
	fmt.Println("耗时：", (time.Now().UnixNano()-s)/1e6)

	fmt.Println(len("oh shit fucker fucccker shit fuckee fuc s shit shiit  shhit !"))

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

func RoundFormatFloat(f float64, scale int) float64 {
	result, _ := strconv.ParseFloat(strconv.FormatFloat(f, 'f', scale+1, 64), 64)

	pow := math.Pow(10, float64(scale))

	return math.Round(result*pow) / pow
}

func Mqtt() {
	opt := mqtt.NewClientOptions().AddBroker("tcp://v6prev-wss.rvaka.cn:1883").SetClientID("ltt_send").SetUsername("ltt")
	c := mqtt.NewClient(opt)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("connect err: ", token.Error())
		return
	}

	for {
		topic := fmt.Sprintf("ltt_go")
		err := c.Publish(topic, 0, false, []byte("{\"msg\":\"罗天文测试\"}")).Error()
		if err != nil {
			fmt.Println("pub err:", err)
			return
		}
		fmt.Println(time.Now(), "send success!", topic)
		time.Sleep(time.Second)
	}
}

func MqttSub() {
	opt := mqtt.NewClientOptions().AddBroker("tcp://v6prev-wss.rvaka.cn:1883").SetClientID("ltt_receive")
	c := mqtt.NewClient(opt)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("connect err: ", token.Wait(), token.Error())
		return
	}

	topic := fmt.Sprintf("ltt_go")
	token := c.Subscribe(topic, 0, func(client mqtt.Client, message mqtt.Message) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.999"), string(message.Payload()))
		message.Ack()
	})
	token.Wait()
	err := token.Error()
	if err != nil {
		fmt.Println("sub err:", err)
		return
	}
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

type ClearRule struct {
	Min  int32 `json:"min"`  // 小数 -999999~999999
	Max  int32 `json:"max"`  // 大数 -999999~999999
	Keep int32 `json:"keep"` // 保留 -999999~999999
}

// 单型发放规则
type PullSingle struct {
	Open bool `json:"open"` // 开关

	Unit  int32 `json:"unit"`  // 单位
	Type  int8  `json:"type"`  // 类型 0固定 1比例
	Param int32 `json:"param"` // 参数
}

// 阶梯型发放规则
type PullStep struct {
	Open bool         `json:"open"`        // 开关
	Step []PullSingle `json:"pull_single"` // 阶梯 特性：数量最大6，unit递增 打开的情况下第一个不为0
}

func csGorm() {
	db, err := gorm.Open("mysql", "v5prodmcs:Gb7YJ#FP7%W866E@79R@tcp(rm-2ze75q86i46cbp80f0o.mysql.rds.aliyuncs.com:3306)/orders?charset=utf8mb4")
	if err != nil {
		fmt.Println("1", err)
		return
	}
	defer db.Close()
	db.LogMode(true)
	rows, err := db.Table("orders").Select("id").Where("id in (?)", []int64{1920, 1921, 1922}).Rows()
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int64
		rows.Scan(&id)
		fmt.Println(id)
	}
}

func csMysql() {
	db, err := sql.Open("mysql", "xiaoka:Xiaoka520@tcp(rm-2ze0624x75gk25025fo.mysql.rds.aliyuncs.com:3306)/?charset=utf8&parseTime=true&loc=Local")
	//db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(600 * time.Second)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)

	rows, err := db.Query("select  `id` from `employ_center`.`employ_infos` where `id` in (123)")
	if err != nil {
		fmt.Println("1", err)
	}
	defer rows.Close()
	for rows.Next() {
		id := 0
		rows.Scan(&id)
		fmt.Println("对咯~", id)
	}
	fmt.Println("对咯22222~")
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

func csQiniu() {
	key := uuid.New().String()

	putPolicy := storage.PutPolicy{
		Scope:   "images",
		Expires: 3600 * 24 * 365 * 2,
	}

	mac := qbox.NewMac("5bJMhEn4DSLNAJT-JiIw9rhmk8coOxMpVwGZoCRc", "mSDqSTWRySYhEatdMuGlNGKFQLhYD4Ue97XYiSD3")
	upToken := putPolicy.UploadToken(mac)

	fmt.Println(upToken)

	cfg := storage.Config{}
	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	recorder, err := storage.NewFileRecorder(os.TempDir())
	if err != nil {
		fmt.Println(err)
		return
	}
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	err = resumeUploader.PutFile(context.Background(), &ret, upToken, key+".docx", "E:\\zhanyia\\src\\server.docx", &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}

func csRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.10.4.9:30111",
		Password: "xiaoka520", // no password set
		DB:       0,           // use default DB
	})

	//fmt.Println(s,err)
	s, _ := client.ZRevRangeWithScores("employ_rank_cnt_month_202112_3437", 0, -1).Result()
	fmt.Println(s)
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

func csHBase() {
	hClient := gohbase.NewClient("localhost:123456")
	defer hClient.Close()
	hrpc.NewGetStr(context.Background(), "gpsrange", "orderId", hrpc.Families(map[string][]string{
		"info": []string{"123", "234"},
	}))

}
