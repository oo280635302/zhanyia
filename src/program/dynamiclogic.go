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

// 通配符匹配
// 思路：动态规划 先设定 0,0 是true 再依次匹配1,1根据0,0 && 1?=1 来决定 最终到m,n是否是正确来判断整体是否匹配
// * 可以将任意一条n列 全部污染成true 哦~
func IsMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	// 先解决 以*开头的情况
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

// 正则表达式匹配---------------------------------------------------------------------------------------------------
// 思路: 动态规划 思路同 通配符匹配 ,不同处在于正则*是 前一个字母的复制体 同时可以将前一个字母置为0 因此遇到* 只需要其前2位数相同即可 4ms/2.4mb
func IsMatchRegexp(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}

	// 检测两数是否相等
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	dp[0][0] = true

	for i := 0; i <= m; i++ { // 第一排是为了 给* 找空间
		for j := 1; j <= n; j++ {

			if p[j-1] == '*' {

				// 当遇到 * 时 因为有可能会将前一个数无效化 需要看其j-2位是否正确
				dp[i][j] = dp[i][j] || dp[i][j-2]
				// 在有效化的情况时 只需要位于s前一位是正确的 他也是正确的 aa == a* 当aa相等
				if matches(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}

				// 正常的数字 只需要对应匹配即可 同时i-1,j-1正确即可
			} else if matches(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}

		}
	}

	return dp[m][n]
}

// 不同路径
// 思路：动态规划 要想知道 到达1,1的路径 == 0,1 + 1,0的路径和以此推论 遇到石头说明此路不通 0ms/2.4mb
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, n)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}
