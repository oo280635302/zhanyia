package program

// 双指针有关的算法问题 -- LeetCode

// 长度最小的子数组------------------------------------------------------------------------------------------------------
// 思路:双指针 一个指向头,一个指向尾,尾移动到和>s 后 移动头缩小两者间间距,移动不了了 再移动头 8 ms/3.8mb
func MinSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	head, tail := 0, 0
	sum, lenth := nums[0], 0
	for head <= len(nums)-1 {
		//fmt.Println(head,tail,sum)
		// 和>s 记录长度
		if sum >= s {
			if lenth == 0 {
				lenth = tail - head + 1
			}
			lenth = min(lenth, tail-head+1)
			// 移动头指针
			sum -= nums[head]
			head++
			continue
		}

		if len(nums)-1 > tail {
			sum += nums[tail+1]
			tail++
		} else {
			break
		}
	}
	return lenth
}

// 两个数组的交集 II--------------------------------------------------------------------------------------------------
// 思路: 用一个 map 存放小的那方的数组 每有一个重复的+1 然后用大的一方去匹配小的一方 有就减少 直到0  4ms/3.2mb
// 额外思路: 双指针 排序后 双指针移动 捕捉相同的数
func intersect(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	if m < n {
		nums1, nums2 = nums2, nums1
	}

	map1 := make(map[int]int)
	for _, v := range nums1 {
		map1[v]++
	}

	res := make([]int, 0)
	for _, v := range nums2 {
		if r, ok := map1[v]; ok && r > 0 {
			map1[v]--
			res = append(res, v)
		}
	}

	return res
}

// 两数之和2 有序数组-------------------------------------------------------------------------------------------------
// 思路：遍历 map保存已经找到的数+对应的位置  8ms/3.1mb
func twoSum(numbers []int, target int) []int {
	checkMap := make(map[int]int)
	res := make([]int, 0)
	for k, v := range numbers {

		if l, ok := checkMap[target-v]; ok {
			res = append(res, l+1)
			res = append(res, k+1)
			break
		}
		checkMap[v] = k
	}

	return res
}

// 优解:双指针 避免了内存浪费 4ms/3mb
func TwoSumBetterSlv(numbers []int, target int) []int {
	n := len(numbers)
	left, right := 0, n-1
	for left < right {
		s := numbers[left] + numbers[right]
		if s == target {
			return []int{left + 1, right + 1}
		} else if s < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// 旋转数组的最小数字 + 寻找旋转排序数组中的最小值 + 寻找旋转排序数组中的最小值2---3道题共同解法
// 思路：双指针 左右共同寻找最小值(左>右 就= 右最小),需要注意的是初始最小值设为number[0] 为没旋转的数组最小的值 0ms/3.1mb
func minArray(numbers []int) int {
	n := len(numbers)
	// 数组为0直接返回
	if n == 0 {
		return 0
	}

	minNum := numbers[0]
	// 左右两边共同开工,寻找最小值
	l, r := 0, n-1
	for l < r {
		if numbers[l] > numbers[l+1] || numbers[r-1] > numbers[r] {
			minNum = min(minNum, numbers[l+1])
			minNum = min(minNum, numbers[r])
			break
		}
		l++
		r--
	}
	return minNum
}

// 优化解: 两份查找 中间切割 如果中间值<右边 = min在左边的区间里 将m赋值给r 同理将m+1赋给l ,如果相等r-- 破坏平衡 4ms/3.1mb
func minArrayBetterSlv(numbers []int) int {
	l := 0
	r := len(numbers) - 1
	for l < r {
		m := l + (r-l)/2
		if numbers[m] < numbers[r] {
			r = m
		} else if numbers[m] > numbers[r] {
			l = m + 1
		} else {
			r--
		}
	}
	return numbers[l]
}
