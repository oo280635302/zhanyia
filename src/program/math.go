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

// 七进制数
// 思路： 从个数开始找，取7的模得出来的就是当前位数的值
func convertToBase7(num int) string {
	if num == 0 {
		return ""
	}
	ans := make([]byte, 0)
	re := num < 0
	if re {
		num = -num
	}

	for num > 0 {
		ans = append(ans, byte('0'+num%7))
		num /= 7
	}

	reverseString(ans)

	if re {
		return "-" + string(ans)
	}

	return string(ans)
}

// 阶乘后的零 的个数
func trailingZeroes(n int) (ans int) {
	/*
		就是计算因数为5的个数
		eg: 130: 5、10、15...130	26个
				 25、50...125		5 个
				 125				1 个
			总计32个
	*/
	for n > 0 {
		n /= 5
		ans += n
	}
	return
}

// 找出缺失的观测数据
func missingRolls(rolls []int, mean int, n int) []int {
	m := len(rolls)
	all := (m + n) * mean
	lave := all
	for _, v := range rolls {
		lave -= v
	}

	// 判断合理性
	if lave < n || lave > n*6 {
		return nil
	}

	// 平均、余数
	ave, l := lave/6, lave%6

	ans := make([]int, n)
	for i := range ans {
		ans[i] = ave
		if i < l {
			ans[i]++
		}
	}

	return ans
}

// 自除数
func selfDividingNumbers(left int, right int) (ans []int) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22, 24, 33, 36, 44, 48, 55, 66, 77, 88, 99, 111, 112, 115, 122, 124, 126, 128, 132, 135, 144, 155, 162, 168, 175, 184, 212, 216, 222, 224, 244, 248, 264, 288, 312, 315, 324, 333, 336, 366, 384, 396, 412, 424, 432, 444, 448, 488, 515, 555, 612, 624, 636, 648, 666, 672, 728, 735, 777, 784, 816, 824, 848, 864, 888, 936, 999, 1111, 1112, 1113, 1115, 1116, 1122, 1124, 1128, 1131, 1144, 1155, 1164, 1176, 1184, 1197, 1212, 1222, 1224, 1236, 1244, 1248, 1266, 1288, 1296, 1311, 1326, 1332, 1335, 1344, 1362, 1368, 1395, 1412, 1416, 1424, 1444, 1448, 1464, 1488, 1515, 1555, 1575, 1626, 1632, 1644, 1662, 1692, 1715, 1722, 1764, 1771, 1824, 1848, 1888, 1926, 1935, 1944, 1962, 2112, 2122, 2124, 2128, 2136, 2144, 2166, 2184, 2196, 2212, 2222, 2224, 2226, 2232, 2244, 2248, 2262, 2288, 2316, 2322, 2328, 2364, 2412, 2424, 2436, 2444, 2448, 2488, 2616, 2622, 2664, 2688, 2744, 2772, 2824, 2832, 2848, 2888, 2916, 3111, 3126, 3132, 3135, 3144, 3162, 3168, 3171, 3195, 3216, 3222, 3264, 3276, 3288, 3312, 3315, 3324, 3333, 3336, 3339, 3366, 3384, 3393, 3432, 3444, 3492, 3555, 3612, 3624, 3636, 3648, 3666, 3717, 3816, 3864, 3888, 3915, 3924, 3933, 3996, 4112, 4116, 4124, 4128, 4144, 4164, 4172, 4184, 4212, 4224, 4236, 4244, 4248, 4288, 4332, 4344, 4368, 4392, 4412, 4416, 4424, 4444, 4448, 4464, 4488, 4632, 4644, 4824, 4848, 4872, 4888, 4896, 4932, 4968, 5115, 5155, 5355, 5515, 5535, 5555, 5775, 6126, 6132, 6144, 6162, 6168, 6192, 6216, 6222, 6264, 6288, 6312, 6324, 6336, 6366, 6384, 6432, 6444, 6612, 6624, 6636, 6648, 6666, 6696, 6762, 6816, 6864, 6888, 6912, 6966, 6984, 7112, 7119, 7175, 7224, 7266, 7371, 7448, 7476, 7644, 7728, 7777, 7784, 8112, 8128, 8136, 8144, 8184, 8224, 8232, 8248, 8288, 8328, 8424, 8448, 8488, 8496, 8616, 8664, 8688, 8736, 8824, 8832, 8848, 8888, 8928, 9126, 9135, 9144, 9162, 9216, 9288, 9315, 9324, 9333, 9396, 9432, 9612, 9648, 9666, 9864, 9936, 9999}
	m := make(map[int]bool, len(arr))
	for _, v := range arr {
		m[v] = true
	}
	for i := left; i <= right; i++ {
		if m[i] {
			ans = append(ans, i)
		}
	}
	return
}

// 有效的正方形
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	points := [][]int{p1, p2, p3, p4}
	counter := map[int]int{}
	distance := func(a []int, b []int) int {
		return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
	}
	for i, n := 0, len(points); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			counter[distance(points[i], points[j])] += 1
		}
	}
	if len(counter) != 2 {
		return false
	}
	x, y := int(1e9), 0
	for k, _ := range counter {
		if x == 1e9 {
			x = k
		} else {
			y = k
		}
	}
	if x > y {
		x, y = y, x
	}
	return y == 2*x && counter[x] == 4 && counter[y] == 2
}

// 检查「好数组」 - 子集能组合成公约数1的数组
func isGoodArray(nums []int) bool {
	g := 0
	for _, val := range nums {
		g = gcd(g, val)
		if g == 1 {
			return true
		}
	}
	return false
}

// 获取最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a // 将a削弱到0，看其余数b
	}
	return b
}

// 负二进制数相加
func addNegabinary(arr1 []int, arr2 []int) (ans []int) {
	i := len(arr1) - 1
	j := len(arr2) - 1
	carry := 0
	for i >= 0 || j >= 0 || carry != 0 {
		x := carry
		if i >= 0 {
			x += arr1[i]
		}
		if j >= 0 {
			x += arr2[j]
		}

		if x >= 2 {
			ans = append(ans, x-2)
			carry = -1
		} else if x >= 0 {
			ans = append(ans, x)
			carry = 0
		} else {
			ans = append(ans, 1)
			carry = 1
		}
		i--
		j--
	}
	for len(ans) > 1 && ans[len(ans)-1] == 0 {
		ans = ans[:len(ans)-1]
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return ans
}

// 可被三整除的偶数的平均值
func averageValue(nums []int) int {
	res := 0
	cnt := 0
	for _, v := range nums {
		if v%6 == 0 {
			res += v
			cnt++
		}
	}
	if cnt == 0 {
		return res
	}
	return res / cnt
}
