package program

import (
	"math"
)

// 字符串反转----------------------------------------------------------------------------------------------------------
func reverseString(s []byte) {
	length := len(s)

	for i := 0; i < length/2; i++ {
		s[i], s[length-i-1] = s[length-i-1], s[i]
	}
}

// 数字反转------------------------------------------------------------------------------------------------------------
// 思路：栈
func reverseInt(x int) int {
	rev := 0
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return rev
}
