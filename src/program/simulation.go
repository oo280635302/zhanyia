package program

// 数字流的秩
type streamRankInfo struct {
	val int
	num int
}
type StreamRank struct {
	sortArr []streamRankInfo
}

func ConstructorStreamRank() StreamRank {
	return StreamRank{}
}

func (this *StreamRank) Track(x int) {
	for idx, cur := range this.sortArr {
		if cur.val > x {
			tmp := append([]streamRankInfo{}, this.sortArr[idx:]...)
			this.sortArr = append(this.sortArr[:idx], streamRankInfo{val: x, num: 1})
			this.sortArr = append(this.sortArr, tmp...)
			return
		} else if cur.val == x {
			this.sortArr[idx].num += 1
			return
		}
	}
	this.sortArr = append(this.sortArr, streamRankInfo{val: x, num: 1})
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	ans := 0
	for _, v := range this.sortArr {
		if v.val <= x {
			ans += v.num
		} else {
			break
		}
	}
	return ans
}
