package program

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

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

// 复数乘法
func complexNumberMultiply(num1 string, num2 string) string {
	num1Arr := strings.Split(num1, "+")
	num2Arr := strings.Split(num2, "+")

	real1, _ := strconv.Atoi(num1Arr[0])
	image1, _ := strconv.Atoi(num1Arr[1][:len(num1Arr[1])-1])
	real2, _ := strconv.Atoi(num2Arr[0])
	image2, _ := strconv.Atoi(num2Arr[1][:len(num2Arr[1])-1])

	return fmt.Sprintf("%d+%di", real1*real2-image1*image2, real1*image2+real2*image1)
}

// 各位相加 反复将各个位上的数字相加，直到结果为一位数
func addDigits(num int) int {
	/*
		求根数：除0以外其余都是 1-9的
			9是取模 +1
			非9取模
	*/

	return (num-1)%9 + 1
}
