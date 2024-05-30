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

func maximumLength2(s string) int {
	n := len(s)
	cnt := make(map[byte][]int)
	// 获得每个字符每段连续相同字符的长度
	for i, j := 0, 0; i < n; i = j {
		for j < n && s[j] == s[i] {
			j++
		}
		cnt[s[i]] = append(cnt[s[i]], j-i)
	}

	res := -1
	for _, vec := range cnt {
		lo, hi := 1, n-2
		// 两分查找 看满足总字符长度1/2的3连续。找不到再找1/4 找到了就找3/4...
		for lo <= hi {
			mid := (lo + hi) >> 1
			count := 0
			for _, x := range vec {
				if x >= mid {
					count += x - mid + 1
				}
			}
			if count >= 3 {
				if mid > res {
					res = mid
				}
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return res
}
