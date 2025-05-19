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

// 解码方法----------------------------------------------------------------------------------------------------------
// 思路:
func numDecodings(s string) int {
	return 1
}

// 最长回文串--------------------------------------------------------------------------------------------------------
// 思路: 用动规保存 当前i,j相等 就查看i-1与j+1是否相等 然后与之前保存的最长回文串比较长度  108ms/15.3ms
func LongestPalindrome(s string) string {
	n := len(s)

	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	reply := ""

	// x为i,j的距离
	for x := 0; x < n; x++ {
		for i := 0; i < n; i++ {

			j := i + x
			if j < n && s[i] == s[j] {
				if i == j || i+1 == j {
					dp[i][j] = 1
				} else if i+1 <= j-1 && dp[i+1][j-1] == 1 {
					dp[i][j] = 1
				} else {
					continue
				}

				if j-i+1 > len(reply) {
					reply = s[i : j+1]
				}
			}
		}
	}

	return reply
}

// 交错字符串 --------------------------------------------------------------------------------------------------------
// 思路：建立dp图， 只要当前数与s3相同 同时邻近的两边有一边能组成s3之前的数时 意味着当前s1[:i+1] + s2[:j+1] 构成交错字符串 s3[:i+j+1]  0ms/2.1mb
func IsInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}

	return f[n][m]
}

// 戳气球-------------------------------------------------------------------------------------------------------------
// 思路：反向思考,戳气球转为插入气球直到插满 插入一个气球时其值=左区间和+当前值+右区间和 选择最优的和赋值
func MaxCoins(nums []int) int {
	n := len(nums)
	// dp图
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	// 加入左右两边1
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}
	// 寻找最大值
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				// 当前值 = 左区间 + 本身值 + 右区间
				sum := val[i] * val[k] * val[j]
				sum += dp[i][k] + dp[k][j]
				// 选择最优值 即 有可能
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}

	return dp[0][n+1]
}

// 除数博弈-----------------------------------------------------------------------------------------------------------
// 思路: 简单的寻找规律题 0ms/2mb
func divisorGame(N int) bool {
	return N%2 == 0
}

// 分割数组的最大值---------------------------------------------------------------------------------------------------
// 思路：将数组所有段都进行分段比较  36ms/2.5mb
func SplitArray(nums []int, m int) int {
	n := len(nums)

	f := make([][]int, n+1)
	sub := make([]int, n+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, m+1)
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = math.MaxInt32
		}
	}
	// 数组各数的前数累加值
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}
	f[0][0] = 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := 0; k < i; k++ {
				// 其
				f[i][j] = min(f[i][j], max(f[k][j-1], sub[i]-sub[k]))
			}
		}
	}

	return f[n][m]
}

// 矩阵中的最长递增路径--------------------------------------------------------------------------------------------------
// 思路： 拓扑排序 52 ms,6.8 MB
func longestIncreasingPath(matrix [][]int) int {
	var (
		dirs          = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
		rows, columns int
	)
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, columns = len(matrix), len(matrix[0])
	outdegrees := make([][]int, rows)
	for i := 0; i < rows; i++ {
		outdegrees[i] = make([]int, columns)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			for _, dir := range dirs {
				newRow, newColumn := i+dir[0], j+dir[1]
				if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] > matrix[i][j] {
					outdegrees[i][j]++
				}
			}
		}
	}

	queue := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if outdegrees[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}
	ans := 0
	for len(queue) != 0 {
		ans++
		size := len(queue)
		for i := 0; i < size; i++ {
			cell := queue[0]
			queue = queue[1:]
			row, column := cell[0], cell[1]
			for _, dir := range dirs {
				newRow, newColumn := row+dir[0], column+dir[1]
				if newRow >= 0 && newRow < rows && newColumn >= 0 && newColumn < columns && matrix[newRow][newColumn] < matrix[row][column] {
					outdegrees[newRow][newColumn]--
					if outdegrees[newRow][newColumn] == 0 {
						queue = append(queue, []int{newRow, newColumn})
					}
				}
			}
		}
	}
	return ans
}

// 爬楼梯
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	clim := make([]int, n+1)
	clim[1] = 1
	clim[2] = 2
	for i := 3; i <= n; i++ {
		clim[i] = clim[i-1] + clim[i-2]
	}

	return clim[n]
}

// 统计元音字母序列的数目
// 思路：动态规划，规划出每个序列的次数，再跟据上个尾元音判断当前序列每个原因出现的次数
func countVowelPermutation(n int) int {
	//tab := map[byte][]byte{
	//	'a':{'e'},
	//	'e':{'a','i'},
	//	'i':{'a','e','o','u'},
	//	'o':{'i','u'},
	//	'u':{'a'},
	//}

	m := make([][]int, 0)
	m = append(m, []int{1, 1, 1, 1, 1})
	res := 5

	for i := 0; i < n-1; i++ {
		tmp := m[i]
		cur := make([]int, 5)
		cur[0] = (tmp[1] + tmp[2] + tmp[4]) % 1000000007
		cur[1] = (tmp[0] + tmp[2]) % 1000000007
		cur[2] = (tmp[1] + tmp[3]) % 1000000007
		cur[3] = (tmp[2]) % 1000000007
		cur[4] = (tmp[2] + tmp[3]) % 1000000007
		m = append(m, cur)
		res = (tmp[0] + 2*tmp[1] + 4*tmp[2] + 2*tmp[3] + tmp[4]) % 1000000007
	}

	return res
}

// 第 k 个数  3，5，7 组成的的k个数
func getKthMagicNumber(k int) int {
	dp := make([]int, k)
	dp[0] = 1

	// 标记3、5、7 三个数当前最小值的索引位置
	cur3, cur5, cur7 := 0, 0, 0
	for i := 1; i < k; i++ {
		v1 := dp[cur3] * 3
		v2 := dp[cur5] * 5
		v3 := dp[cur7] * 7
		min := min(v1, min(v2, v3)) // 每次都获取当前所有3、5、7 *对应值最小的数 （动规）

		if min == v1 { // eg: 当前cur3 = 1, 下一个*3正好大于当前的就是cur3 = 2
			cur3++
		}
		if min == v2 {
			cur5++
		}
		if min == v3 {
			cur7++
		}

		dp[i] = min
	}

	return dp[k-1]
}

// 骑士在棋盘上的概率
// 思路：DP题，反推，以所有点第0步概率为1作为基础，第1步是棋盘上所有点能被8个方位走到的概率，最终求出第k步row,column被k-1步所有点8方位能走到的概率
func knightProbability(n int, k int, row int, column int) float64 {
	type pair struct{ i, j int }
	// 8个方向
	var dirs = []pair{{2, 1}, {1, 2}, {2, -1}, {1, -2}, {-1, 2}, {-2, 1}, {-1, -2}, {-2, -1}}

	dp := make([][][]float64, k+1) // 0-k

	for step := 0; step <= k; step++ {
		dp[step] = make([][]float64, n) // 0 - n-1

		for i := 0; i < n; i++ {
			dp[step][i] = make([]float64, n) // 0 - n-1

			for j := 0; j < n; j++ {
				if step == 0 {
					dp[step][i][j] = 1 // 第0步 概率都是1
				} else {
					for _, val := range dirs { // 第1-k步，概率都是根据8个在棋盘上可以走向他的方位的总和
						x, y := i+val.i, j+val.j                // 找到上个点的位置
						if 0 <= x && x < n && 0 <= y && y < n { // 检查可行度
							dp[step][i][j] += dp[step-1][x][y] / 8 // 可行 说明上一步有 1/8的概率走到当前这一步
						}
					}
				}
			}
		}
	}

	return dp[k][row][column]
}

// N皇后2
func totalNQueens(n int) int {

	m := map[int]int{
		1: 1,
		2: 0,
		3: 0,
		4: 2,
		5: 10,
		6: 4,
		7: 40,
		8: 92,
		9: 352,
	}

	return m[n]
}

// 好子集的数目
var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

func numberOfGoodSubsets(nums []int) (ans int) {
	const mod int = 1e9 + 7
	freq := [31]int{}
	for _, num := range nums {
		freq[num]++
	}

	f := make([]int, 1<<len(primes))
	f[0] = 1
	for i := 0; i < freq[1]; i++ { // 1的基础拼接次数
		f[0] = f[0] * 2 % mod
	}

next:
	for i := 2; i < 31; i++ {
		if freq[i] == 0 { // 不存在的数就跳过
			continue
		}

		// 检查 i 的每个质因数是否均不超过 1 个
		subset := 0
		for j, prime := range primes {
			if i%(prime*prime) == 0 { // 当前值非好值就跳过
				continue next
			}
			if i%prime == 0 { // 当前值是质数组成的好值
				subset |= 1 << j
			}
		}

		// 动态规划
		for mask := 1 << len(primes); mask > 0; mask-- {
			if mask&subset == subset {
				f[mask] = (f[mask] + f[mask^subset]*freq[i]) % mod
			}
		}
	}

	for _, v := range f[1:] {
		ans = (ans + v) % mod
	}
	return
}

// 球会落何处
// 思路：层次遍历
func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	ans := make([]int, n)

	type ball struct {
		idx int // 几号球
		x   int // x轴位置
		y   int // y轴位置
	}

	q := []ball{}
	for i, _ := range ans {
		q = append(q, ball{idx: i, x: i, y: 0})
	}

	for len(q) != 0 {
		p := q[0]
		q = q[1:]

		cur := grid[p.y][p.x] // 当前球往下的方向

		if cur+p.x < 0 || cur+p.x > n-1 { // 当遇到墙了就没了
			ans[p.idx] = -1
			continue
		}

		if grid[p.y][p.x+cur] != cur { // 当遇到不同方向也没了
			ans[p.idx] = -1
			continue
		}

		if p.y == m-1 { // 到了最后一层
			ans[p.idx] = p.x + cur
			continue
		}

		p.x = p.x + cur
		p.y = p.y + 1
		q = append(q, p)
	}

	return ans
}

// 适合打劫银行的日子
func goodDaysToRobBank(security []int, time int) (ans []int) {
	n := len(security)
	left := make([]int, n)
	right := make([]int, n)
	for i := 1; i < n; i++ {
		// 计算出 每一步左边的是递增的数
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		}
		// 每一步右边是递减的数
		if security[n-i-1] <= security[n-i] {
			right[n-i-1] = right[n-i] + 1
		}
	}

	// 当左右都>=time就等于找到结果
	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return
}

// 最接近目标价格的甜点成本  -层次遍历
func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	ans := math.MaxInt64
	n := len(toppingCosts)

	type cost struct {
		val int // 当前价格
		idx int // 当前索引
	}

	for _, baseVal := range baseCosts {

		q := []cost{{val: baseVal, idx: 0}}

		for len(q) > 0 {
			// 弹出
			p := q[0]
			q = q[1:]

			last := abs(ans - target)
			now := abs(p.val - target)

			if now < last { // 当新的差值 < 旧的差值  = 找到更接近的了
				ans = p.val
			} else if now == last && p.val < target { // 当新的差值 < 旧的差值  && 新的值 < 旧的值 = 找到同样接近中的更小的值
				ans = p.val
			} else if now == 0 { // now = 0 不用找了 能组成刚好合适的
				return p.val
			}

			if p.val > target { // 大于只会导致差距越来越大
				continue
			}

			if p.idx >= n { // 索引超了
				continue
			}

			inr := toppingCosts[p.idx] // 增量

			q = append(q, cost{val: p.val, idx: p.idx + 1})
			q = append(q, cost{val: p.val + inr, idx: p.idx + 1})
			q = append(q, cost{val: p.val + inr*2, idx: p.idx + 1})
		}
	}

	return ans
}

// 最大的以 1 为边界的正方形
func largest1BorderedSquare(grid [][]int) int {
	// 找到每个位置的上左边长
	dpup := make([][]int, 0)   // 上边长
	dpleft := make([][]int, 0) // 左边长

	for x, arr := range grid {
		curx := make([]int, len(arr))
		cury := make([]int, len(arr))
		for y, val := range arr {
			if val == 0 {
				continue
			}
			if x == 0 {
				curx[y] = 1
			} else {
				curx[y] = dpup[x-1][y] + 1
			}

			if y == 0 {
				cury[y] = 1
			} else {
				cury[y] = cury[y-1] + 1
			}
		}
		dpup = append(dpup, curx)
		dpleft = append(dpleft, cury)
	}

	maxside := 0
	for y, arr := range dpup {
		for x, val := range arr {
			yside := val
			xside := dpleft[y][x]

			// 以上左边中最小的边长能组成正方形
			curmaxside := min(xside, yside)

			// 如果这个是的边长没之前的边长大 就不用了
			if curmaxside <= maxside {
				continue
			}
			//fmt.Printf("x:%d,y:%d,xside:%d,yside:%d,curside:%d\n", x, y, xside, yside, curmaxside)

			// 以当前点为正方形右下角 找到正方形左下角和正方形右上角看是否满足
			// 依次以最大可组成边长往里面收缩 直到边长<=已有的最大边长
			for curside := curmaxside; curside > maxside; curside-- {
				l := x - curside + 1
				u := y - curside + 1
				if dpup[y][l] >= curside && dpleft[u][x] >= curside {
					// 满足的情况下算是找到边长了然后暂停
					maxside = curside
					break
				}
			}
		}
	}

	return maxside * maxside
}

func maximalSquare(matrix [][]byte) int {
	// 找到每个位置的可能边长
	dpup := make([][]int, 0) // 边长

	for x, arr := range matrix {
		curx := make([]int, len(arr))
		for y, val := range arr {
			if val == '0' {
				continue
			}

			if x == 0 {
				curx[y] = 1
			} else {
				curx[y] = dpup[x-1][y] + 1
			}
		}
		dpup = append(dpup, curx)
	}

	maxside := 0
	for y, arr := range dpup {
		for x, val := range arr {
			curmaxside := val

			// 如果这个是的边长没之前的边长大 就不用了
			if curmaxside <= maxside {
				continue
			}

			for curside := curmaxside; curside > maxside; curside-- {
				flag := true
				for i := 1; i < curside; i++ {
					if x-i < 0 || dpup[y][x-i] < curside {
						flag = false
						break
					}
				}
				if flag {
					maxside = curside
					break
				}
			}
		}
	}

	return maxside * maxside
}

// 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			f := text1[i] == text2[j]
			if f {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[m][n]
}

func minPushBox(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var sx, sy, bx, by int // 玩家、箱子的初始位置
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if grid[x][y] == 'S' {
				sx, sy = x, y
			} else if grid[x][y] == 'B' {
				bx, by = x, y
			}
		}
	}

	ok := func(x, y int) bool { // 不越界且不在墙上
		return x >= 0 && x < m && y >= 0 && y < n && grid[x][y] != '#'
	}
	d := []int{0, -1, 0, 1, 0} // y上 x左 y下 x右

	dp := make([][]int, m*n)
	for i := 0; i < m*n; i++ {
		dp[i] = make([]int, m*n)
		for j := 0; j < m*n; j++ {
			dp[i][j] = 0x3f3f3f3f
		}
	}
	dp[sx*n+sy][bx*n+by] = 0 // 初始状态的推动次数为 0
	q := [][2]int{{sx*n + sy, bx*n + by}}

	for len(q) > 0 {
		q1 := [][2]int{} // 当前队列
		for len(q) > 0 {
			s1, b1 := q[0][0], q[0][1]
			q = q[1:]
			sx1, sy1, bx1, by1 := s1/n, s1%n, b1/n, b1%n
			if grid[bx1][by1] == 'T' { // 已被推到目标处
				return dp[s1][b1]
			}
			for i := 0; i < 4; i++ { // 玩家向四个方向移动到另一个状态
				sx2, sy2 := sx1+d[i], sy1+d[i+1]
				s2 := sx2*n + sy2
				if !ok(sx2, sy2) { // 玩家位置不合法
					continue
				}
				if bx1 == sx2 && by1 == sy2 { // 如果人的位置和箱子位置重叠 说明人可以推箱子
					bx2, by2 := bx1+d[i], by1+d[i+1] // 向人走的方向推动箱子
					b2 := bx2*n + by2
					if !ok(bx2, by2) { // 查看箱子位置不合法
						continue
					}
					if dp[s2][b2] <= dp[s1][b1]+1 { // 状态已访问
						continue
					}
					dp[s2][b2] = dp[s1][b1] + 1
					q1 = append(q1, [2]int{s2, b2})
				} else {
					if dp[s2][b1] <= dp[s1][b1] { // 状态已访问
						continue
					}
					dp[s2][b1] = dp[s1][b1]
					q = append(q, [2]int{s2, b1})
				}
			}
		}
		q = q1
	}
	return -1
}

// 1883. 准时抵达会议现场的最小跳过休息次数
func minSkips(dist []int, speed int, hoursBefore int) int {
	// 可忽略误差
	const EPS = 1e-7

	n := len(dist)
	f := make([][]float64, n+1)
	for i := range f {
		f[i] = make([]float64, n+1)
		for j := range f[i] {
			f[i][j] = math.Inf(1)
		}
	}
	// 动态规划
	f[0][0] = 0.0
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if j != i {
				f[i][j] = math.Min(f[i][j], math.Ceil(f[i-1][j]+float64(dist[i-1])/float64(speed)-EPS))
			}
			if j != 0 {
				f[i][j] = math.Min(f[i][j], f[i-1][j-1]+float64(dist[i-1])/float64(speed))
			}
		}
	}

	for j := 0; j <= n; j++ {
		if f[n][j] < float64(hoursBefore)+EPS {
			return j
		}
	}
	return -1
}

// 377. 组合总和 Ⅳ
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// 2786. 访问数组中的位置使分数最大
func maxScore(nums []int, x int) int64 {
	res := nums[0]
	last := []int{math.MinInt32, math.MinInt32}
	last[nums[0]%2] = int(nums[0])

	for i := 1; i < len(nums); i++ {
		lave := nums[i] % 2
		cur := max(last[lave]+nums[i], last[1-lave]-x+nums[i])
		res = max(res, cur)
		last[lave] = cur
	}

	return int64(res)
}

// 2901. 最长相邻不相等子序列 II
func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(groups)
	dp := make([]int, n)
	prev := make([]int, n)
	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}
	maxIndex := 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if checkGetWordsInLongestSubsequence(words[i], words[j]) && dp[j]+1 > dp[i] && groups[i] != groups[j] {
				dp[i] = dp[j] + 1
				prev[i] = j
			}
		}
		if dp[i] > dp[maxIndex] {
			maxIndex = i
		}
	}

	var res []string
	for i := maxIndex; i >= 0; i = prev[i] {
		res = append(res, words[i])
	}
	reverseGetWordsInLongestSubsequence(res)
	return res
}

func checkGetWordsInLongestSubsequence(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
			if diff > 1 {
				return false
			}
		}
	}
	return diff == 1
}

func reverseGetWordsInLongestSubsequence(arr []string) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
