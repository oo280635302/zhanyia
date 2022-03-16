package program

import (
	"container/list"
)

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

// 全 O(1) 的数据结构
// 查询/增加1/扣除1 复杂度都为1的计数排序
type AllOne struct {
	Count map[string]*list.Element // 计数
	List  *list.List               // 升序排序
}

type allOneNode struct {
	count int             // 当前节点的分数
	word  map[string]bool // 当前分数的词库
}

func ConstructorAllOne() AllOne {
	return AllOne{
		Count: make(map[string]*list.Element),
		List:  list.New(),
	}
}

func (a *AllOne) Inc(key string) {
	curList := a.Count[key]
	// 不存在 就插入链首
	if curList == nil {
		lf := a.List.Front()
		// 空链表，头链首cnt>1 新增头节点插入
		if lf == nil || lf.Value.(*allOneNode).count > 1 {
			element := a.List.PushFront(&allOneNode{word: map[string]bool{key: true}, count: 1})
			a.Count[key] = element
			return
		}

		// 有1的情况 直接插入1的链数据里
		lf.Value.(*allOneNode).word[key] = true
		a.Count[key] = lf
		return
	}

	// 存在
	nextList := curList.Next()
	curNode := curList.Value.(*allOneNode)
	// 需要插入链表
	if nextList == nil || nextList.Value.(*allOneNode).count > curNode.count+1 {
		// 增加下一位，指向下一位，删除当前位的key
		a.Count[key] = a.List.InsertAfter(&allOneNode{count: curNode.count + 1, word: map[string]bool{key: true}}, curList)
		delete(curNode.word, key)
		if len(curNode.word) == 0 {
			a.List.Remove(curList)
		}
		return
	}
	// 不需要插入链表只需要将数据加入链表的m里面
	nextList.Value.(*allOneNode).word[key] = true
	a.Count[key] = nextList
	delete(curNode.word, key)
	if len(curNode.word) == 0 {
		a.List.Remove(curList)
	}
}

func (a *AllOne) Dec(key string) {
	curList := a.Count[key]
	curNode := curList.Value.(*allOneNode)
	prevList := curList.Prev()
	// 踢出链表 cnt<=0
	if curNode.count-1 == 0 {
		delete(curNode.word, key)
		if len(curNode.word) == 0 {
			a.List.Remove(curList)
		}
		delete(a.Count, key)
		return
	}

	// 需要插入链表
	if prevList == nil || prevList.Value.(*allOneNode).count < curNode.count-1 {
		a.Count[key] = a.List.InsertBefore(&allOneNode{count: curNode.count - 1, word: map[string]bool{key: true}}, curList)
		delete(curNode.word, key)
		if len(curNode.word) == 0 {
			a.List.Remove(curList)
		}
		return
	}
	// 不需要插入链表 直接加入链表的m
	prevList.Value.(*allOneNode).word[key] = true
	a.Count[key] = prevList
	delete(curNode.word, key)
	if len(curNode.word) == 0 {
		a.List.Remove(curList)
	}
}

func (a *AllOne) GetMaxKey() string {
	// 最大的在最后面 back
	if f := a.List.Back(); f != nil {
		for key := range f.Value.(*allOneNode).word {
			return key
		}
	}
	return ""
}

func (a *AllOne) GetMinKey() string {
	// 最小的最前面 front
	if f := a.List.Front(); f != nil {
		for key := range f.Value.(*allOneNode).word {
			return key
		}
	}
	return ""
}
