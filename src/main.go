package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-redis/redis/v8"
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
	"zhanyia/src/common"
	"zhanyia/src/must"
	"zhanyia/src/program"
)

func main() {
	// 创建must组件实例
	fmt.Println("run starting")
	program.Ingress()

	a := "/home/ubuntu/titan/bak handleAutoQuest FinishCreateKill"
	fmt.Println(strings.Trim(a, ","))

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
	time.Sleep(time.Second)
}

func getGitCommitHash() (string, error) {
	// 执行git命令获取最后一次提交的哈希
	cmd := exec.Command("git", "rev-parse", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// 求解直角三角形的边长
func solveTriangle(side, angle float64) (float64, float64) {
	// 将角度转换为弧度
	radian := angle * math.Pi / 180

	// 计算对边和邻边的长度
	opposite := side * math.Sin(radian)
	adjacent := side * math.Cos(radian)

	return opposite, adjacent
}

// 计算两个坐标点之间的夹角度数（以度数为单位）
func calculateAngle(srcX, srcY, targetX, targetY float64) float64 {
	// 计算向量之间的夹角
	dx := targetX - srcX
	dy := targetY - srcY
	radian := math.Atan2(dy, dx)
	angle := radian * 180 / math.Pi

	// 将角度转换为0到360度的范围
	if angle < 0 {
		angle += 360
	}

	return angle
}

type League struct {
	Area   int
	Locale int
}

func LeagueSort(arr []League, area, locale int) []League {
	sort.Slice(arr, func(i, j int) bool {
		// 比较优先级
		if arr[i].Area == area && arr[i].Locale == locale {
			return true
		}
		if arr[j].Area == area && arr[j].Locale == locale {
			return false
		}
		if arr[i].Area == area {
			return true
		}
		if arr[j].Area == area {
			return false
		}
		if arr[i].Locale == locale {
			return true
		}
		return false
	})

	return arr
}

type IPPosition struct {
	Status   string `json:"status"`   // 返回状态
	Info     string `json:"info"`     // 错误原因
	InfoCode string `json:"infocode"` // 错误码
	Province string `json:"province"` // 省
	City     string `json:"city"`     // 城市
	AdCode   string `json:"adcode"`   // adcode编码
	Lock     sync.RWMutex
}

func GetPositionByIP(ip string) *IPPosition {
	urlPath := "https://restapi.amap.com/v5/ip?"
	param := url.Values{}
	param.Set("key", "323abb4090b1ac829cef26be6c768e14")
	param.Set("ip", ip)
	param.Set("type", "4")

	req, err := http.NewRequest("GET", urlPath+param.Encode(), strings.NewReader(""))
	if err != nil {
		return nil
	}
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil || len(body) == 0 {
		return nil
	}

	ans := &IPPosition{}
	json.Unmarshal(body, ans)
	return ans
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
	//db, err := sql.Open("mysql", "xiaoka:Xiaoka520@tcp(rm-2ze0624x75gk25025fo.mysql.rds.aliyuncs.com:3306)/?charset=utf8&parseTime=true&loc=Local")
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(600 * time.Second)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)

	r, err := db.Exec("update cs.cs set rid = 10085 where rid = 10086")
	if err != nil {
		fmt.Println("1111", err)
		return
	}
	cnt, err := r.RowsAffected()
	if err != nil || cnt == 0 {
		fmt.Println("2222", err, cnt)
		return
	}
	fmt.Println(33333)
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

const PUSH_PRIVATE_MESSAGE_SCRIPT_STRING = `
	local keyCmp = function(el1,el2) 
		return tonumber(el1) < tonumber(el2)
	end

	local hashKey = KEYS[1]
	local memRecordKey = KEYS[2]
	local pid = ARGV[1]
	local msgId = ARGV[2]
	local msg = ARGV[3]
	local now = tonumber(ARGV[4])
	local max = tonumber(ARGV[5])

	redis.pcall("hset",hashKey,msgId,msg)

	local msgCnt = redis.pcall("hlen",hashKey)
	if msgCnt > max then
		local tbAllMsgId = redis.pcall("hkeys",hashKey)
		table.sort(tbAllMsgId,keyCmp)
		local delMsgId = {}
		for i = 1, #tbAllMsgId - max do
			delMsgId[i] = tbAllMsgId[i]
		end
		redis.pcall("hdel",hashKey,unpack(delMsgId))
		msgCnt = redis.pcall("hlen",hashKey)
	end
	
	redis.pcall("hset",memRecordKey,pid,string.format("%d:%d",msgCnt,now))

	return msgId
`

func csRedis() {
	ctx := context.TODO()
	client := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	for i := 0; i < 1000; i++ {

		client.ZRevRangeWithScores(ctx, "ltt", 0, 99).Result()

		continue
		score := float64(i) + float64(time.Now().Unix())*0.00000000001
		client.ZAdd(ctx, "ltt", &redis.Z{Score: score, Member: i})
	}

	return
	script := `
		local k1 = KEYS[1]
		local v1 = ARGV[1]
		local v2 = ARGV[2]
		
		redis.pcall("zadd",k1,v1,v2)
		
		local cnt = redis.pcall("zcard",k1)

		local mem = 1
		if cnt > tonumber(5) then
			mem = redis.pcall("zrange",k1,0,0)
			redis.pcall("zrem",k1,mem[1])
		end
		return mem
	`

	redisScript := redis.NewScript(script)

	for i := 0; i < 100; i++ {
		key := i
		value := float64(i) + float64(math.MaxInt32-time.Now().Unix())*0.0000000001
		val, err := redisScript.Run(ctx, client, []string{"cs_rank"}, value, key).Result()
		fmt.Println(val, err)
		time.Sleep(time.Second / 10)
	}
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

// 反射体结构注入
func reflectStruct() {
	type Cs1 struct {
		Val1 int8
	}
	type Cs struct {
		Val1 int8
		Cs   *Cs1
	}

	c := &Cs{}
	v := reflect.ValueOf(c)

	ve := v.Elem()
	for i := 0; i < ve.NumField(); i++ {

		switch ve.Field(i).Kind() {
		case reflect.Ptr:
			ve.Field(i).Set(reflect.New(ve.Field(i).Type().Elem()))
		case reflect.Int8:
			ve.Field(i).SetInt(11)
		}
	}

	fmt.Println(c)
}

func read(c reflect.Value) {
	switch c.Kind() {
	case reflect.Int8:
		c.SetInt(123)
	case reflect.Ptr:
		read(c.Elem())
	case reflect.Struct:
		for i := 0; i < c.NumField(); i++ {
			fmt.Println("11", c.Field(i).Kind())
			switch c.Field(i).Kind() {
			case reflect.Ptr:
				reflect.New(c.Field(i).Elem().Type())
				//c.Field(i).Set())
			default:
				read(c.Field(i))
			}
		}
	}
}
