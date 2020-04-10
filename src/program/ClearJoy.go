package program

import (
	"fmt"
	"math/rand"
)

// 制作一张 横*竖 的二维地图
func MakeMap(horizontal int, vertical int) [][]int {
	a := make([][]int, 0)
	for i := 0; i < vertical; i++ {
		b := make([]int, 0)
		for i := 0; i < horizontal; i++ {
			b = append(b, 0)
		}
		a = append(a, b)
	}
	return a
}

//  二维地图 中的0进行 填充 1-5 的随机数
func FullMap(writeMap [][]int) {
	for i := 0; i < len(writeMap); i++ {
		for j := 0; j < len(writeMap[i]); j++ {
			if writeMap[i][j] == 0 {
				writeMap[i][j] = rand.Intn(5) + 1
			}
		}
	}
}

// console输出成 二维图 的样子
func PrintDoubleMap(writeMap [][]int) {
	for i := 0; i < len(writeMap); i++ {
		fmt.Println(writeMap[i])
	}
	fmt.Println()
}

// 消掉的图标下坠
// 		eg：
//		1 0 0 0 0		0 0 0 0 0
//		0 2 1 4 6 ->    1 2 0 4 6
//		2 3 0 7 5       2 3 1 7 5
func IconFall(writeMap [][]int) {
	// 一共下沉 高度-1 次
	// 例如 3层 要进行2次
	// 3*5 = 2*2*5 = 20次

	for x := 0; x < len(writeMap)-1; x++ {
		// 从低到高 - 跳过第1层
		for i := len(writeMap) - 1; i > 0; i-- {
			for j := 0; j < len(writeMap[i]); j++ {
				// 当图标是0 将上一层的 对应j的图标进行交换
				if writeMap[i][j] == 0 {
					writeMap[i][j], writeMap[i-1][j] = writeMap[i-1][j], writeMap[i][j]
				}
			}
		}
	}

}
