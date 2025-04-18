package program

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MAX_LEVEL = 3
)

type SkipNode struct {
	key   int64
	score int64
	value interface{}

	back []*SkipNode // 上个节点 用于反向遍历 或者 删除
	next []*SkipNode // 多层级 可能有
}

type SkipList struct {
	header *SkipNode
	r      *rand.Rand
	level  int
	m      map[int64]*SkipNode // 加速查找 key:skipNode.key
	length int
}

func NewSkipList() *SkipList {
	l := &SkipList{
		level: MAX_LEVEL,
		r:     rand.New(rand.NewSource(time.Now().UnixNano())),
		m:     make(map[int64]*SkipNode),
	}
	l.header = &SkipNode{
		next: make([]*SkipNode, MAX_LEVEL),
	}
	return l
}

func (l *SkipList) GetWithValue(key int64) (interface{}, bool) {
	node, ok := l.m[key]
	if !ok {
		return nil, ok
	}
	return node.value, ok
}

func (l *SkipList) GetWithScore(key int64) (int64, bool) {
	node, ok := l.m[key]
	if !ok {
		return 0, ok
	}
	return node.score, ok
}

func (l *SkipList) Set(key int64, value interface{}, score int64) {
	node := l.get(key)
	if node != nil {
		l.remove(node)
	}
	if node == nil {
		node = l.newNode(key, value, score)
	}
	l.insert(node)
}

func (l *SkipList) GetRange(startRank, endRank int) []any {
	if endRank < startRank {
		return nil
	}
	curRank := 1

	var res []any

	node := l.header.next[0]

	for {
		if curRank > endRank {
			break
		}

		if node == nil {
			break
		}

		if startRank <= curRank && curRank <= endRank {
			res = append(res, node.value)
		}

		node = node.next[0]
		curRank++
	}
	return res
}

func (l *SkipList) get(key int64) *SkipNode {
	return l.m[key]
}

// 分数降序排列
func (l *SkipList) insert(node *SkipNode) {
	path := make([]*SkipNode, l.level)
	var prev = l.header
	var next *SkipNode

	// 高度递减查找
	for i := l.level - 1; i >= 0; i-- {
		next = prev.next[i]
		// 同层向右移动查找
		for next != nil && next.score >= node.score {
			prev = next
			next = prev.next[i]
		}
		path[i] = prev
	}

	// 将新增的数据插入到结构里面
	for i := 0; i < len(node.next); i++ {
		node.back[i] = path[i]         // 自己back指针 指向 前指针
		node.next[i] = path[i].next[i] // 自己next指针 指向 前指针的后指针

		if node.next[i] != nil {
			node.next[i].back[i] = node // 后指针的back指针 指向自己
		}
		path[i].next[i] = node // 前指针的next指针 指向自己
	}
	l.m[node.key] = node
	l.length++
}

func (l *SkipList) remove(node *SkipNode) {
	for i := 0; i < len(node.next); i++ {
		back := node.back[i]
		if back.next[i] == node {
			back.next[i] = node.next[i]
		}
	}

	delete(l.m, node.key)
	l.length--
}

func (l *SkipList) newNode(key int64, value interface{}, score int64) *SkipNode {
	level := l.randomLevel()
	n := &SkipNode{
		key:   key,
		score: score,
		value: value,
		back:  make([]*SkipNode, level),
		next:  make([]*SkipNode, level),
	}
	return n
}

func (l *SkipList) print() {
	for i := l.level - 1; i >= 0; i-- {
		fmt.Printf("当前层数:%d  ", i)
		cur := l.header.next[i]
		for cur != nil {
			fmt.Printf("----key:%d,score:%d", cur.key, cur.score)
			cur = cur.next[i]
		}
		fmt.Println()
	}
}

func (l *SkipList) randomLevel() int {
	return l.r.Int()%l.level + 1
}
