package program

// 最大化股票利润--------------------------------------------------------------------------------------------------------
// 思路：贪心，只取上下坡利润里面上坡的点
func maxProfitGreedy(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	r := 0
	l := prices[0]
	for i := 1; i < len(prices); i++ {
		n := prices[i]
		if n-l > 0 {
			r += n - l
		}
		l = n
	}
	return r
}

// s是否是t的顺序子集合 [abc,axxbeecxx]---------------------------------------------------------------------------------
// 思路：贪心+双指针，循环移动t，找到s[l] = t[r]的话s就往移动一位，直到s，t其中一个走完，如果s走完了就符合
func isSubsequence(s string, t string) bool {
	l, r := 0, 0

	for l < len(s) && r < len(t) {
		if s[l] == t[r] {
			l++
		}
		r++
	}

	return len(s) == l
}

// 花盆问题，花盆左右不能有花盆，找能加入当前花盆数量----------------------------------------------------------------------
// 思路：贪心，每次插入判断左右是否有花盆同时自身是否有花盆的情况，需要注意左右边界等特殊情况
func canPlaceFlowers(flowerbed []int, n int) bool {
	can := 0
	for i, f := range flowerbed {
		if f == 1 {
			continue
		}
		if ((i-1 >= 0 && flowerbed[i-1] == 0) || i-1 < 0) && ((i+1 <= len(flowerbed)-1 && flowerbed[i+1] == 0) || i+1 > len(flowerbed)-1) {
			can++
			flowerbed[i] = 1
		}
	}
	return can >= n
}
