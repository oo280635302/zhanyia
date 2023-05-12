package program

// 枚举

// 有效时间的数目
func countTime(time string) int {
	m := map[int]int{
		0: 3,  // 0,1,2
		1: 10, // 0,1,2,3,4,5,6,7,8,9
		2: 1,  // :
		3: 6,  // 0,1,2,3,4,5
		4: 10, // 0,1,2,3,4,5,6,7,8,9

		5: 2, // 0,1
		6: 4, // 0,1,2,3
		7: 24,
	}

	ans := 1
	for idx, val := range time {
		if val != '?' {
			continue
		}

		switch idx {
		case 0:
			if time[1] >= '4' && time[1] != '?' {
				ans *= m[idx+5]
			} else if time[1] == '?' {
				ans *= m[idx+7]
			} else {
				ans *= m[idx]
			}
		case 1:
			if time[0] == '2' {
				ans *= m[idx+5]
			} else if time[0] == '?' {
				continue
			} else {
				ans *= m[idx]
			}
		default:
			ans *= m[idx]
		}
	}

	return ans
}
