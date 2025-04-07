package rank

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
	value RankElemInfo

	back []*SkipNode // 上个节点 用于删除
	next []*SkipNode // 多层级 可能有
	span []int
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
		back: make([]*SkipNode, MAX_LEVEL),
		next: make([]*SkipNode, MAX_LEVEL),
		span: make([]int, MAX_LEVEL),
	}
	return l
}

func (l *SkipList) Get(key int64) (RankElemInfo, bool) {
	node, ok := l.m[key]
	if !ok {
		return RankElemInfo{}, ok
	}

	rank := l.getRank(key)

	v := node.value
	v.CacheRank = int32(rank)
	return v, ok
}

func (l *SkipList) Set(key int64, value RankElemInfo, score int64) {
	node := l.get(key)
	if node != nil {
		l.remove(node)
	}
	if node == nil {
		node = l.newNode(key, value, score)
	}
	l.insert(node)
}

func (l *SkipList) Range(startRank, endRank int32) []RankElemInfo {
	if endRank < startRank {
		return nil
	}
	var curRank int32 = 1

	var res []RankElemInfo
	node := l.header.next[0]

	for {
		if curRank > endRank {
			break
		}

		if node == nil {
			break
		}

		if startRank <= curRank && curRank <= endRank {
			vs := node.value
			vs.CacheRank = curRank
			res = append(res, vs)
		}

		node = node.next[0]
		curRank++
	}
	return res
}

func (l *SkipList) Remove(key int64) {
	node := l.get(key)
	if node == nil {
		return
	}
	l.remove(node)
}

func (l *SkipList) get(key int64) *SkipNode {
	return l.m[key]
}

func (l *SkipList) getRank(key int64) int {
	need, ok := l.m[key]
	if !ok {
		return 0
	}

	node := l.header.next[0]
	curRank := 0
	for {
		if node == nil {
			break
		}
		curRank++
		if node == need {
			break
		}
		node = node.next[0]
	}

	return curRank
}

// 分数降序排列
func (l *SkipList) insert(node *SkipNode) {
	path := make([]*SkipNode, l.level)
	span := make([]int, l.level)
	var prev = l.header
	var next *SkipNode

	// 高度递减查找
	for i := l.level - 1; i >= 0; i-- {
		if i == l.level-1 {
			span[i] = 0
		} else {
			span[i] = span[i+1]
		}

		next = prev.next[i]
		for next != nil && (next.score > node.score || next.score == node.score && next.value.CreateAt.Before(node.value.CreateAt)) {
			prev = next
			span[i] += prev.span[i]

			next = prev.next[i]
		}
		path[i] = prev
	}

	for i := 0; i < len(node.next); i++ {
		node.back[i] = path[i]
		node.next[i] = path[i].next[i]

		oldSpan := path[i].span[i]

		node.span[i] = oldSpan - span[i] + 1

		if node.next[i] != nil {
			node.next[i].back[i] = node
		}

		path[i].next[i] = node

		if path[i].next[i] != nil {
			path[i].span[i] = span[i] + 1
		}
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

func (l *SkipList) newNode(key int64, value RankElemInfo, score int64) *SkipNode {
	level := l.randomLevel()
	n := &SkipNode{
		key:   key,
		score: score,
		value: value,
		back:  make([]*SkipNode, level),
		next:  make([]*SkipNode, level),
		span:  make([]int, level),
	}
	return n
}

func (l *SkipList) randomLevel() int {
	return l.r.Int()%l.level + 1
}

func (l *SkipList) print() {
	for i := l.level - 1; i >= 0; i-- {
		fmt.Printf("当前层数:%d  ", i)
		cur := l.header.next[i]
		for cur != nil {
			fmt.Printf("----key:%d,score:%d,span:%d", cur.key, cur.score, cur.span[i])
			cur = cur.next[i]
		}
		fmt.Println()
	}
}
