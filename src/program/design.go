package program

// 考场就座
type ExamRoom struct {
	n int   // 座位数
	s []int // 已落座的位置
}

func ConstructorExamRoom(n int) ExamRoom {
	return ExamRoom{n: n, s: []int{}}
}

func (this *ExamRoom) Seat() int {
	student := 0
	idx := 0
	if len(this.s) > 0 {
		dist := this.s[0] // 位置宽度
		pre := -1         // 以0为起步
		for i, v := range this.s {
			if pre != -1 {
				d := (v - pre) / 2 // 新增位置到周围位置的宽度
				if d > dist {
					dist = d
					student = pre + d
					idx = i
				}
			}
			pre = v
		}
		if this.n-1-this.s[len(this.s)-1] > dist {
			student = this.n - 1
			idx = len(this.s)
		}
	}
	this.s = append(this.s, 0) // 添加学生到新位置
	copy(this.s[idx+1:], this.s[idx:])
	this.s[idx] = student
	return student
}

func (this *ExamRoom) Leave(p int) {
	idx := 0
	for i := 0; i < len(this.s); i++ { // 离开位置
		if this.s[i] == p {
			idx = i
			break
		}
	}
	this.s = append(this.s[:idx], this.s[idx+1:]...)
}
