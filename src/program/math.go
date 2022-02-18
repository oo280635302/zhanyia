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

// 银行中的激光束数量
func numberOfBeams(bank []string) int {
	banks := make([]int, 0) // 统计每行的设备数，0设备跳过
	for _, val := range bank {
		cnt := 0
		for _, v := range val {
			if v == '1' {
				cnt++
			}
		}
		if cnt != 0 {
			banks = append(banks, cnt)
		}
	}

	if len(banks) <= 1 { // 没有多排设备无法组成激光 直接返回
		return 0
	}

	ans := 0
	cur := banks[0]
	for _, v := range banks[1:] { // 相邻两排相乘追加
		ans += cur * v
		cur = v
	}

	return ans
}
