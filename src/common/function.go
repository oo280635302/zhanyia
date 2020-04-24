package common

import (
	"os"
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

//
