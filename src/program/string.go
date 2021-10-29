package program

import (
	"fmt"
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

// 第一个不重复字符-----------------------------------------------------------------------------------------------------
// 思路：数组保存每次字母出现次数，
func firstUniqChar(s string) int {
	m := [26]int{}
	for _, v := range s {
		m[v-'a']++
	}
	for idx, v := range s {
		if m[v-'a'] == 1 {
			return idx
		}
	}
	return -1
}

// 验证回文串-----------------------------------------------------------------------------------------------------------
// 思路：前后双指针
func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
	i, j := 0, len(s)-1

	for i < j {
		if s[i] < 48 || (s[i] > 57 && s[i] < 65) || (s[i] > 90 && s[i] < 97) || s[i] > 122 {
			i++
			continue
		}
		if s[j] < 48 || (s[j] > 57 && s[j] < 65) || (s[j] > 90 && s[j] < 97) || s[j] > 122 {
			j--
			continue
		}

		tempI := s[i]
		if tempI >= 65 && tempI <= 90 {
			tempI += 32
		}

		tempJ := s[j]
		if tempJ >= 65 && tempJ <= 90 {
			tempJ += 32
		}
		fmt.Println(string(tempI), string(tempJ))
		if tempI != tempJ {
			return false
		}
		i++
		j--
	}

	return true
}

// 字符串转换成整数-----------------------------------------------------------------------------------------------------
func myAtoi(s string) int {
	res := 0
	reverse := false

	for _, v := range s {
		if v == ' ' {
			s = s[1:]
		} else {
			break
		}
	}

	for idx, v := range s {
		if v != '-' && v != '+' && (v < '0' || v > '9') {
			break
		}
		if v == '-' {
			if idx == 0 {
				reverse = true
				continue
			}
			break
		}
		if v == '+' {
			if idx == 0 {
				continue
			}
			break
		}

		res = res*10 + int(v-'0')

		if res > math.MaxInt32 && !reverse {
			res = math.MaxInt32
			break
		} else if res > math.MaxInt32+1 && reverse {
			res = math.MaxInt32 + 1
			break
		}
	}

	if reverse {
		res = -res
	}

	return res
}
