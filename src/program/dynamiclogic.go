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
