package program

import (
	"math/bits"
)

/*
	位元数知识点：
		异^：
			规则：	1^1=0 1^0=1 0^0=0
			非零数异0 = 非零数本身	  5^0=5
			两个相同数异 =  0		  5^5=0 	自反性
			异运算满足交换结合律		  1^2^3^4 = (3^4)^(1^2)
		与|：
			规则：	1|1=1 1|0=1 0|0=0
		或&：
			规则	1&1=1 1&0=0 0&0=0
			n&n-1 会让他的最低1位去掉一个 eg: 1111&1110=1110  1110&1101=1100 1100&1011=1000
*/

// 多数元素，数组出现次数大于n/2的元素-----------------------------------------------------------------------------------
// 思路：计数即可，遇到不相同计算-1，相同+1，cnt=0就继续计当前数
func majorityElement(nums []int) int {
	r := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if r == nums[i] {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			r = nums[i]
		}
	}
	return r
}

// 多数元素2，数组出现次数大于n/3的元素-----------------------------------------------------------------------------------
// 思路2 : 摩尔投票法 抵消，将3个不同的数掉抵消掉获取
func majorityElement2(nums []int) []int {
	elem1, elem2 := 0, 0
	num1, num2 := 0, 0
	cnt1, cnt2 := 0, 0

	// 大于n/3的元素 可知最多又2个，用摩尔投票法如果出现3个不同的就清理掉
	for _, v := range nums {
		if num1 > 0 && elem1 == v {
			num1++

		} else if num2 > 0 && elem2 == v {
			num2++

		} else if num1 == 0 {
			elem1 = v
			num1++
		} else if num2 == 0 {
			elem2 = v
			num2++
		} else {
			num1--
			num2--
		}

	}

	// 筛选出来的两个数最有可能超过n/3 ，但存在[1,2] 这种特殊情况将筛选出来的数统计下数量
	res := make([]int, 0)
	for _, v := range nums {
		if v == elem1 {
			cnt1++
		} else if v == elem2 {
			cnt2++
		}
	}

	// 超过n/3的就添加
	if cnt1 > len(nums)/3 {
		res = append(res, elem1)
	}
	if cnt2 > len(nums)/3 {
		res = append(res, elem2)
	}

	return nums
}

// 数组中只有1个数存在1个，其余都存在2个，找出这个数----------------------------------------------------------------------
// 思路：异运算， x^0=x，x^x=0，1^0^1^2=2
func singleNumberByBit(nums []int) int {
	r := 0

	for _, v := range nums {
		r ^= v
	}

	return r
}

// 2进制中1的个数--------------------------------------------------------------------------------------------------------
// 思路：或运算，n与n-1或会拼掉位运算的最底位的1，直到拼完计算次数
func hammingWeight(num uint32) int {
	r := 0
	for ; num > 0; num &= num - 1 {
		r++
	}
	return r
}

// 两个字符串有一个字母不同，找出来---------------------------------------------------------------------------------------
// 思路：异运算的，相同抵消特性
func findTheDifference(s string, t string) byte {
	r := int32(0)
	for _, v := range s {
		r ^= v
	}
	for _, v := range t {
		r ^= v
	}
	return byte(r)
}

// 两个数字对应二进制位不同的位置的数目-----------------------------------------------------------------------------------
// 思路：异
func hammingDistance(x int, y int) int {
	z := x ^ y
	return bits.OnesCount(uint(z))
}
