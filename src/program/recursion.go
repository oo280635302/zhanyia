package program

// 整数替换
func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	// 偶数 只需要移动一步
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}
	// 遇到奇数需要移动2步 ，并且找到n+1,n-1规划路径中比较小的那一条
	return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
}
