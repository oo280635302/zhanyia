package program

import (
	"container/heap"
	"fmt"
)

func Ingress() {
	res := lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
	fmt.Println("res:", res)
}

//输入：[2,7,4,1,8,1]
//输出：1
//解释：
//先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
//再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
//接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
//最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。
func lastStoneWeight(stones []int) int {
	h := IntHeap{}
	h = stones
	heap.Init(&h)

	for h.Len() > 1 {
		// 吐两个最大出来
		a1 := heap.Pop(&h).(int)
		a2 := heap.Pop(&h).(int)
		fmt.Println(a1, a2, a1-a2)
		if a1-a2 == 0 {
			continue
		} else {
			heap.Push(&h, a1-a2)
		}
	}
	r := 0
	if h.Len() != 0 {
		r = heap.Pop(&h).(int)
	}

	return r
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
