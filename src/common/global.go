package common

// 全局实例
var AllGlobal map[string]interface{}

var Log LogInterface

// 初始化 启动全局变量
func init() {
	AllGlobal = make(map[string]interface{})
}
