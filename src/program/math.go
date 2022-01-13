package program

import "math"

// Pow(x, n) 计算 x 的 n 次幂函数
// 思路：API 造什么轮子？
func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
