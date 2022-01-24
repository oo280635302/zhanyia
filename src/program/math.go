package program

import "math"

// Pow(x, n) 计算 x 的 n 次幂函数
// 思路：API 造什么轮子？
func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}

// 计算力扣银行的钱
// 思路：底为28元，每周比上周多7元 + 每天比昨天多1元 + 周 * 天
func totalMoney(n int) int {
	x := n / 7
	y := n % 7

	res := x*21 + (1+x)*x/2*7 + (y * x) + ((1 + y) * y / 2)
	return res
}

// 比赛中的配对次数
// 思路：每两队晋级1对，奇数自动晋级一对，比赛场数n-1
func numberOfMatches(n int) int {
	return n - 1
}
