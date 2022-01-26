package program

import (
	"math"
	"sort"
)

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

// 盛最多水的容器--------------------------------------------------------------------------------------------------------
// 思路：双指针 边界最高为基准即可找到最多的水容器
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0

	for l < r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}

// 三数之和
// 思路：排序，然后定最左位的点，用双指针找其右边的两个与其组合为0的
func threeSum(nums []int) [][]int {
	// 排序
	sort.Ints(nums)
	res := make([][]int, 0)
	n := len(nums)

	for idx, v := range nums {
		// 第一个点>=0 或 后面没有2个数了 就等于找到头了
		if v > 0 || idx > n-3 {
			break
		}

		// 相同的数跳过
		if idx-1 >= 0 && nums[idx] == nums[idx-1] {
			continue
		}

		// 建立双指针
		l, r := idx+1, n-1
		for l < r {

			// 跳过相同左数据
			if nums[l] == nums[l-1] && l-1 != idx {
				l++
				continue
			}

			// 跳过相同右数据
			if r+1 < n && nums[r] == nums[r+1] {
				r--
				continue
			}

			// 合并项大于0移动右指针让其缩小
			if nums[l]+nums[r]+v > 0 {
				r--
				// 合并项小于0移动左指针让其变大
			} else if nums[l]+nums[r]+v < 0 {
				l++
			} else {
				// 找到结果了，可以随便移动指针
				res = append(res, []int{v, nums[l], nums[r]})
				l++
			}
		}
	}

	return res
}

// 最接近的三数和
// 思路：核心思路与三数之和相同，利用双指针特性 在固定第一个数的情况下减少2、3数的匹配次数，每次提炼出来3个数合并 看是否最接近目标
func threeSumClosest(nums []int, target int) int {
	// 排序
	sort.Ints(nums)
	res := -9999999
	n := len(nums)

	for idx, v := range nums {
		// 后面没有2个数了 就等于找到头了
		if idx > n-3 {
			break
		}

		// 相同的数跳过
		if idx-1 >= 0 && nums[idx] == nums[idx-1] {
			continue
		}

		// 建立双指针
		l, r := idx+1, n-1
		for l < r {

			// 跳过相同左数据
			if nums[l] == nums[l-1] && l-1 != idx {
				l++
				continue
			}

			// 跳过相同右数据
			if r+1 < n && nums[r] == nums[r+1] {
				r--
				continue
			}

			x := nums[l] + nums[r] + v

			// 合并项大于target移动右指针让其缩小
			if x > target {
				r--
				// 合并项小于target移动左指针让其变大
			} else if x < target {
				l++
			} else {
				return target
			}
			res = closeNum(target, res, x)
		}
	}

	return res
}

func closeNum(target, x, y int) int {
	a, b := int(math.Abs(float64(x-target))), int(math.Abs(float64(y-target)))
	if a >= b {
		return y
	}
	return x
}

// 删除回文子序列
// 思路：只有a,b 删除 是回文的子序列，如果本身就是回文只需要1次，如果本身不是回文先删除a再删除b只需要2次
func removePalindromeSub(s string) int {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return 2
		}
	}
	return 1
}