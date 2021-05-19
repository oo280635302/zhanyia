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
	"math/big"
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
	"zhanyia/src/program"
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
	//csGorm()
	//httpReq()
	//csMysql()
	//csMongo()
	//fmt.Println("123")
	//cs()
	program.Ingress()

	big.NewInt(1)

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

	req := map[string]interface{}{
		"app_key":    "36db954333b14bf9ad14f7f5bfa24fde",
		"company_id": 1,
		"channel":    1,
		"timestamp":  1618983874,
	}

	param, _ := json.Marshal(req)

	request, err := http.NewRequest("PUT", "https://api.xiaokacloud.com/api/v1/small/read_advertise", strings.NewReader(string(param)))
	if err != nil {
		fmt.Println("绑定隐私号 http newRequest has err:", err)
		return
	}

	request.Header.Set("Content-Type", "application/json;charset=utf-8")

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

	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(600 * time.Second)
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(20)

	//arr := []string{"xxx", "yyy"}

	stm, err := db.Prepare("insert into `cs`.`cs` (`key`,`value`) values (?,?);")
	if err != nil {
		fmt.Println("err1", err)
		return
	}
	for i := 19; i < 29; i++ {
		r, err := stm.Exec(i, i)
		if err != nil {
			fmt.Println("err2", err)
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

//告ㄴ可以透过公会聊天视窗进行公告编辑、修改13、改善英雄/装备平衡、Bug修复■改善英雄／
//装备平衡-增加部分英雄普通攻击值(图鉴5星，1等为基准)ㄴ米雅:攻击力168 (+4),治愈1440
//(+36)ㄴ瑞皮娜:防御力51 (+5)-兰儿特殊能力在非战斗时也会套用，每30秒启动一次。ㄴ更新
//前:更新前:生命值低于50%的受伤友军回复20%生命值。苏醒符对每位友军仅触发一次。ㄴ更新后
//:生命值低于50%的受伤友军回复20%生命值。苏醒符每30秒触发一次。-普利菲斯特殊技能效果启
//动机率从15%升至20%。-亚马洛克的地狱猎犬准备期间减少0.05秒，距离及速度增加10%亚马洛
//克专武特殊加成，以30%的机率召唤影狼，增加至40%，影狼的间隔减少20%。-香格里拉的煽风点
//火武器技能伤害从每秒200%增加至280%。■ Bug修复-维萝妮卡与诺克西亚进化能力值设定错误
//问题已修复。ㄴ三星进化至四星的加成效果，直到四星进化至五星后才套用的问题已修复。ㄴ从
//4星进化至5星后，特殊能力没有显示的问题已修复。-蒂尼亚普通攻击的音效已修复。- Mk.99使
//用普通攻击，光束麦格农时，错误判定目标位置的问题已修复。-未来公主使用连锁技能时，异
//常穿墙的问题已修复。-偶像骑士团长伊娃、米雅、Mk.99与贝丝的特殊能力于部份情况异常套用
//问题已修复。-拉娜的护腕伤害减免效果套用异常问题已修复。14、BUG修正和功能改善=-周边商
//品铸造结果动画，Mk.99设定为代表角色，异常显示问题已修复。- MK.99骑乘炸弹虫时，高度及
//方向变化异常问题已修复。-进行攻击转场时若是有其他视窗跳出，转场画面会停住的问题已修
//复。-于裂谷副本，自动战斗按键将于开始时自动启用。-训练室ㄴ放置或释出英雄时，将可查看
//相关图像。ㄴ可于训练室中查看确认英雄的状态资讯。-守护者基地自动编队ㄴ在守护者基地自
//动编队时，确认守护者点数产量的增加幅度。.ㄴ如果守护者基地产量已达优化最大值，自动编
//队按键将无法点击。-优化新增英雄经验值的介面设计。-公会-超级时装将套用至稻草人雕像上。1
//5、商店—礼包-英雄成长礼包(恶魔女王莉莉丝)※独特英雄成长礼包会于召唤对应英雄后出现。※
//独特英雄成长礼包将会于首次出现后显示7天，如果于7天内都没有购买，将不会再次显示。-该礼包只
//限4/27(二)创立帐号后的玩家(限时7天)ㄴ期间: 2021-05-18维护后第2季飞跃成长礼包#1-自选x2 /固
//定内容x1ㄴ独特英雄选择x1ㄴ专属武器选择x1-购买获得ㄴ Lv. 70觉醒副本箱x 10第2季飞跃成长礼包#
//2-自选x1 /固定内容x2ㄴ独特英雄选择x1或是专属武器选择x1-购买获得ㄴ召唤控制器x 100ㄴ Lv. 70觉
//醒副本箱x 10-回归礼包-超过20天没有登入游戏的玩家(限时7天)ㄴ期间: 2021-05-04维护后~ 2021-06
//-20 07:59:59第2季回归守护者飞跃成长礼包-自选x2 /固定内容x2ㄴ独特英雄选择x1ㄴ专属武器选择x1-
//购买获得ㄴ召唤控制器x 60ㄴ Lv. 70觉醒副本箱x10-一般礼包ㄴ期间: 2021-05-18维护后~ 2021-06-0
//1 07:59:59第2季限时特价礼包ㄴ付费宝石8100ㄴ金币800000ㄴ经验值1200000第2季高级召唤礼包ㄴ订阅
//礼包(10天)-购买获得召唤控制器x100史诗突破上限锤x1ㄴ订阅内容1000宝石x 10天
