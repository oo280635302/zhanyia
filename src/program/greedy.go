package program

import (
	"fmt"
	"math"
	"sort"
)

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

// 石子游戏 IX
// 思路：把012的关系找清楚
func stoneGameIX(stones []int) bool {
	cnt0, cnt1, cnt2 := 0, 0, 0
	for _, v := range stones {
		if v%3 == 1 {
			cnt1++
		} else if v%3 == 2 {
			cnt2++
		} else {
			cnt0++
		}
	}

	/*
		特殊：0为奇数可以反转胜负，除了石子下完blob赢的情况。
		1和2组成的两种情况：
			1 12121212121...
			2 21212121212...

		Alice作为先手可以任意选择1、2开始。
		当0偶数：
			情况1:
				有1并且1<=2 alice赢
			情况2：
				有2并且2<=1 alice赢
			综上：
				只要既有1又有2，alice就可以跟据选择不同的路线获胜
				1 > 0 && 2 > 0
		当0奇数：
			0奇数可以反转胜负
			情况1：
				1比2多3个及以上  剩余为1101的情况 alice赢
			情况2：
				2比1多3个及以上  剩余为2202的情况 alice赢
			综上：
				alice就可以跟据选择不同的路线获胜
				存在情况1和情况2任意一种alice都可以赢
	*/

	if cnt0%2 == 0 {
		return cnt1 >= 1 && cnt2 >= 1
	}

	return cnt1-cnt2 > 2 || cnt2-cnt1 > 2
}

// 移除石子的最大得分
func maximumScore(a, b, c int) int {
	sum := a + b + c
	maxVal := max(max(a, b), c)
	// 假设 a<=b<=c
	// a+b <= c 匹配数量：a+b = abc和 - 最大的c
	if sum < maxVal*2 {
		return sum - maxVal
		// a+b > c 匹配数量： a+b+c / 2
	} else {
		return sum / 2
	}
}

// 转换字符串的最少操作次数
func minimumMoves(s string) int {
	ans := 0
	l := 0
	// 贪心 遇到X就覆盖 并跳过连续的3个
	for l < len(s) {
		if s[l] == 'X' {
			l += 3
			ans += 1
		} else {
			l++
		}
	}

	return ans
}

// 有界数组中指定下标处的最大值
func maxValue(n, index, maxSum int) int {
	left := index          // 到左边界的距离
	right := n - index - 1 // 到右边界的距离
	if left > right {      // 置换左右，使左边<=右边
		left, right = right, left
	}

	upper := ((left+1)*(left+1)-3*(left+1))/2 + left + 1 + (left + 1) + ((left+1)*(left+1)-3*(left+1))/2 + right + 1
	if upper >= maxSum {
		a := 1.0
		b := -2.0
		c := float64(left + right + 2 - maxSum)
		return int((-b + math.Sqrt(b*b-4*a*c)) / (2 * a))
	}

	upper = (2*(right+1)-left-1)*left/2 + (right + 1) + ((right+1)*(right+1)-3*(right+1))/2 + right + 1
	if upper >= maxSum {
		a := 1.0 / 2
		b := float64(left) + 1 - 3.0/2
		c := float64(right + 1 + (-left-1)*left/2.0 - maxSum)
		return int((-b + math.Sqrt(b*b-4*a*c)) / (2 * a))
	} else {
		a := float64(left + right + 1)
		b := float64(-left*left-left-right*right-right)/2 - float64(maxSum)
		return int(-b / a)
	}
}

// 灌溉花园的最少水龙头数目
func minTaps(n int, ranges []int) int {
	// 能被最广覆盖就能减少水龙的数量 去掉不必要的能被其余水龙头覆盖的水龙头
	land := make([]int, n)
	for i, v := range ranges {
		l := max(i-v, 0)
		r := min(i+v, n)
		for l < r {
			// 当前区域 最广水龙头
			land[l] = max(land[l], r)
			l++
		}
	}

	ans := 0
	cur := 0
	for cur < n {
		// 没被任何水龙头覆盖
		if land[cur] == 0 {
			return -1
		}
		cur = land[cur]
		ans++
	}

	return ans
}

// 距离相等的条形码
func rearrangeBarcodes(barcodes []int) []int {
	// 贪心：让最大出现频率的词间隔排布 先排偶数列再排奇数列
	res := make([]int, len(barcodes))
	h := map[int]int{}
	for _, v := range barcodes {
		h[v]++
	}
	arr := []int{}
	for k := range h {
		arr = append(arr, k)
	}
	sort.Slice(arr, func(i, j int) bool {
		return h[arr[i]] > h[arr[j]]
	})
	ptr := 0
	// 先排布偶列
	for idx := range res {
		if idx%2 == 0 {
			for {
				num := arr[ptr]
				if h[num] > 0 {
					res[idx] = num
					h[num]--
					break
				} else {
					ptr++
				}
			}
		}
	}
	// 在排布奇列
	for idx := range res {
		if idx%2 == 1 {
			for {
				num := arr[ptr]
				if h[num] > 0 {
					res[idx] = num
					h[num]--
					break
				} else {
					ptr++
				}
			}
		}
	}

	return res
}

// 受标签影响的最大值
func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return values[idx[i]] > values[idx[j]]
	})

	fmt.Println(values)
	fmt.Println(labels)
	useLabel := make(map[int]int)
	useValueCnt := 0
	res := 0
	for i := 0; i < len(idx); i++ {
		// 已经使用完了
		if useValueCnt >= numWanted {
			break
		}
		label := labels[idx[i]]
		if useLabel[label] >= useLimit {
			continue
		}
		res += values[idx[i]]
		useValueCnt++
		useLabel[label]++
	}
	return res
}
