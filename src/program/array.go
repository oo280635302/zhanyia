package program

import (
	"container/heap"
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

// 至少是其他数字两倍的最大数  至少1个数并且大于0
// 思路：贪心，找到最大数和第二大的数，如果最大数是第二大数的2倍+就说明比数组所有其他数都大2倍+
func dominantIndex(nums []int) int {
	firstMaxNum := nums[0]
	secondMaxNum := 0
	firstIdx := 0

	for idx, val := range nums[1:] {

		if val > firstMaxNum { // 如果值比当前最大还大就成为最大，并把原最大值给老二
			secondMaxNum = firstMaxNum
			firstMaxNum = val
			firstIdx = idx + 1
		} else if val > secondMaxNum { // 第二大 < val < 最大，那他就是第二大
			secondMaxNum = val
		}

		fmt.Println(firstMaxNum, secondMaxNum)
	}

	if firstMaxNum >= secondMaxNum*2 {
		return firstIdx
	}

	return -1
}

// 查找和最小的 K 对数字
// 逻辑：堆 每次推出最小的同时，推入比他大的组合进行堆排序，只要找够k对位置
// 因为正序排序，对于已经被推出去的组合比他大的数要么i+1，要么j+1
func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ { // 先将 第一个数组的 推进去排序
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair) // 堆在推的时候排序在推出
		i, j := p.i, p.j         // 当前最大ij
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1}) // 因为当前i+1都有，所有只需要j+1的数据即可  其他数据要么都是i，j+1的 要么已经被推出去了
		}
	}
	return
}

type pair struct{ i, j int }
type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}                                // 对比 和
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v } // 逆序，最大的是最后一个 小顶堆
