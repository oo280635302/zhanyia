package common

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"
)

// 取int64绝对值
func AbsInt64(n int64) int64 {
	y := n >> 63       // y ← x >> 63
	return (n ^ y) - y // (x ⨁ y) - y
}

// 检测文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// TimerCallBack 定时器回掉函数
type TimerCallBack func() bool

// 定时器
// 定时 运行callBack -回调返回true就循环执行 -false就关闭
// 可用chan bool 关闭
func GoTimer(d time.Duration, callBack TimerCallBack) (*time.Timer, chan bool) {
	t := time.NewTimer(d)
	stop := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-t.C:
				next := callBack()
				if !next {
					t.Stop()
					return
				}
				t.Reset(d)
			case <-stop:
				t.Stop()
				return
			}
		}
	}()
	return t, stop
}

// 关闭定时器
func CloseTimer(tm *time.Timer, ch chan bool) {
	ch <- true
	for !tm.Stop() {
		close(ch)
	}
}

// 错误日志输出
func LogErr(a ...interface{}) {
	Log.LogErr(a...)
}

// 调试日志输出
func LogDeBug(a ...interface{}) {
	Log.LogDebug(a...)
}

// 判断字符串是不是浮点数
func JudgeStringIsFloat(val string) bool {
	pattern := `^(\d+)\.\d+$`
	result, _ := regexp.MatchString(pattern, val)
	return result
}

// 判断字符床是不是纯数字
func JudgeStringIsInt(val string) bool {
	pattern := `^(\d+)$`
	result, _ := regexp.MatchString(pattern, val)
	return result
}

// 反转字符串
func ReverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

// 将任意数字 转成 整数及其小数位数
func TakeEveryNumToInt(num string) (string, int) {
	nums := strings.Split(num, ".")
	if len(nums) == 1 {
		return nums[0], 0
	} else if len(nums) == 2 {
		return nums[0] + nums[1], len(nums[1])
	}
	return "", -1
}

// 去除字符串数字多余的0
func RmStringNumRemainZero(a string) string {
	if JudgeStringIsFloat(a) {
		for k, v := range a {
			if v == 46 {
				break
			}
			if a[k] == 48 && a[k+1] != 46 {
				a = a[1:]
			}
		}
		for i := len(a) - 1; i > 0; i-- {
			if a[i] != 48 {
				break
			}
			a = a[:i]
		}
		return a
	}
	for _, v := range a {
		if v != 48 {
			break
		}
		a = a[1:]
	}
	return a
}

// 起止时间戳
type StartEndTimeStamp struct {
	start int64
	end   int64
}

type Unit int32

const (
	Hour Unit = 1
	Day  Unit = 2
)

// 获取X年X月X日 -> X年X月X日 的所有时间戳  天/小时
func getAllTimeStamp(sYear, sMonth, sDay, eYear, eMonth, eDay int, uint Unit) []*StartEndTimeStamp {
	start := time.Date(sYear, time.Month(sMonth), sDay, 0, 0, 0, 0, time.Local).Unix()
	end := time.Date(eYear, time.Month(eMonth), eDay, 0, 0, 0, 0, time.Local).Unix()
	var addT int64
	if uint == 1 {
		addT = 60 * 60
	} else if uint == 2 {
		addT = 60 * 60 * 24
	}
	arr := make([]*StartEndTimeStamp, 0)

	for {
		//到最后结束
		if start == end {
			break
		}
		oneData := &StartEndTimeStamp{}
		oneData.start = start
		start += addT
		oneData.end = start
		arr = append(arr, oneData)
	}
	return arr
}

func UnmarshalPb2Url(message proto.Message) {
	vm := reflect.ValueOf(message).Elem()
	str := ""
	for i := 0; i < vm.NumField(); i++ {
		refField := vm.Type().Field(i)

		jsonTagRev := refField.Tag.Get("json")
		jsonTagArr := strings.Split(jsonTagRev, ",")
		jsonTag := jsonTagArr[0]
		if jsonTag == "-" {
			continue
		}
		kd := vm.FieldByName(refField.Name).Kind()

		switch kd {
		case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			str += fmt.Sprintf("%v=%s&", jsonTag, strconv.Itoa(int(kd)))
		case reflect.String:
			str += fmt.Sprintf("%v=%s&", jsonTag, string(kd))
		case reflect.Bool:
			str += fmt.Sprintf("%v=%s&", jsonTag, kd)
		}
	}
	fmt.Println(str)
}

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu    sync.Mutex
	calls map[string]*call
}

var g Group

// 同一执行时间内的同一操作 只执行一次且返回相同数据
func Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.calls == nil {
		g.calls = make(map[string]*call)
	}

	if c, ok := g.calls[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := &call{}
	c.wg.Add(1)
	g.calls[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()
	g.mu.Lock()
	delete(g.calls, key)
	g.mu.Unlock()
	return c.val, c.err
}

// 不应gc的线程安全指针类型
var builderPool = sync.Pool{
	New: func() interface{} {
		return &strings.Builder{}
	},
}

// 获取可以自动扩容字符串builder 减少时空浪费
func GetBuilder() *strings.Builder {
	return builderPool.Get().(*strings.Builder)
}

// 清理builder
func DeleteBuilder(b *strings.Builder) {
	b.Reset()
	builderPool.Put(b)
}

// 0拷贝 string to bytes
func StringToBytes(s string) []byte {
	l := len(s)
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
		Len:  l,
		Cap:  l,
	}))
}
