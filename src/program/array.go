package program

import (
	"fmt"
	"math"
)

// 提莫攻击的持续时长
// 思路：每次都保存上次攻击的结束时间 如果当前攻击的时间点小于上次攻击结算时间点 持续时长只增加当前时间+持续时间-上次结算时间点即可
func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 0
	last := 0

	for _, v := range timeSeries {
		if v >= last {
			res += duration
		} else {
			res += duration + v - last
		}
		last = v + duration
	}
	return res
}

// 寻找两个正序数组的中位数
// 思路：双指针，去掉两个数组的最小直到找到中间数为止
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i, j, k := 0, 0, 0
	last, cur := 0, 0
	m, n := len(nums1), len(nums2)
	mid := (m + n) / 2 // 中间数，如果是偶数应该后一个

	for k <= mid {
		last = cur
		// 当i,j都没到达数组的末尾时比较
		if i < m && j < n {
			if nums1[i] < nums2[j] {
				cur = nums1[i]
				i++
			} else {
				cur = nums2[j]
				j++
			}
			// i<m说明j到达尾巴了，移动i
		} else if i < m {
			cur = nums1[i]
			i++
			// 说i到达尾巴了，移动j
		} else {
			cur = nums2[j]
			j++
		}
		k++
		fmt.Println(last, cur, k, mid)
	}

	if (m+n)%2 == 1 {
		return float64(cur)
	}

	return float64(last+cur) / 2
}

// 递增的三元子序列
// 思路：贪心 让1<2的同时 1,2尽可能小 然后找到3  "贪"字诀
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	st1, st2 := nums[0], math.MaxInt64 // 1号从头开始， 2号先设置为无限大(以方便找到比1大的数)
	for _, v := range nums[1:] {
		if v > st2 { // 当前值 > 2号 说明找到了
			return true
		} else if v > st1 { // 2号 > 当前值 > 1号 说明他是当前的最小2号
			st2 = v
		} else { // 当前值 < 1号 说明他是当前最小1号
			st1 = v
		}
	}
	return false
}
