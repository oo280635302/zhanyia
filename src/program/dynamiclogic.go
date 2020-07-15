package program

import (
	"fmt"
	"math"
)

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

// 不同路径2-----------------------------------------------------------------------------------------------------------
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

// 跳水板-----------------------------------------------------------------------------------------------------------
// 思路: 因为k=大小板总和  因此排列方式有k+1种 0个短板->k个短板  20ms/7.1mb
func divingBoard(shorter int, longer int, k int) []int {
	res := make([]int, 0)
	if k == 0 {
		return nil
	}
	if longer == shorter {
		res = append(res, k)
		return res
	}
	for i := 0; i <= k; i++ {
		res = append(res, shorter*(k-i)+longer*(i))
	}
	return res
}

// 恢复空格-----------------------------------------------------------------------------------------------------------
// 思路:字典树Trie + 动态规划 96ms/56.8mb
func respace(dictionary []string, sentence string) int {
	n, inf := len(sentence), 0x3f3f3f3f
	root := &Trie{next: [26]*Trie{}}
	for _, word := range dictionary {
		root.insert(word)
	}
	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = inf
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		curPos := root
		for j := i; j >= 1; j-- {
			t := int(sentence[j-1] - 'a')
			if curPos.next[t] == nil {
				break
			} else if curPos.next[t].isEnd {
				dp[i] = min(dp[i], dp[j-1])
			}
			if dp[i] == 0 {
				break
			}
			curPos = curPos.next[t]
		}
	}
	return dp[n]
}

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func (this *Trie) insert(s string) {
	curPos := this
	for i := len(s) - 1; i >= 0; i-- {
		t := int(s[i] - 'a')
		if curPos.next[t] == nil {
			curPos.next[t] = &Trie{next: [26]*Trie{}}
		}
		curPos = curPos.next[t]
	}
	curPos.isEnd = true
}

// 最佳买卖股票时机含冷冻期--------------------------------------------------------------------------------------------
// 思路:dynamic看注释 将操作状态分为3种,到每一股时都根据前一个状态的最佳收益计算出本次状态的最佳收益 	0ms/2.5mb
func MaxProfitByFreeze(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	// f[i][0]: 手上持有股票的最大收益累计
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	f := make([][3]int, n)
	f[0][0] = -prices[0]
	for i := 1; i < n; i++ {
		f[i][0] = max(f[i-1][0], f[i-1][2]-prices[i]) // 说明当前手上一定有股票, 两种情况 本身就有f[i-1][0] 或者 不属于冰冻期这次购买f[i-1][2]-prices[i]
		f[i][1] = f[i-1][0] + prices[i]               // 说明 前一天有股票这次卖出 即纯收入
		f[i][2] = max(f[i-1][1], f[i-1][2])           // 说明 前一天必定 不持有股票 要么处于冰冻期f[i-1][1] 要么处于非冰冻期f[i-1][2]
	}
	fmt.Println(f)
	return max(f[n-1][1], f[n-1][2]) // 最终结果必定从不持有股票中选举出来  因为持有股票最佳也仅仅是持平的金额
}

// 不同路径1-----------------------------------------------------------------------------------------------------------
// 思路:到当前点的路径 = 到起左边点的路径 + 到起上边点路径的;最左边与最上边的路径都是1 0ms/2mb
func UniquePaths(m int, n int) int {

	if m == 0 || n == 0 {
		return 0
	}

	path := make([][]int, m)
	for i := 0; i < m; i++ {
		path[i] = make([]int, n)
	}
	path[0][0] = 1

	for i := 0; i < len(path); i++ {
		for j := 0; j < len(path[0]); j++ {
			if i == 0 || j == 0 {
				path[i][j] = 1
				continue
			}
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}

	return path[m-1][n-1]
}

// 最小路径和-----------------------------------------------------------------------------------------------------------
// 思路: 动态规划 每个路径的和都是右边或者下边中选择路径和最小的 从右下角开始找 每个格子都找最小路径和 一直往左上角找  8ms/3.9mb
// 优化:该题甚至都不用建立 二维动态图 利用grid就可以 节约了0.6mb的内存
func MinPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				continue
			} else if i == m-1 && j != n-1 {
				grid[i][j] = grid[i][j] + grid[i][j+1]
			} else if i != m-1 && j == n-1 {
				grid[i][j] = grid[i][j] + grid[i+1][j]
			} else {
				grid[i][j] = grid[i][j] + min(grid[i+1][j], grid[i][j+1])
			}
		}
	}

	return grid[0][0]
}

// 地下城游戏-----------------------------------------------------------------------------------------------------------
// 思路：动态规划  根据右下角生命值 一直往上 推到出最优的开始计算的生命值 (思路与最小路径和相同) 4ms/3.7mb
func calculateMinimumHP(dungeon [][]int) int {
	// 建立 二维图
	m, n := len(dungeon), len(dungeon[0])
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	// m-1,n-1值由其右边 和 下边 两值选举出来
	dp[m][n-1], dp[m-1][n] = 1, 1

	// 从右下角开始 计算出要往下一格走需要的生命值
	// 从右边和下边两个生命格中选出需要最少生命的格子， 走这个格子需要的生命= 下个格子带来的生命 - 这个格子能+/-的生命 但最低只能为1
	// 依次向上推导 出0,0需要的最低生命
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			minn := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max(minn-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}

// 爬楼梯--------------------------------------------------------------------------------------------------------------
// 思路:动态规划  因为只能1,2个台阶走法: 该台阶 = 该台阶-1 + 台阶-2,0号台阶与1号台阶 走法都只有1种 0ms/1.9mb
// 优化:反复使用2个变量 让空间复杂度o(1)
func ClimbStairs(n int) int {
	if n <= 1 {
		return 1
	}
	res1, res2 := 1, 1

	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			res1 = res1 + res2
		} else {
			res2 = res1 + res2
		}
	}

	return max(res1, res2)
}

// 买卖股票的最佳时机---------------------------------------------------------------------------------------------------
// 思路: 有两种操作 买/卖 最优化买:对比 上一次买与这次买 最省钱的方式,最优化卖:对比 上次卖与这次卖 最赚钱的方式 4ms/3.6mb
// 优化: 内存可以压缩 不用建立二维数组 重复使用买卖变量 4ms/3.1mb
func maxProfit(prices []int) int {
	// 0 = 持有股票
	// 1 = 不持有股票
	n := len(prices)
	if n == 0 {
		return 0
	}

	buy := -prices[0]
	sell := 0

	for i := 1; i < n; i++ {
		// 这个时候购入
		tBuy := max(buy, -prices[i])
		// 这个时候卖出
		tSell := max(buy+prices[i], sell)
		buy, sell = tBuy, tSell
	}
	return sell
}

// 三角形最小路径和-----------------------------------------------------------------------------------------------------
// 思路: 从下往上 每次都给上一层提供下一层相邻两个最小的,直到最上层就是最小的路径和 4ms/3.1mb
// 优化: 直接利用triangle 不浪费额外内存 （实际情况应该不能这样做）
func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	if n == 1 {
		return triangle[0][0]
	}

	for i := n - 2; i >= 0; i-- {
		for k, v := range triangle[i] {
			triangle[i][k] = v + min(triangle[i+1][k], triangle[i+1][k+1])
		}
	}

	return triangle[0][0]
}

// 编辑距离-----------------------------------------------------------------------------------------------------------
// 思路：展开w1,w2两个字符串 然后用依次匹配的模式来获取最终 当前w1字符与w2字符相等时不用修改 不同时需要在前一步中找出最优解 同时+1 4ms/5.7mb
func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			// 当前位置的数相同 = 各数前一个匹配的值 (因为相同 需要前面的数对齐)
			// 当前位置的数不同 = 需要从周边数中找到最优的匹配 + 1 (加1是因为自身数不同)
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i-1][j-1]), dp[i][j-1]) + 1
			}
		}
	}

	return dp[m][n]
}

// 解码方法----------------------------------------------------------------------------------------------------
// 思路:
func numDecodings(s string) int {
	return 1
}
