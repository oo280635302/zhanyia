package must

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"zhanyia/src/common"
)

type Log struct {
	logDir   string
	lock     sync.Mutex
	logFile  *os.File
	logPrint *log.Logger
}

// 创建Mq实例
func init() {
	common.AllGlobal["Log"] = &Log{}
	common.AllGlobal["Log"].(*Log).loadLogDir()
}

// 创建日志目录
func (l *Log) loadLogDir() {
	// 获取当前文件夹位置
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Log Getwd err :%v", err))
	}
	l.logDir = dir + "\\log"

	// 创建log目录

	err = os.MkdirAll(l.logDir, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Log Mkdir err :%v", err))
	}
	l.createLog()
}

// 记录日志
func (l *Log) createLog() {
	// 锁
	l.lock.Lock()
	defer l.lock.Unlock()

	// 日志文件名
	nowTime := time.Now().Format("20060102150405")
	logName := ""
	for {
		logName = "log/log_" + nowTime + ".log"

		// 关闭之前的文件
		if l.logFile != nil {
			l.closeLastFile(l.logFile)
		}

		// 创建日志文件 - 失败重建 - 保存进内存
		f, err := os.OpenFile(logName, os.O_CREATE|os.O_EXCL|os.O_RDWR, os.ModePerm)
		if err != nil {
			continue
		}
		l.logFile = f

		// 带时间代码行号的日志
		l.logPrint = log.New(f, "", log.LstdFlags|log.Lshortfile)

		// 输出重定向
		os.Stdout = f
		os.Stderr = f

		break
	}

	// 2个小时整点 换日志文件
	timeNow := time.Now()
	timeNext := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), timeNow.Hour()+2, 0, 0, 0, time.Local)
	common.GoTimer(time.Second*time.Duration(timeNext.Unix()-timeNow.Unix()), func() bool {
		l.createLog()
		return true
	})
}

// 关闭日志
func (l *Log) closeLastFile(f *os.File) {
	go func() {
		common.GoTimer(time.Minute, func() bool {
			err := f.Close()
			if err != nil {
				return true
			}
			return false
		})
	}()
}

// 错误日志
func (l *Log) LogErr(a ...interface{}) {
	b := append([]interface{}{" ERROR: "}, a...)
	l.logPrint.Println(b...)
}

// Debug日志
func (l *Log) LogDebug(a ...interface{}) {
	l.logPrint.Println(a...)
}
