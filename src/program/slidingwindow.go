package program

// 2981. 找出出现至少三次的最长特殊子字符串 I
// 思路：滑动窗口----找到了连续相同字符的长度即可
func maximumLength(s string) int {
	if len(s) <= 2 {
		return -1
	}

	m := map[int32]int{}
	var last int32
	var cnt int
	for _, v := range s {
		if v == last {
			cnt++
			for i := 0; i <= cnt; i++ {
				cur := int32(i)*10000 + v
				m[cur]++
			}
		} else {
			cnt = 0
			m[v]++
			last = v
		}
	}

	maxLen := -1
	for cur, num := range m {
		if num >= 3 {
			length := cur / 10000
			maxLen = max(maxLen, int(length+1))
		}
	}

	return maxLen
}
