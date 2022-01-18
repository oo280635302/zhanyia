package program

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

// 字符串反转------------------------------------------------------------------------------------------------------------
func reverseString(s []byte) {
	length := len(s)

	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

// 数字反转--------------------------------------------------------------------------------------------------------------
// 思路：栈
func reverseInt(x int) int {
	rev := 0
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return rev
}

// 第一个不重复字符------------------------------------------------------------------------------------------------------
// 思路：数组保存每次字母出现次数，
func firstUniqChar(s string) int {
	m := [26]int{}
	for _, v := range s {
		m[v-'a']++
	}
	for idx, v := range s {
		if m[v-'a'] == 1 {
			return idx
		}
	}
	return -1
}

// 验证回文串------------------------------------------------------------------------------------------------------------
// 思路：前后双指针
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	i, j := 0, len(s)-1

	for i < j {
		if s[i] < 48 || (s[i] > 57 && s[i] < 65) || (s[i] > 90 && s[i] < 97) || s[i] > 122 {
			i++
			continue
		}
		if s[j] < 48 || (s[j] > 57 && s[j] < 65) || (s[j] > 90 && s[j] < 97) || s[j] > 122 {
			j--
			continue
		}

		tempI := s[i]
		if tempI >= 65 && tempI <= 90 {
			tempI += 32
		}

		tempJ := s[j]
		if tempJ >= 65 && tempJ <= 90 {
			tempJ += 32
		}
		fmt.Println(string(tempI), string(tempJ))
		if tempI != tempJ {
			return false
		}
		i++
		j--
	}

	return true
}

// 字符串转换成整数------------------------------------------------------------------------------------------------------
func myAtoi(s string) int {
	res := 0
	reverse := false

	for _, v := range s {
		if v == ' ' {
			s = s[1:]
		} else {
			break
		}
	}

	for idx, v := range s {
		if v != '-' && v != '+' && (v < '0' || v > '9') {
			break
		}
		if v == '-' {
			if idx == 0 {
				reverse = true
				continue
			}
			break
		}
		if v == '+' {
			if idx == 0 {
				continue
			}
			break
		}

		res = res*10 + int(v-'0')

		if res > math.MaxInt32 && !reverse {
			res = math.MaxInt32
			break
		} else if res > math.MaxInt32+1 && reverse {
			res = math.MaxInt32 + 1
			break
		}
	}

	if reverse {
		res = -res
	}

	return res
}

// z字形变换 ------------------------------------------------------------------------------------------------------------
// 思路：按顺序找到每个字符的应该在的位置，然后赋值给返回值
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([][]byte, numRows)

	for idx, b := range s {

		x := idx % (numRows - 1)
		if idx/(numRows-1)%2 == 1 {
			x = numRows - 1 - x
		}
		rows[x] = append(rows[x], byte(b))
	}

	res := ""
	for _, v := range rows {
		res += string(v)
	}

	return res
}

// 数字转换罗马字符------------------------------------------------------------------------------------------------------
// 思路：穷举求商求余
func intToRoman(num int) string {
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	res := ""
	arr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, n := range arr {
		cnt := num / n
		num = num % n
		for i := 0; i < cnt; i++ {
			res += m[n]
		}
	}
	return res
}

// 四树之和--------------------------------------------------------------------------------------------------------------
// 思路：与三树之和相似
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)

	n := len(nums) - 1

	// 固定1
	for i := 0; i <= n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 2 > 1
		for j := i + 1; j <= n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 3、4 >= 2 && 3 < 4
			l, r := j+1, n
			for l < r {
				if l != j+1 && nums[l] == nums[l-1] {
					l++
					continue
				}
				if r != n && nums[r] == nums[r+1] {
					r--
					continue
				}
				x := nums[i] + nums[j] + nums[l] + nums[r]
				if x > target {
					r--
				} else if x < target {
					l++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
				}
			}
		}
	}
	return res
}

// 括号生成 -------------------------------------------------------------------------------------------------------------
// 思路：动规+迭代写法 思路如下
func generateParenthesis(n int) []string {
	res := make([]string, 0)

	type CurString struct {
		Str   string // 当前括号
		Score int    // 当前分数 以分数代栈 不能为负数
		Lave  int    // 当前剩余左括号
	}

	// 保存上一层的括号及参数
	last := []CurString{{Str: "(", Score: 1, Lave: n - 1}}

	for i := 2; i <= 2*n; i++ {
		// 保存当前层的括号及参数
		cur := make([]CurString, 0)
		for _, v := range last {

			// 当既 分数 > 0 && 左括号有剩余 时 下个括号可以为随意
			if v.Score > 0 && v.Lave > 0 {
				// 根据题目的输出先左后右 - 其余无意义
				cur = append(cur, CurString{Str: v.Str + "(", Score: v.Score + 1, Lave: v.Lave - 1})
				cur = append(cur, CurString{Str: v.Str + ")", Score: v.Score - 1, Lave: v.Lave})
			}

			// 当 分数 > 0 && 左括号无剩余 时 下个必定是右括号
			if v.Score > 0 && v.Lave == 0 {
				cur = append(cur, CurString{Str: v.Str + ")", Score: v.Score - 1, Lave: v.Lave})
			}

			// 当 分数 = 0 && 左括号右剩余 时 因为分数不能为负 下个必定是左括号
			if v.Score == 0 && v.Lave > 0 {
				cur = append(cur, CurString{Str: v.Str + "(", Score: v.Score + 1, Lave: v.Lave - 1})
			}
		}

		last = cur
	}

	for _, v := range last {
		res = append(res, v.Str)
	}

	return res
}

// 字符串相乘------------------------------------------------------------------------------------------------------------
// 思路：从小到大依次相乘 保存结果 然后再从结果从小到大依次进位
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x * y
		}
	}
	for i := m + n - 1; i > 0; i-- {
		ansArr[i-1] += ansArr[i] / 10
		ansArr[i] %= 10
	}
	ans := ""
	idx := 0
	if ansArr[0] == 0 {
		idx = 1
	}
	for ; idx < m+n; idx++ {
		ans += strconv.Itoa(ansArr[idx])
	}
	return ans
}

// 小写字母异位词分组-----------------------------------------------------------------------------------------------------
// 思路：用26位数组来保存每个字母出现的次数，出现次数相同的就是异位词
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int32][]string)
	for _, str := range strs {
		var x [26]int32
		for _, s := range str {
			x[s-'a']++
		}
		m[x] = append(m[x], str)
	}
	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 两数相除--------------------------------------------------------------------------------------------------------------
func divide(dividend int, divisor int) int {
	if dividend/divisor > math.MaxInt32 {
		return 1<<31 - 1
	}
	if dividend/divisor < math.MinInt32+1 {
		return math.MinInt32
	}

	return dividend / divisor
}

// 宝石与石头
func numJewelsInStones(jewels string, stones string) int {
	isJewel := make(map[int32]bool)
	for _, v := range jewels {
		isJewel[v] = true
	}
	ans := 0
	for _, v := range stones {
		if isJewel[v] {
			ans++
		}
	}
	return ans
}

// 最长回文子序列  - 子序列定义为：不改变剩余字符顺序的情况下，删除某些字符或者不删除任何字符形成的一个序列
func longestPalindromeSubSeq(s string) int {
	// 由i,j来表示左边界与右边界，如果i > j 就为0
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 边界是逐渐扩大的 固定右边界 移动左边界
	for j := 0; j < n; j++ {
		dp[j][j] = 1                  // i=j，就一个字符说明是1的回文符号
		for i := j - 1; i >= 0; i-- { // 固定右边界 移动左边界
			if s[i] == s[j] { // 如果相同 = 缩小范围2后的字符的最大回文数+2
				dp[i][j] = dp[i+1][j-1] + 2 // 如果i+1>j-1时=2，i+1=j-1时=3
			} else { // 如果不相同 = 减少左边界的字符串 或者 减少右边界字符串 里面最大回文数
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
