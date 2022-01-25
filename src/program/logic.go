package program

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// 罗马数值转普通数字
// 枚举型
func romanToInt(s string) int {
	romans := [...]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	nums := [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	var start int
	var v int
	for i := len(romans) - 1; i >= 0; {
		charLen := len(romans[i])

		end := start + charLen
		fmt.Println(end)
		if end > len(s) || s[start:end] != romans[i] {
			i--
			continue
		}

		start += charLen
		v += nums[i]
	}
	return v
}

// 公共前缀 长度
// 取第一个 然后匹配
func longestCommonPrefix(arr []string) string {
	reply := ""
	if len(arr) < 1 {
		return reply
	} else if len(arr) == 1 {
		return arr[0]
	}

	for k, _ := range arr[0] {
		for k1, v := range arr[1:] {
			if len(v) > k && arr[0][k] == v[k] {
				if k1 == len(arr[1:])-1 {
					reply += string(arr[0][k])
				}
			} else {
				return reply
			}
		}
	}
	return reply
}

// 有效括号
// 用栈匹配 循环一次 将左括号都存起来 后存先匹配机制判断是否取完 - 同类都可以用栈知识
func isValid(s string) bool {
	if s == "" {
		return true
	}

	var stack []uint8

	m := map[uint8]uint8{
		'}': '{',
		')': '(',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			if m[s[i]] != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 合并链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	reply := &ListNode{}

	if l1 == nil && l2 == nil {
		return nil
	}

	nowNode := reply
	for {

		if l1 == nil && l2 == nil {
			nowNode = nil
			return reply
		}

		newNode := &ListNode{}

		if l1 == nil {
			nowNode.Val = l2.Val
			nowNode.Next = newNode
			l2 = l2.Next
		} else if l2 == nil {
			nowNode.Val = l1.Val
			nowNode.Next = newNode
			l1 = l1.Next
		} else {
			if l1.Val > l2.Val {
				nowNode.Val = l2.Val
				nowNode.Next = newNode
				l2 = l2.Next
			} else {
				nowNode.Val = l1.Val
				nowNode.Next = newNode
				l1 = l1.Next
			}
		}
		if l1 == nil && l2 == nil {
			nowNode.Next = nil
			return reply
		}
		nowNode = nowNode.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 把数字翻译成字符串
// 递归玩法 - 每2位数 查下他可以变化的方式，可以就继续变化就探索他的变化，没有就丢弃1位
func TranslateNum(num int) int {
	// 当数字不能在被分解了 就返回
	if num < 10 {
		return 1
	}
	var res int
	// 如果是 10-25 是 可以被解析的值 ， 就多1种变化
	// 如果 其他就 只有1种变化
	if num%100 <= 25 && num%100 > 9 {
		res += TranslateNum(num / 100)
		res += TranslateNum(num / 10)
	} else {
		res += TranslateNum(num / 10)
	}

	return res
}

// 删除排序数组中的重复项
// 双指针 36ms 4.6mb
func RemoveDuplicates(nums []int) int {
	r, l := 0, 1
	for l < len(nums) {
		if nums[r] == nums[l] {
			nums = append(nums[:r], nums[r+1:]...)
		} else {
			r++
			l++
		}
	}
	return len(nums)
}

// 移除元素
// 反向遍历 0ms,2.1mb
func RemoveElement(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
		fmt.Println(nums, i)
	}
	return len(nums)
}

// 重复字符串
// 滑动窗口 双指针 循环头为左指针 右指针才做指针+1开始	8 ms 2.6 MB
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	ls := 0
	for i := 0; i < n; i++ {
		r := i + 1
		m := [128]int{}
		m[s[i]]++
		for r < n && m[s[r]] == 0 {
			m[s[r]]++
			r++
		}
		ls = Max(ls, r-i)
	}
	return ls
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// 实现strStr()
// KMP算法
func StrStr(haystack string, needle string) int {
	lenRoot := len(haystack)
	lenTmpl := len(needle)
	if lenTmpl == 0 {
		return 0
	}

	next := KMPNext(needle)
	for i, j := 0, 0; i < lenRoot; { //i=haystack j=needle
		// 找到 第一次不匹配的位置
		for j < lenTmpl && i < lenRoot && haystack[i] == needle[j] {
			i++
			j++
		}
		// 当这是j已经被匹配完了 就返回
		if j == lenTmpl {
			return i - j
		}
		// 当i被匹配完了，就说明没有
		if i == lenRoot {
			return -1
		}

		// i 每次都会往后移动
		// j 根据返回值 重新定位要开始匹配的位置
		if j > 0 {
			j = next[j-1]
		} else {
			i++
		}
	}
	return -1
}
func KMPNext(s string) []int {
	lenth := len(s)
	next := make([]int, lenth)
	next[0] = 0
	i, j := 1, 0
	for i < lenth {
		if s[i] == s[j] {
			next[i] = j + 1 // 一下个匹配位置为下一位
			i++
			j++
		} else {
			if j == 0 {
				next[i] = 0 // 重头开始匹配
				i++
			} else {
				j = next[j-1] // 回退
			}
		}
	}
	fmt.Println(next)
	return next
}

// 外观数列
// 思路：递归/循环 4ms/6.3mb
func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	// 获取上一n-1的外观数列
	pre := CountAndSay(n - 1)
	res := ""
	count := 1
	for i := 0; i < len(pre); i++ {
		if i+1 < len(pre) && pre[i] == pre[i+1] {

			count++
		} else {
			res += strconv.Itoa(count) + string(pre[i])
			count = 1
		}
	}
	return res
}

// 最大子序和  [-2,1,-3,4,-1,2,1,-5,4]
// 解法:动态规划 8 ms/3.8 MB -- 使用
func MaxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	fmt.Println(nums)
	return max
}

// 一到10W 排序
func Match10WSort() {
	// 生成10W条数据
	arr := make([]int, 100000)
	for i := 1; i <= 100000; i++ {
		arr[i-1] = i
	}

	// 找最大的100个数
	top := make([]int, 0)

	// 乱序
	randSlice(arr)
	//fmt.Println(arr)
	// 比较
	start := time.Now().UnixNano()
	for i := 0; i < 100; i++ {
		t := getMaxNum(arr)
		top = append(top, arr[t])
		arr[t] = 0
	}
	end := time.Now().UnixNano()
	fmt.Println("耗时", (end-start)/1e6)
	fmt.Println(top)
}

func getMaxNum(arr []int) (index int) {
	compare := make([]int, 0)

	for i := 100000; i > 1; {
		if i == 100000 {
			for j := 0; j < len(arr)/2+len(arr)%2; j++ {
				if j*2+1 >= len(arr) {
					compare = append(compare, j*2)
					break
				}
				if arr[j*2] > arr[j*2+1] {
					compare = append(compare, j*2)
				} else {
					compare = append(compare, j*2+1)
				}
			}
			i = i/2 + i%2
			continue
		}
		n := len(compare)
		for j := 0; j < i; j++ {
			if j%2 == 1 {
				continue
			}
			if n-i+j+1 > n+i || arr[compare[n-i+j]] > arr[compare[n-i+j+1]] {
				compare = append(compare, compare[n-i+j])
			} else {
				compare = append(compare, compare[n-i+j+1])
			}
		}

		i = i/2 + i%2
	}
	return compare[len(compare)-1]
}

func randSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	if rv.Type().Kind() != reflect.Slice {
		return
	}

	length := rv.Len()
	if length < 2 {
		return
	}

	swap := reflect.Swapper(slice)
	rand.Seed(time.Now().Unix())
	for i := length - 1; i >= 0; i-- {
		j := rand.Intn(length)
		swap(i, j)
	}
	return
}

// 旋转数组 --------------------------------------------------------------------------------------------------------
// 思路1：根据余数拆分数组然后交换数组
// 思路2: 反转数组，根据k拆分数组再反转，合并
func rotate(nums []int, k int) {
	m := k % len(nums)
	if m == 0 {
		return
	}

	m = len(nums) - m

	arr := nums[m:]
	arr = append(arr, nums[:m]...)

	copy(nums, arr)
}

// 检查数组是否存在重复元素-------------------------------------------------------------------------------------------
// 思路hash
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > 1 {
			return true
		}
	}
	return false
}

// 加一 --------------------------------------------------------------------------------------------------------------
func plusOne(digits []int) []int {
	t := 1
	for i := len(digits) - 1; i >= 0; i-- {
		j := (digits[i]+t)/10 > 0
		digits[i] = (digits[i] + t) % 10
		if j {
			t = 1
		} else {
			t = 0
		}
	}

	if t == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

// 将数组的0移动到最后 --------------------------------------------------------------------------------------
func moveZeroes(nums []int) {
	i, j := 0, 0
	for _, v := range nums {
		if v == 0 {
			j++
		} else {
			if i != j {
				nums[i] = nums[j]
			}
			i++
			j++
		}
	}

	for idx := i; idx < len(nums); idx++ {
		nums[idx] = 0
	}
}

// 九宫格数独-------------------------------------------------------------------------------------------------------
// 思路： 只需要验证已经给出来的数独是不是在同行同列同33里面有相同的即可
func isValidSudoku(board [][]byte) bool {
	var row, column [9][9]int
	var block [3][3][9]int

	for i, v := range board {
		for j, v1 := range v {
			if v1 == '.' {
				continue
			}

			num := v1 - '0' - 1

			row[i][num]++
			column[j][num]++
			block[i/3][j/3][num]++

			if row[i][num] > 1 {
				return false
			}
			if column[j][num] > 1 {
				return false
			}
			if block[i/3][j/3][num] > 1 {
				return false
			}
		}
	}

	return true
}

// 顺时针旋转正方形图像------------------------------------------------------------------------------------------------------
// 思路：不过就是4个位置的交换位置罢了
func rotate2D(matrix [][]int) {
	length := len(matrix)

	for i := 0; i < length/2; i++ {
		for j := i; j < length-i-1; j++ {
			m := length - i - 1 // i的反值
			n := length - j - 1 // j的反值

			fmt.Println(matrix[i][j], matrix[m][j], matrix[m][n], matrix[i][n])

			temp := matrix[i][j]
			matrix[i][j] = matrix[n][i]
			matrix[n][i] = matrix[m][n]
			matrix[m][n] = matrix[j][m]
			matrix[j][m] = temp

			fmt.Println(matrix[i][j], matrix[m][j], matrix[m][n], matrix[i][n])

		}
	}
}

// 排列序列    找由n个数字组成排序的第k个数
func getPermutation(n int, k int) string {
	// 获取各个位数的递归值 n!
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}

	ans := ""

	// 已用数字标记
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}

	// 因为 k <= (n-1)! 为了方便计算排序除所以k-1
	k--
	for i := 1; i <= n; i++ {
		// 获取当前位数的排序该有的
		order := k/factorial[n-i] + 1

		// 根据排序找到当前没被使用的当前位数的数字
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}

		// 摸掉已经被使用的位数阶乘
		k %= factorial[n-i]
	}
	return ans
}
