package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"zhanyia/src/common"
	"zhanyia/src/must"
	"zhanyia/src/program"
	pb "zhanyia/src/proto"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(program.MergeTwoLists(&program.ListNode{Val: 10}, &program.ListNode{Val: 11}))

	// 创建must组件实例
	must.Init()
	mustComponent()

	//
	a := &pMergedAB{}
	var length *LenMerged
	MergedStructAB1([]*ProA{{'a', 'a', 1}, {'c', 'a', 3}}, 1, []*ProB{{'a', 'a', 1}}, 1, a, length)

	fmt.Println("run start")
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

// 精灵云 笔试第一题 合并结构体
type ProA struct {
	key1 byte
	key2 byte
	val1 int
}

type ProB struct {
	key1 byte
	key2 byte
	val2 int
}

type MergedAB struct {
	key1 byte
	key2 byte
	val1 int
	val2 int
}

type pMergedAB struct {
	val  *MergedAB
	next *pMergedAB
}

type LenMerged int

func (p *pMergedAB) insert(pMergedAB *pMergedAB) {

	temp := p
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	if temp.val == nil {
		temp.val = pMergedAB.val
	} else {
		temp.next = pMergedAB
	}
}

// method 1
func MergedStructAB1(arrA []*ProA, lenA int, arrB []*ProB, lenB int, mergedAB *pMergedAB, lengthMergedAB *LenMerged) {

	for _, v1 := range arrA {
		for _, v2 := range arrB {
			if v1.key1 == v2.key1 && v1.key2 == v2.key2 {
				mergedAB.insert(&pMergedAB{val: &MergedAB{key1: v1.key1, key2: v2.key2, val1: v1.val1, val2: v2.val2}})
			}
		}
	}
	fmt.Println(mergedAB)
}

// method 2
func MergedStructAB2(arrA []*ProA, lenA int, arrB []*ProB, lenB int, mergedAB *pMergedAB, lengthMergedAB *LenMerged) {

	for _, v1 := range arrA {
		for _, v2 := range arrB {
			if v1.key1 == v2.key1 && v1.key2 == v2.key2 {
				mergedAB.insert(&pMergedAB{val: &MergedAB{key1: v1.key1, key2: v2.key2, val1: v1.val1, val2: v2.val2}})
			}
		}
	}
	fmt.Println(mergedAB)
}
