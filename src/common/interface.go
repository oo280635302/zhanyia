package common

// 日志
type LogInterface interface {
	// 错误输出
	LogErr(a ...interface{})
	// 调试输出
	LogDebug(a ...interface{})
}
