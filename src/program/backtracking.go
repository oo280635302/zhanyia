package program

import "strconv"

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
