package program

// 电话号码的字母组合
// 思路：递归,每次都找对应数值的几个不同对应值 当每个的剩余数为空时 就是找到底了 就将其追加到返回值中	0ms/2mb
var phone = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func LetterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	res := make([]string, 0)
	dfsLetterCombinations("", digits, &res)
	return res
}

func dfsLetterCombinations(nowStr string, digits string, res *[]string) {
	if len(digits) == 0 {
		*res = append(*res, nowStr)
		return
	}

	for _, v := range phone[digits[0]] {
		a := nowStr + v
		dfsLetterCombinations(a, digits[1:], res)
	}
}
