package common

import (
	"os"
	"regexp"
	"strings"
	"time"
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
