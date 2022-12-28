package program

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

// 邻位交换的最小次数
func getMinSwaps(num string, k int) int {
	n := len(num)

	// 获取第k个最小秒数
	lastNumByte := []int32(num)
	kNumByte := []int32(num)
	for i := 0; i < k; i++ {
		nextPermutation(kNumByte)
	}

	ans := 0

	// 贪心：找到最小交换次数
	// 每次遇到不相同的数都从右边找最近的相同的数然后整体移动过来，
	for i := 0; i <= n-1; i++ {

		if lastNumByte[i] != kNumByte[i] { // 从左到右找到不同的数 i
			for j := i + 1; j <= n-1; j++ {
				if lastNumByte[i] == kNumByte[j] { // 从右边找到 i相同的j的位置

					for k := j; k > i; k-- { // 将j逐渐移动i位置，并记录步数
						kNumByte[k], kNumByte[k-1] = kNumByte[k-1], kNumByte[k]
						ans++
					}
					break
				}
			}
		}
	}

	return ans
}

// 句子中的有效单词数
func countValidWords(sentence string) int {
	words := strings.Fields(sentence)
	ans := 0
	for _, word := range words {
		if validWord(word) {
			ans++
		}
	}

	return ans
}
func validWord(s string) bool {
	n := len(s)
	cnt := 0
	for idx, v := range s {
		if !(v == '-' || v == '!' || v == ',' || v == '.' || (v >= 'a' && v <= 'z')) {
			return false
		}
		if v == '-' {
			if cnt != 0 {
				return false
			}
			if idx <= 0 || idx >= n-1 || s[idx-1] < 'a' || s[idx-1] > 'z' || s[idx+1] < 'a' || s[idx+1] > 'z' {
				return false
			}
			cnt++
		}

		if (v == '!' || v == '.' || v == ',') && idx != n-1 {
			return false
		}
	}
	return true
}

// 比较版本号
// 思路：双指针，跟据获取到的当前版本号阶段数字进行对比
func compareVersion(version1 string, version2 string) int {
	m, n := len(version1), len(version2)
	i, j := 0, 0

	for i < m || j < n {
		cur1 := 0
		cur2 := 0
		for i < m && version1[i] != '.' { // version1当前阶段的数字
			cur1 = cur1*10 + int(version1[i]-'0')
			i++
		}
		i++ // 跳过逗号

		for j < n && version2[j] != '.' { // version2当前阶段的数字
			cur2 = cur2*10 + int(version2[j]-'0')
			j++
		}
		j++ // 跳过逗号

		if cur1 > cur2 {
			return 1
		}

		if cur1 < cur2 {
			return -1
		}
	}

	return 0
}

// “气球” 的最大数量 balloon
// 思路：统计每个字母的出现次数，求最小即可
func maxNumberOfBalloons(text string) int {
	m := make(map[int32]int)

	for _, v := range text {
		m[v]++
	}

	ans := 10000
	if m['b'] < ans {
		ans = m['b']
	}
	if m['a'] < ans {
		ans = m['a']
	}
	if m['l']/2 < ans {
		ans = m['l'] / 2
	}
	if m['o']/2 < ans {
		ans = m['o'] / 2
	}
	if m['n'] < ans {
		ans = m['n']
	}

	return ans
}

// 1比特与2比特字符
// 思路: 遇到1记录2步 遇到0记录1步，记录走完除最后一位数的全程的需要的步数 是否与n-1相等
func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	step := 0

	for i := 0; i < n-1; i++ {
		if bits[i] == 1 {
			step += 2
			i++
		} else {
			step += 1
		}
		// fmt.Println(bits[i],n,step,i)
	}

	return n-1 == step
}

// 推多米诺
// 模拟：模拟真实情况，推导过程、还有另外的解 广度遍历
func pushDominoes(dominoes string) string {
	ans := []byte(dominoes)

	l, n := 0, len(dominoes)
	left := 'L'

	for l < n {
		// 找到没有被推动的一段
		r := l
		for r < n && dominoes[r] == '.' {
			r++
		}
		//fmt.Println(l,r)

		// 除了最后一段设右边R以外 其他情况都以当前L、R为右边界
		right := 'R'
		if r < n {
			right = int32(dominoes[r])
		}

		// 如果相等是同向 填充跳过的.
		if left == right {
			for l < r {
				ans[l] = byte(left)
				l++
			}
			// 如果不相等 同时是R-L的分布 就互推
		} else if left == 'R' {
			r--
			for l < r {
				ans[l], ans[r] = 'R', 'L'
				l++
				r--
			}
		}

		// 记录左边界
		left = right
		l = r + 1
	}

	return string(ans)
}

// 仅仅反转字母
// 思路：双指针，遇到非字母就跳过
func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	l, r := 0, len(s)-1

	for l < r {
		if !((ans[l] >= 'A' && ans[l] <= 'Z') || (ans[l] >= 'a' && ans[l] <= 'z')) {
			l++
			continue
		}
		if !((ans[r] >= 'A' && ans[r] <= 'Z') || (ans[r] >= 'a' && ans[r] <= 'z')) {
			r--
			continue
		}
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}

	return string(ans)
}

// 最优除法
// 思路：所有数想除获得最大的值，因为都是正整数相除只会越来越小，所以只需要保证 左边固定idx=0 / (右边相除) ，分母因为相除越来越小，分子固定就越来越大
func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	ans := ""
	for idx, val := range nums {
		if idx == 0 {
			ans += strconv.Itoa(val)
		} else if idx == 1 {
			ans += "/(" + strconv.Itoa(val)
		} else if idx == len(nums)-1 {
			ans += "/" + strconv.Itoa(val) + ")"
		} else {
			ans += "/" + strconv.Itoa(val)
		}
	}

	return ans
}

// 寻找最近的回文数
func nearestPalindromic(n string) string {
	m := len(n)
	candidates := []int{int(math.Pow10(m-1)) - 1, int(math.Pow10(m)) + 1}
	selfPrefix, _ := strconv.Atoi(n[:(m+1)/2])
	for _, x := range []int{selfPrefix - 1, selfPrefix, selfPrefix + 1} {
		y := x
		if m&1 == 1 {
			y /= 10
		}
		for ; y > 0; y /= 10 {
			x = x*10 + y%10
		}
		candidates = append(candidates, x)
	}

	ans := -1
	selfNumber, _ := strconv.Atoi(n)
	for _, candidate := range candidates {
		if candidate != selfNumber {
			if ans == -1 ||
				abs(candidate-selfNumber) < abs(ans-selfNumber) ||
				abs(candidate-selfNumber) == abs(ans-selfNumber) && candidate < ans {
				ans = candidate
			}
		}
	}
	return strconv.Itoa(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 最长特殊序列 Ⅰ
// 思路： 这TM什么阅读理解题
func findLUSlength(a, b string) int {
	if a != b {
		return max(len(a), len(b))
	}
	return -1
}

// 蜡烛之间的盘子
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	preSum := make([]int, n) // 当前节点累计的球数
	left := make([]int, n)   // 当前节点最近的左蜡烛 - 包括自身
	sum, l, r := 0, -1, -1
	for i, ch := range s {
		if ch == '*' {
			sum++
		} else {
			l = i
		}
		preSum[i] = sum
		left[i] = l
	}

	right := make([]int, n)
	for i := n - 1; i >= 0; i-- { // 当前节点最近的右蜡烛
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		x, y := right[q[0]], left[q[1]]
		if x >= 0 && y >= 0 && x < y {
			ans[i] = preSum[y] - preSum[x] //  右节点最近的左蜡烛的球数 - 左节点的最近的右蜡烛的球数
		}
	}
	return ans
}

// 两个列表的最小索引总和
func findRestaurant(list1 []string, list2 []string) []string {
	ans := []string{}

	// 建立list1的map
	m1 := make(map[string]int, 0)
	for idx, val := range list1 {
		m1[val] = idx
	}

	min := 1000000

	for idx, val := range list2 {
		if num, ok := m1[val]; ok {
			if num+idx < min {
				min = num + idx
				ans = []string{val}
			} else if num+idx == min {
				ans = append(ans, val)
			}
		}
	}

	return ans
}

// 词典中最长的单词
func longestWord(words []string) string {
	// 排序 长度升序，字典序逆序（目的是为了让相同长度的字符字典序最小的排最后 让最终结果的字典序最小在最后弹出去）
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j]) || (len(words[i]) == len(words[j]) && words[i] > words[j])
	})

	ans := ""
	m := make(map[string]bool)
	m[""] = true
	for _, v := range words {
		if m[v[:len(v)-1]] == true { // 因为升序，所以能匹配到的 单词长度一定>前一个 || 单词长度=前一个&&字典序<前一个
			ans = v
			m[v] = true
		}
	}

	return ans
}

// 如果相邻两个颜色均相同则删除当前颜色
func winnerOfGame(colors string) bool {
	ans := 0 // 只有3个相连的颜色才会抵消中间那个，且被抵消后的不会让相反的颜色组成3连，因此只需要计算能组成3连的次数就行
	n := len(colors)

	for idx, val := range colors {
		if idx-1 >= 0 && idx+1 < n {
			if val == 'A' && colors[idx-1] == 'A' && colors[idx+1] == 'A' {
				ans++
			}
			if val == 'B' && colors[idx-1] == 'B' && colors[idx+1] == 'B' {
				ans--
			}
		}
	}
	return ans > 0 // Alice先手所以ans必须>0
}

// 环绕字符串中唯一的子字符串  环形字符串=abcdef...zabcdef
func findSubstringInWraproundString(p string) int {
	dp := [26]int{}

	// 能组成环字符串的子串的 一定是连续的英文字符串 如 abc 找出每个字母最长的连续字串数量和 就是 唯一字串数量
	// 用动规的方式找到 每个字母的最长连续字符串

	k := 1
	for idx, val := range p {

		// 如果前一个字符是当前字符的连续 k++
		if idx > 0 && (val-int32(p[idx-1])+26)%26 == 1 {
			k++
			// 不是就重置
		} else {
			k = 1
		}
		// 每个字母最大的连续子串
		dp[int(val)-'a'] = max(dp[int(val)-'a'], k)
	}

	ans := 0
	for _, v := range dp {
		ans += v
	}
	return ans
}

// 生成每种字符都是奇数个的字符串
func generateTheString(n int) string {
	// 如果本身n 就是奇数之间返回
	if n%2 == 1 {
		return strings.Repeat("a", n)
		// 如果是偶数 就是n-1奇数 + 1奇数
	} else {
		return strings.Repeat("a", n-1) + strings.Repeat("b", 1)
	}
}

// Second Largest Digit in a String
func secondHighest(s string) int {
	first := int32(-1)
	second := int32(-1)
	for _, n := range s {
		if n < '0' || n > '9' {
			continue
		}
		num := n - '0'
		if num > first {
			second = first
			first = num
		} else if num < first && num > second {
			second = num
		}
	}

	return int(second)
}

// Number of Different Integers in a String
// 思路：双指针 找到数字开始往右边找直到找到非数字，再把左指针移动到右指针处继续找数字
func numDifferentIntegers(word string) int {
	m := make(map[string]bool)

	left := 0
	for left < len(word) {
		leftIsDigit := unicode.IsDigit(rune(word[left]))
		if !leftIsDigit {
			left++
			continue
		}

		cur := ""
		right := left
		for right < len(word) && unicode.IsDigit(rune(word[right])) {
			if word[right] == '0' && cur == "" {
				right++
				continue
			}
			cur += string(word[right])
			right++
		}
		fmt.Println(cur)
		m[cur] = true
		left = right
	}

	return len(m)
}

// 执行操作后的变量值
func finalValueAfterOperations(operations []string) int {
	res := 0
	for _, val := range operations {
		switch val {
		case "++X", "X++":
			res += 1
		case "--X", "X--":
			res -= 1
		}
	}
	return res
}

// 统计同构子字符串的数目
func countHomogenous(s string) (res int) {
	prev := rune(s[0])
	cnt := 0
	for _, c := range s {
		// 遇到相同字母就计数
		if c == prev {
			cnt++
			// 遇到不同字母就根据长度 算出排列 (n*n-1...*2*1) = (n+1)*n/2
		} else {
			res += (cnt + 1) * cnt / 2
			cnt = 1
			prev = c
		}
	}
	// 遇到最后的再算一次数
	res += (cnt + 1) * cnt / 2
	return res % (1e9 + 7)
}

// 构造字典序最大的合并字符串
func largestMerge(word1 string, word2 string) string {
	ans := ""
	for len(word1) != 0 || len(word2) != 0 {
		l, r := int32(0), int32(0)
		if len(word1) > 0 {
			l = int32(word1[0])
		}
		if len(word2) > 0 {
			r = int32(word2[0])
		}
		// 比较两个字符串的第一个字符：如果遇到大的就直接追加，遇到相等的就再比较下字典序
		if l > r {
			ans += string(l)
			word1 = word1[1:]
		} else if l < r {
			ans += string(r)
			word2 = word2[1:]
		} else {
			if word1 > word2 {
				ans += string(l)
				word1 = word1[1:]
			} else {
				ans += string(r)
				word2 = word2[1:]
			}
		}
	}

	return ans
}
