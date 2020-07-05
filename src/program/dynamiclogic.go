package program

// 动态规划有关的算法问题 -- LeetCode

// 单词拆分---------------------------------------------------------------------------------------------------
// 思路：双指针匹配移动 先将数组转成map方便匹配 再将数据能匹配的点记录下来 根据上一个正确点匹配下一个正确点 最终查看末尾点是否被正确匹配
func WordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	// 0点 默认是 空 被任意正确匹配
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// 当头点正确时 - 找出能正确全部匹配到的点记录下正确的下一个点
			if dp[j] && wordDictSet[s[j:i]] {
				//fmt.Println(i,j)
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 最长重复子数组-----------------------------------------------------------------------------------------------
// 思路：滑动窗口 以A、B各为固定点 移动B、A 每次移动都 遍历A B求出其当前最长重复子数组 然后汇总 32ms/3.4mb
// 多解题思路：动态规划（弱）  二分hash法（强）
func findLength(A []int, B []int) int {
	n, m := len(A), len(B)
	ret := 0
	for i := 0; i < n; i++ {
		lenA := min(m, n-i)
		maxLen := maxLength(A, B, i, 0, lenA)
		ret = max(ret, maxLen)
	}
	for i := 0; i < m; i++ {
		lenB := min(n, m-i)
		maxLen := maxLength(A, B, 0, i, lenB)
		ret = max(ret, maxLen)
	}
	return ret
}

func maxLength(A, B []int, addA, addB, len int) int {
	ret, k := 0, 0
	for i := 0; i < len; i++ {
		if A[addA+i] == B[addB+i] {
			k++
		} else {
			k = 0
		}
		ret = max(ret, k)
	}
	return ret
}

// 动态规划解法: 基于暴力解法之上，保存 i,j 位置 当i,j相等时 他们的公共长度 是i+1,j+1的公共长度+1 否则他们公共长度为0
func findLength2(A []int, B []int) int {
	n, m := len(A), len(B)
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if A[i] == B[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			if ans < dp[i][j] {
				ans = dp[i][j]
			}
		}
	}
	return ans
}
