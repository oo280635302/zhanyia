package program

import (
	"sort"
	"strconv"
	"unicode"
)

// 电话号码的字母组合
// 思路：递归,每次都找对应数值的几个不同对应值 当每个的剩余数为空时 就是找到底了 就将其追加到返回值中	0ms/2mb
var phone = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func LetterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	res := make([]string, 0)
	dfsLetterCombinations("", digits, &res)
	return res
}

func dfsLetterCombinations(nowStr string, digits string, res *[]string) {
	if len(digits) == 0 {
		*res = append(*res, nowStr)
		return
	}

	for _, v := range phone[digits[0]] {
		a := nowStr + v
		dfsLetterCombinations(a, digits[1:], res)
	}
}

// 复原 IP 地址
// 思路：回溯，先枚举每一步的各种可能性然后dfs找到最下一步的可能性，如果能组成ip就算结果之一，如果不能组成就回溯进行下一个可能性
var (
	ans      []string
	segments []int
)

func restoreIpAddresses(s string) []string {
	segments = make([]int, 4)
	ans = []string{}
	dfsRestoreIpAddresses(s, 0, 0)
	return ans
}

func dfsRestoreIpAddresses(s string, segId, segStart int) {
	// 如果找到了 4 段 IP 地址并且遍历完了字符串，那么就是一种答案
	if segId == 4 {
		if segStart == len(s) {
			ipAddr := ""
			for i := 0; i < 4; i++ {
				ipAddr += strconv.Itoa(segments[i])
				if i != 4-1 {
					ipAddr += "."
				}
			}
			ans = append(ans, ipAddr)
		}
		return
	}

	// 如果还没有找到 4 段 IP 地址就已经遍历完了字符串，那么提前回溯
	if segStart == len(s) {
		return
	}
	// 由于不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
	if s[segStart] == '0' {
		segments[segId] = 0
		dfsRestoreIpAddresses(s, segId+1, segStart+1)
	}
	// 一般情况，枚举每一种可能性并递归
	addr := 0
	for segEnd := segStart; segEnd < len(s); segEnd++ {
		addr = addr*10 + int(s[segEnd]-'0')
		if addr > 0 && addr <= 0xFF {
			segments[segId] = addr
			dfsRestoreIpAddresses(s, segId+1, segEnd+1)
		} else {
			break
		}
	}
}

// 最多可达成的换楼请求数目
// 思路：回溯+枚举
func maximumRequests(n int, requests [][]int) (ans int) {
	delta := make([]int, n)
	cnt, zero := 0, n
	var dfs func(int)
	dfs = func(pos int) {
		if pos == len(requests) {
			if zero == n && cnt > ans {
				ans = cnt
			}
			return
		}

		// 不选 requests[pos]
		dfs(pos + 1)

		// 选 requests[pos]
		z := zero
		cnt++
		r := requests[pos]
		x, y := r[0], r[1]
		if delta[x] == 0 {
			zero--
		}
		delta[x]--
		if delta[x] == 0 {
			zero++
		}
		if delta[y] == 0 {
			zero--
		}
		delta[y]++
		if delta[y] == 0 {
			zero++
		}
		dfs(pos + 1)
		delta[y]--
		delta[x]++
		cnt--
		zero = z
	}
	dfs(0)
	return
}

// 强密码检验器
func strongPasswordChecker(password string) int {
	hasLower, hasUpper, hasDigit := 0, 0, 0
	for _, ch := range password {
		if unicode.IsLower(ch) {
			hasLower = 1
		} else if unicode.IsUpper(ch) {
			hasUpper = 1
		} else if unicode.IsDigit(ch) {
			hasDigit = 1
		}
	}
	categories := hasLower + hasUpper + hasDigit

	switch n := len(password); {
	case n < 6:
		return max(6-n, 3-categories)
	case n <= 20:
		replace, cnt, cur := 0, 0, '#'
		for _, ch := range password {
			if ch == cur {
				cnt++
			} else {
				replace += cnt / 3
				cnt = 1
				cur = ch
			}
		}
		replace += cnt / 3
		return max(replace, 3-categories)
	default:
		// 替换次数和删除次数
		replace, remove := 0, n-20
		// k mod 3 = 1 的组数，即删除 2 个字符可以减少 1 次替换操作
		rm2, cnt, cur := 0, 0, '#'
		for _, ch := range password {
			if ch == cur {
				cnt++
				continue
			}
			if remove > 0 && cnt >= 3 {
				if cnt%3 == 0 {
					// 如果是 k % 3 = 0 的组，那么优先删除 1 个字符，减少 1 次替换操作
					remove--
					replace--
				} else if cnt%3 == 1 {
					// 如果是 k % 3 = 1 的组，那么存下来备用
					rm2++
				}
				// k % 3 = 2 的组无需显式考虑
			}
			replace += cnt / 3
			cnt = 1
			cur = ch
		}

		if remove > 0 && cnt >= 3 {
			if cnt%3 == 0 {
				remove--
				replace--
			} else if cnt%3 == 1 {
				rm2++
			}
		}

		replace += cnt / 3

		// 使用 k % 3 = 1 的组的数量，由剩余的替换次数、组数和剩余的删除次数共同决定
		use2 := min(min(replace, rm2), remove/2)
		replace -= use2
		remove -= use2 * 2

		// 由于每有一次替换次数就一定有 3 个连续相同的字符（k / 3 决定），因此这里可以直接计算出使用 k % 3 = 2 的组的数量
		use3 := min(replace, remove/3)
		replace -= use3
		remove -= use3 * 3
		return (n - 20) + max(replace, 3-categories)
	}
}

// 全排列
func permute(nums []int) [][]int {
	res, path := make([][]int, 0), make([]int, 0, len(nums))
	isUse := make(map[int]bool, 0)

	var dfs = func(cur int) {}
	dfs = func(cur int) {
		// 已使用数量 和 数组数量相等等于到底了
		if cur == len(nums) {
			// []int是指针类型 要深copy 不然之前的数组会改变
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for _, num := range nums {
			if !isUse[num] {
				path = append(path, num)
				isUse[num] = true
				dfs(cur + 1)
				path = path[:len(path)-1]
				isUse[num] = false
			}
		}
	}
	// 从0 开始
	dfs(0)
	return res
}

// 全排列 II
func permuteUnique(nums []int) [][]int {
	res, path := make([][]int, 0), make([]int, 0, len(nums))
	isUse := make(map[int]bool, 0)
	sort.Ints(nums)

	var dfs = func(cur int) {}
	dfs = func(cur int) {
		// 已使用数量 和 数组数量相等等于到底了
		if cur == len(nums) {
			// []int是指针类型 要深copy 不然之前的数组会改变
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for idx, num := range nums {
			// 跳过已经被使用的索引
			if isUse[idx] {
				continue
			}
			// 跳过 同层级（!isUse[idx-1]）当前数与前一个数相等(nums[idx] == nums[idx-1] )的情况
			if idx > 0 && nums[idx] == nums[idx-1] && !isUse[idx-1] {
				continue
			}

			path = append(path, num)
			isUse[idx] = true
			dfs(cur + 1)
			path = path[:len(path)-1]
			isUse[idx] = false
		}
	}
	// 从0 开始
	dfs(0)
	return res
}
