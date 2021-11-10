package program

import (
	"sort"
)

// 两个字符串是否是字母异位----------------------------------------------------------------------------------------------
// 思路：排序+比较   other哈希(但只能只能针对有限个比较)
func isAnagram(s, t string) bool {
	si, ti := []byte(s), []byte(t)
	sort.Slice(si, func(i, j int) bool { return si[i] < si[j] })
	sort.Slice(ti, func(i, j int) bool { return ti[i] < ti[j] })
	return string(si) == string(ti)
}

// 2数组交集 -----------------------------------------------------------------------------------------------------------
// 思路：排序+双指针  排序后移动双指针那个小移动那个直到一方到达终点，相同且与已有返回不一样时加入
func intersection(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool { return nums1[i] < nums1[j] })
	sort.Slice(nums2, func(i, j int) bool { return nums2[i] < nums2[j] })

	res := make([]int, 0)
	l, r := 0, 0
	for l < len(nums1) && r < len(nums2) {

		if nums1[l] <= nums2[r] {
			if len(res) == 0 || (len(res) != 0 && res[len(res)-1] != nums2[r] && nums1[l] == nums2[r]) {
				res = append(res, nums2[r])
				r++
				l++
				continue
			}
			l++
			continue
		}
		r++

	}
	return res
}

// 数组中只有1个数存在1个，其余都存在多个，找出这个数----------------------------------------------------------------------
// 思路：排序 判断拐角时数组是否之前数的个数，是1就=找到了，   其余优秀思路：参考 singleNumberByBit
func singleNumberBySort(nums []int) int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	r := nums[0]
	cnt := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == r {
			cnt++
		} else if nums[i] != r && cnt == 1 {
			break
		} else {
			r = nums[i]
		}
	}

	return r
}

// 合并两个有序数组，思路要求时间复杂度m+n，空间复杂度1 -------------------------------------------------------------------
// 思路：双指针  其实双指针已经给出来了，同时
func mergeNum(nums1 []int, m int, nums2 []int, n int) {
	idx := len(nums1) - 1

	for idx >= 0 {
		if m-1 < 0 {
			nums1[idx] = nums2[n-1]
			n--
		} else if n-1 < 0 {
			nums1[idx] = nums1[m-1]
			m--
		} else {

			if nums1[m-1] > nums2[n-1] {
				nums1[idx] = nums1[m-1]
				m--
			} else {
				nums1[idx] = nums2[n-1]
				n--
			}

		}
		idx--
	}

}

// 第一个错误行为-------------------------------------------------------------------------------------------------------
// 思路：两分
func firstBadVersion(n int) int {
	s, e := 1, n

	for s < e {
		mid := (s + e) / 2
		if true {
			s = mid + 1
		} else {
			e = mid
		}
	}

	return s
}

// 有效的平方和 -------------------------------------------------------------------------------------------------------
// 思路：两分，1-num直接一定存在1个数等于
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	s, e := 1, num

	for s < e {
		mid := (s + e) / 2

		if mid*mid < num {
			s = mid + 1
		} else if mid*mid > num {
			e = mid
		} else {
			return true
		}
	}

	return false
}
