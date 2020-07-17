package program

import "fmt"

// 二分查找算法问题 -- leetCode

// 有序矩阵中第K小的元素------------------------------------------------------------------------------------------
// 思路：二分查找  根据矩形特性 左上最小->右下最大  以中间断开查找数字 32ms/6.2mb
// 左边(小的一方)的个数 < k = mid应该在右边 同理 否则在左边
// 将边界移动到之前的中值点 ,继续二分 一直到左 >= 右时 即找到了该数
func KthSmallest2(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		//fmt.Println(left,right)
		mid := left + (right-left)/2
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0
	num := 0
	for i >= 0 && j < n {
		fmt.Println(num, mid, k, n)
		if matrix[i][j] <= mid {
			num += i + 1
			j++
		} else {
			i--
		}
	}
	return num >= k
}

// 搜索插入位置-------------------------------------------------------------------------------------------------------
// 思路：二分/tree  或者直接 遍历  均为 4ms,3.1mb
func SearchInsert(nums []int, target int) int {
	// 左右指针
	left, right := 0, len(nums)-1
	// 中间值
	mid := 0
	// 当做 左<右 时能继续寻找
	for left <= right {
		mid = (left + right) / 2

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	return left
}
func searchInsert(nums []int, target int) int {
	for k, v := range nums {
		if v >= target {
			return k
		}
	}
	return len(nums)
}
