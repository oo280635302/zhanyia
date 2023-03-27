package program

// 统计只差一个字符的子串数目
func countSubstrings(s string, t string) int {
	ans := 0

	m, n := len(s), len(t)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff := 0

			for k := 0; k+i < m && k+j < n; k++ {
				if s[k+i] != t[k+j] {
					diff++
				}

				if diff > 1 {
					break
				}
				if diff == 1 {
					ans++
				}
			}
		}
	}

	return ans
}
