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
