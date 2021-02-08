package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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
	"net/url"
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

	//must.GinListener(must.NewLimitTicker(60*time.Second, 10))
	//csXlsx()
	csGorm()
	//httpReq()
	//csMysql()
	//csMongo()

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

func CC() *cs {
	return nil
}

func RAND() int {
	Rank := make(map[int]int, 3395)
	for i := 1; i <= 3395; i++ {
		Rank[i] = i
	}

	i := 1
	for _, v := range Rank {
		if v == 1244 {
			return i
		}
		i++
	}
	return -1
}

type PPhoneBindReq struct {
	AreaCode    string `protobuf:"bytes,1,opt,name=areaCode,proto3" json:"areaCode,omitempty"`
	AppKey      string `protobuf:"bytes,2,opt,name=appKey,proto3" json:"appKey,omitempty"`
	CallerNum   string `protobuf:"bytes,3,opt,name=callerNum,proto3" json:"callerNum,omitempty"`
	CalleeNum   string `protobuf:"bytes,4,opt,name=calleeNum,proto3" json:"calleeNum,omitempty"`
	Duration    int64  `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	MaxDuration int64  `protobuf:"varint,6,opt,name=maxDuration,proto3" json:"maxDuration,omitempty"`
	RecordFlag  bool   `protobuf:"varint,7,opt,name=recordFlag,proto3" json:"recordFlag,omitempty"`
	NotifyUrl   string `protobuf:"bytes,8,opt,name=notifyUrl,proto3" json:"notifyUrl,omitempty"`
}

func httpReq() {
	token := "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZW1wIiwicGF5bG9hZCI6IjQ4ODQ0MTk5ODk1MjQzNWRhODk1Mjg2NjMyZTgyZjQwIiwiaXNzIjoi5Zub5bed5bCP5ZKW56eR5oqA5pyJ6ZmQ5YWs5Y-4IiwiaWF0IjoxNjA4ODAwODY4LCJleHAiOjE2MDk0MDU2Njh9._P0RBvYGsa4F6FCoVpOrKpIz06-QsRxpbJHH9A4gWyg"

	req := &PPhoneBindReq{}
	req.AppKey = "488441998952435da895286632e82f40"
	req.AreaCode = "028"
	req.CalleeNum = "13982552218"
	req.CallerNum = "13980494026"
	req.RecordFlag = true
	req.Duration = 6 * 60 * 60
	req.MaxDuration = 30
	req.NotifyUrl = "https://api.xiaokayun.cn/api/v1/privacy/phone/notify"
	param, _ := json.Marshal(req)

	request, err := http.NewRequest("POST", "http://117.172.236.74:30011"+"/v1/privacy/open/axb/binding", strings.NewReader(string(param)))
	if err != nil {
		fmt.Println("绑定隐私号 http newRequest has err:", err)
		return
	}

	request.Header.Set("Content-Type", "application/json;charset=utf-8")
	request.Header.Set("Authorization", token)

	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(request)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("绑定隐私号 http 请求失败 err:", err, resp)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("绑定隐私号 io real 失败", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	res := make(map[string]interface{})
	if err = json.Unmarshal(body, &res); err != nil {
		fmt.Println("绑定隐私号 解析body 失败", err)
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

type cs struct {
	Id    int64  `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func csGorm() {
	db, err := gorm.Open("mysql", "root:123@tcp(localhost:3306)/?charset=utf8mb4")
	if err != nil {
		fmt.Println("1", err)
		return
	}
	defer db.Close()
	db.LogMode(true)

	cc := make([]*cs, 0)
	cc = append(cc, &cs{
		Key:   "18",
		Value: "18",
	})
	cc = append(cc, &cs{
		Key:   "18",
		Value: "18",
	})

	for _, v := range cc {
		bd := db.Table("cs.cs").Create(v)
		if bd.Error != nil {
			fmt.Println(bd.Error)
			return
		}
		fmt.Println(bd.Value.(*cs))
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
	//stm, err := db.Prepare("insert into `cs`.`1E` (`key1`,`key2`,`value1`,`value2`) value (?,?,?,?)")
	//if err != nil {
	//	fmt.Println("insert has err:", err)
	//	return
	//}
	//for i := 23004; i <= 100000; i++ {
	//	r, err := stm.Exec(i, i, i, i)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(r.LastInsertId())
	//}
	//a := []int64{1,2,3}
	//b, _ := json.Marshal(a)
	//r,err:= db.Exec("insert into `cs`.`cs` (`key`,`arr`) value (?,?)",1,b)
	//if err != nil {
	//	fmt.Println("ins err:",err)
	//	return
	//}
	//fmt.Println(r.LastInsertId())

	rows, err := db.Query("select `key` from `cs`.`cs` where  `id` in (" + "1,2,3,4" + ")")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		key := ""
		rows.Scan(&key)
		fmt.Println(key)
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
