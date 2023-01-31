package program

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strconv"
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

// 最小绝对差
// 思路：排序后遍历，绝对差最小的一定是相邻的，差值相等就追加/小于就重置返回
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	res := make([][]int, 0)

	min := arr[1] - arr[0]

	for idx, v := range arr[1:] {
		cur := v - arr[idx]
		if cur == min {
			res = append(res, []int{arr[idx], v})
		} else if cur < min {
			res = append([][]int{}, []int{arr[idx], v})
		}
	}

	return res
}

// 最小时间差
// 思路:转为数字，再用最小绝对差的方式找最小，需要注意的是跨天的哪一点
func findMinDifference(timePoints []string) int {
	times := make([]int, 0)
	for _, v := range timePoints {
		h, _ := strconv.Atoi(v[0:2])
		m, _ := strconv.Atoi(v[3:5])
		times = append(times, h*60+m)
	}
	sort.Ints(times)
	times = append(times, times[0]+1440) // 跨天那一天
	ans := times[1] - times[0]
	for idx, v := range times[1:] {
		if v-times[idx] < ans {
			ans = v - times[idx]
		}
	}
	return ans
}

// 存在重复元素 II
// 思路：哈希+滑动窗口 用一个map来 保存当前值的前k步所走过的值
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]bool)
	for idx, val := range nums {
		if idx > k { // 如果idx>k，就删除m中第k+1个前的那个值 已确保m保存的是前k个数
			delete(m, nums[idx-k-1])
		}
		if m[val] == true { // 找到了
			return true
		}
		m[val] = true // 没找到把当前值存进去，让下个数来找
	}
	return false
}

// 下一个排列   找到比当前排列正好大一点的新排列
func nextPermutation(nums []int32) {
	// 逆序找到 左<右的位置
	n := len(nums)
	cur := -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			cur = i //cur的右边是降序的
			break
		}
	}
	// 如果没有说明没有比当前排列更大的排列了
	if cur == -1 {
		reserveInt32(nums)
		return
	}

	// 找到右边比cur大的位数最低的数
	for i := n - 1; i > cur; i-- {
		if nums[i] > nums[cur] {
			nums[i], nums[cur] = nums[cur], nums[i] // 先交换，因为 nums[i-1] < nums[cur] < nums[i]交换后 cur 右边也可以保证降序
			reserveInt32(nums[cur+1:])              // 交换后为了保证是正好大一些的排列，将右边降序排列的大数组 翻转 升序排列的小数组 就是正好大一些
			break
		}
	}

}
func reserveInt32(nums []int32) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 跳跃游戏 IV
func minJumps(arr []int) int {
	if len(arr) == 1 {
		return 0
	}
	type jumpStruct struct {
		i    int // 在arr里面的索引
		step int // 第几步
	}
	n := len(arr)
	graph := make(map[int][]int) // 用来保存相同数的索引
	isDead := make(map[int]bool) // 标记1个索引是否被访问过了
	isDead[0] = true
	for idx, v := range arr {
		graph[v] = append(graph[v], idx)
	}

	path := []jumpStruct{{0, 0}}

	for {
		// 取出数组的第一个
		cur := path[0]
		path = path[1:]

		// 如果它就是最后一个数就结束了
		if cur.i == n-1 {
			return cur.step
		}

		// 不是，找到他能达到的所有位置

		// 纯优化,可以去掉的代码：如果i+1=n-1 说明下一步一定到
		if cur.i+1 == n-1 {
			return cur.step + 1
		}

		// 前一位
		if cur.i+1 < n && !isDead[cur.i+1] {
			path = append(path, jumpStruct{i: cur.i + 1, step: cur.step + 1})
			isDead[cur.i+1] = true
		}

		// 后一位
		if cur.i-1 > 0 && !isDead[cur.i-1] {
			path = append(path, jumpStruct{i: cur.i - 1, step: cur.step + 1})
			isDead[cur.i-1] = true
		}

		// 相同数
		for _, v := range graph[arr[cur.i]] {
			// 纯优化,可以去掉的代码：如果v=n-1说明下一步一定到
			if v == n-1 {
				return cur.step + 1
			}
			if !isDead[v] {
				path = append(path, jumpStruct{i: v, step: cur.step + 1})
				isDead[v] = true
			}
		}

		delete(graph, arr[cur.i]) // 把已经走过的相同数清理掉，减少下次遇到了遍历相同数
	}
}

// 学生分数的最小差值
// 思路： 排序后，每k个连续的学生的最大最小值的差值
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	res := 1000000

	for idx, val := range nums[:len(nums)+1-k] { // 左指针只需要滑动到 1.右往左数第k-1个 len-k+1 2.索引-1  len-k 3.左开右闭 len-k+1
		res = min(res, nums[idx+k-1]-val)
	}

	return res
}

// 有序数组中的单一元素
// 思路：二分查找 找到中位数，判断中位数补位的相邻数是否相等，相等说明前一半是偶数 就找后面的，不相等说明单一数在前一半就找前面   x^1 找到补位数 eg: 11^1=10  10^1=11
func singleNonDuplicate(nums []int) int {
	s, e := 0, len(nums)-1

	for s < e {
		mid := (s + e) / 2

		if nums[mid] == nums[mid^1] {
			s = mid + 1
		} else {
			e = mid
		}
	}

	return nums[s]
}

// 矩阵中的幸运数
func luckyNumbers(matrix [][]int) []int {
	ans := []int{}

	minRow := make(map[int]int)
	maxCol := make(map[int]int)

	for idx1, row := range matrix {
		for idx2, val := range row {

			if val < minRow[idx1] || minRow[idx1] == 0 {
				minRow[idx1] = val
			}
			if val > maxCol[idx2] {
				maxCol[idx2] = val
			}
		}
	}

	for _, v1 := range minRow {
		for _, v2 := range maxCol {
			if v1 == v2 {
				ans = append(ans, v1)
			}
		}
	}

	return ans
}

// 煎饼排序
// 思路：从当前数组找到最大的数翻转到前面再翻转到最后去，然后把最后的数排除找最大的数重复翻转再排除
func pancakeSort(arr []int) []int {
	ans := make([]int, 0)

	for i := len(arr); i > 1; i-- {

		idx := 0 // 找到当前最大的数
		for cur, v := range arr {
			if v == i {
				idx = cur
			}
		}

		if idx == i-1 {
			continue
		}

		reverseInts(arr[:idx+1]) // 先翻转到前面来
		reverseInts(arr[:i])     // 再翻转到有序的地方去

		ans = append(ans, idx+1, i)
	}

	return ans
}
func reverseInts(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
}

// 增量元素之间的最大差值
func maximumDifference(nums []int) int {
	ans := -1

	min := nums[0]

	for _, v := range nums {
		if v > min {
			ans = max(ans, v-min) // 找到当前数大于最小值 就去与结果比较
		} else {
			min = v // 如果没找到，那当前数就是最小值
		}
	}

	return ans
}

// 得分最高的最小轮调
// 思路：需要注意针对这类问题用 前缀和 解决的切入点， 该问题是我们能根据移动规律来找出每次移动的涨掉分规律
func bestRotation(nums []int) int {
	n := len(nums)
	diff := make([]int, n+1) // 用以保存每移动一步的所造成的涨/掉分
	for i := 0; i < n; i++ {
		if num := nums[i]; i >= num { // 针对已经处在 [i,n-1] 之间的数
			diff[0]++       // 不移动就会得分
			diff[i-num+1]-- // 移动到 i以下就会掉分
			diff[i+1]++     // 移动到 i以上就上涨分
		} else { // 针对处于[0,i]之间的数
			diff[i+1]++       // 移动到i以上就涨分
			diff[i-num+n+1]-- // 移动到i以下就掉分
		}
	}

	ans := 0
	max := 0

	for idx, v := range diff {
		if v+max > max {
			max = v + max
			ans = idx
		}
	}

	return ans
}

// 向数组中追加 K 个整数
func minimalKSum(nums []int, k int) int64 {
	// 所有 k+已出现比k小的nums数量 的和 - 已出现比k小的nums数量
	sort.Ints(nums)

	diff := 0
	last := -1
	for _, v := range nums {
		if v <= k {
			if last != v {
				k++
				diff += v
				last = v
			}
		} else {
			break
		}
	}

	return int64((1+k)*k/2 - diff)
}

// 图片平滑器
func imageSmoother(img [][]int) [][]int {
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := range ans[i] {
			sum, num := 0, 0
			// 找每个格子周边所有格子的和/格子数
			for _, row := range img[max(i-1, 0):min(i+2, m)] {
				for _, v := range row[max(j-1, 0):min(j+2, n)] {
					sum += v
					num++
				}
			}
			ans[i][j] = sum / num
		}
	}
	return ans
}

// 考试的最大困扰度
func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveAnswersByBit(answerKey, k, 'T'),
		maxConsecutiveAnswersByBit(answerKey, k, 'F'))
}
func maxConsecutiveAnswersByBit(answerKey string, k int, b byte) int {
	l, x, ans := 0, 0, 0

	for r := range answerKey {
		if answerKey[r] != b { //遇到非b将x++
			x++
		}

		for x > k { // 当x超过限制时，移动l，来减少x的数量以匹配k的数量
			if answerKey[l] != b {
				x--
			}
			l++
		}
		ans = max(ans, r-l+1)
	}

	return ans
}

// 二倍数对数组
func canReorderDoubled(arr []int) bool {
	cnt := make(map[int]int, len(arr))
	for _, x := range arr {
		cnt[x]++
	}
	if cnt[0]%2 == 1 {
		return false
	}

	vals := make([]int, 0, len(cnt))
	for x := range cnt {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) })

	for _, x := range vals {
		if cnt[2*x] < cnt[x] { // 无法找到足够的 2x 与 x 配对
			return false
		}
		cnt[2*x] -= cnt[x]
	}
	return true
}

// 掉落的方块
func fallingSquares(positions [][]int) []int {
	ans := make([]int, len(positions))

	// 计算每个方块掉落能组成的当前高度
	for idx, p := range positions {
		left, right := p[0], p[0]+p[1]-1 // 每个方块的左边界、右边界

		cur := p[1]
		ans[idx] = p[1] // 默认高度

		for i, p2 := range positions[:idx] { // 遍历已落下的方块 看是否有相交
			left2, right2 := p2[0], p2[0]+p2[1]-1
			if right2 >= left && right >= left2 { // 相交
				ans[idx] = max(ans[idx], cur+ans[i])
			}
		}
	}

	for i := 1; i < len(ans); i++ {
		ans[i] = max(ans[i], ans[i-1])
	}

	return ans
}

// 单词距离
func findClosest(words []string, word1 string, word2 string) int {
	// 因为题目提示需要复用， 拿map来保存索引位置
	m := make(map[string][]int)
	for idx, v := range words {
		m[v] = append(m[v], idx)
	}

	// 两个文字的索引的位置 升序的
	arr1 := m[word1]
	arr2 := m[word2]
	if len(arr1) == 0 || len(arr2) == 0 {
		return 0
	}

	ans := math.MaxInt64
	idx1 := 0
	idx2 := 0

	for {
		// 双指针找最小的差值
		ans = min(abs(arr1[idx1]-arr2[idx2]), ans)
		if idx1 == len(arr1)-1 && idx2 == len(arr2)-1 {
			break
		}

		// 除了最大索引以外 都应该交替移动 比如 arr1[idx1]<arr2[idx2] 如果再移动idx2只会让差距变大
		if arr1[idx1] < arr2[idx2] && idx1 < len(arr1)-1 {
			idx1++
		} else if idx2 < len(arr2)-1 && idx2 < len(arr2)-1 {
			idx2++
		} else if idx1 < len(arr1)-1 && idx2 == len(arr2)-1 {
			idx1++
		} else if idx2 < len(arr2)-1 && idx1 == len(arr1)-1 {
			idx2++
		}

	}

	return ans
}

// 独一无二的出现次数 Unique Number of Occurrences
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	numM := make(map[int]bool)
	for _, val := range m {
		if numM[val] {
			return false
		}
		numM[val] = true
	}
	return true
}

// 至少在两个数组中出现的值
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	ans := []int{}
	m := map[int]bool{}
	m1 := map[int]bool{}
	for _, v := range nums1 {
		m1[v] = true
	}

	m2 := map[int]bool{}
	for _, v := range nums2 {
		if m1[v] {
			m[v] = true
		}
		m2[v] = true
	}

	for _, v := range nums3 {
		if m1[v] || m2[v] {
			m[v] = true
		}
	}

	for k := range m {
		ans = append(ans, k)
	}

	return ans
}

// 统计不开心的朋友
func unhappyFriends(n int, preferences [][]int, pairs [][]int) (ans int) {
	order := make([][]int, n)
	for i, preference := range preferences {
		order[i] = make([]int, n)
		for j, p := range preference {
			order[i][p] = j
		}
	}
	match := make([]int, n)
	for _, p := range pairs {
		match[p[0]] = p[1]
		match[p[1]] = p[0]
	}

	for x, y := range match {
		index := order[x][y]
		for _, u := range preferences[x][:index] {
			v := match[u]
			if order[u][x] < order[u][v] {
				ans++
				break
			}
		}
	}
	return
}

// 积压订单中的订单总数
func getNumberOfBacklogOrders(orders [][]int) int {
	// 以后优化点: 积压订单用堆 减少插入的时间复杂度
	buy0 := [][]int{}  //[]int{price, num}
	sell1 := [][]int{} //[]int{price, num}

	for _, order := range orders {
		price := order[0]
		num := order[1]

		switch order[2] {
		case 0:
			for len(sell1) > 0 {
				p := sell1[0]
				if p[0] > price || num == 0 { // 销售价格大于购买价格 或者 当前购买订单没货了 停止
					break
				}
				cnt := min(p[1], num)
				p[1] -= cnt
				num -= cnt

				if p[1] == 0 {
					sell1 = sell1[1:]
				}
			}

			// 还剩订单就追加到 购买积压订单里面
			if num > 0 {
				idx := len(buy0)
				for k, v := range buy0 {
					if price < v[0] {
						idx = k
						break
					}
				}
				buy0 = append(buy0, []int{})
				copy(buy0[idx+1:], buy0[idx:])
				buy0[idx] = []int{price, num}
			}
		case 1:
			for len(buy0) > 0 {
				p := buy0[len(buy0)-1]
				if price > p[0] || num == 0 { // 销售价格大于购买价格 或者 当前购买订单没货了 停止
					break
				}
				cnt := min(p[1], num)
				p[1] -= cnt
				num -= cnt

				if p[1] == 0 {
					buy0 = buy0[:len(buy0)-1]
				}
			}

			// 还剩订单就追加到 购买积压订单里面
			if num > 0 {
				idx := len(sell1)
				for k, v := range sell1 {
					if price < v[0] {
						idx = k
						break
					}
				}
				sell1 = append(sell1, []int{})
				copy(sell1[idx+1:], sell1[idx:])
				sell1[idx] = []int{price, num}
			}
		}

	}

	ans := 0
	for _, v := range buy0 {
		ans += v[1]
		ans %= 1e9 + 7
	}
	for _, v := range sell1 {
		ans += v[1]
		ans %= 1e9 + 7
	}
	return ans
}

// 检查句子中的数字是否递增
func areNumbersAscending(s string) bool {
	last := 0
	for i := 0; i < len(s); i++ {
		cur := 0

		if s[i] < '0' || s[i] > '9' {
			continue
		}

		for {
			num := int(s[i]) - 48
			if num < 0 || num > 9 {
				break
			}
			cur = cur*10 + num
			i++
		}
		//fmt.Println(last, cur)
		if cur <= last {
			return false
		}
		last = cur
	}

	return true
}

// 还原排列的最少操作步数
func reinitializePermutation(n int) int {
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}
	ans := 1

	for {
		arr := make([]int, n)
		copy(arr, perm)

		for i := range arr {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}
		fmt.Println(arr)
		flag := false
		for k, v := range arr {
			if k != v {
				flag = true
				break
			}
		}

		if !flag {
			return ans
		}

		// 不符合继续变化
		perm = arr
		ans++
	}
}

// 统计一个数组中好对子的数目
func countNicePairs(nums []int) (ans int) {
	cnt := map[int]int{}
	for _, num := range nums {
		rev := 0
		for x := num; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		fmt.Println(num, rev, num-rev, cnt[num-rev])
		ans += cnt[num-rev]
		cnt[num-rev]++
	}
	return ans % (1e9 + 7)
}
