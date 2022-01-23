package program

import "container/heap"

// 堆有关的算法问题 -- LeetCode

// 堆工具---------------------------------------------------------------------------------------------------------------

// 堆排序
func HeapSort(s []int) {
	N := len(s) - 1 //s[0]不用，实际元素数量和最后一个元素的角标都为N
	//构造堆
	//如果给两个已构造好的堆添加一个共同父节点，
	//将新添加的节点作一次下沉将构造一个新堆，
	//由于叶子节点都可看作一个构造好的堆，所以
	//可以从最后一个非叶子节点开始下沉，直至
	//根节点，最后一个非叶子节点是最后一个叶子
	//节点的父节点，角标为N/2
	for k := N / 2; k >= 1; k-- {
		sink(s, k, N)
	}
	//下沉排序
	for N > 1 {
		swap(s, 1, N) //将大的放在数组后面，升序排序
		N--
		sink(s, 1, N)
	}
}

//下沉（由上至下的堆有序化）
func sink(s []int, k, N int) {
	for {
		i := 2 * k
		if i > N { //保证该节点是非叶子节点
			break
		}
		if i < N && s[i+1] > s[i] { //选择较大的子节点
			i++
		}
		if s[k] >= s[i] { //没下沉到底就构造好堆了
			break
		}
		swap(s, k, i)
		k = i
	}
}

func swap(s []int, i int, j int) {
	s[i], s[j] = s[j], s[i]
}

type IHeap []int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 有序矩阵中第K小的元素-------------------------------------------------------------------------------------------------
// 思路：将矩阵转成数组 然后 堆排序(任何排序都可以) 40ms/7mb
func kthSmallest(matrix [][]int, k int) int {
	// 将矩阵转成数组
	arr := make([]int, 0)
	for _, val := range matrix {
		arr = append(arr, val...)
	}
	HeapSort(arr)
	return arr[k-1]
}

// 石头组消耗碰撞，得出最后剩余的石头重量---------------------------------------------------------------------------------
// 思路：堆，每次吐出两个比较得出结果后
func lastStoneWeight(stones []int) int {
	h := IHeap{}
	h = stones
	heap.Init(&h)
	for h.Len() > 1 {
		// 吐两个最大出来
		a1 := heap.Pop(&h).(int)
		a2 := heap.Pop(&h).(int)
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

// 股票价格波动
// 思路：最大金额/最小金额 用两个堆来保存（都是小顶堆，用金额的负数来获取最大金额）
type StockPrice struct {
	maxPrice, minPrice heapPair
	timePriceMap       map[int]int
	maxTimestamp       int
}

func ConstructorStockPrice() StockPrice {
	return StockPrice{timePriceMap: map[int]int{}}
}
func (sp *StockPrice) Update(timestamp, price int) {
	heap.Push(&sp.maxPrice, timePrice{-price, timestamp})
	heap.Push(&sp.minPrice, timePrice{price, timestamp})
	sp.timePriceMap[timestamp] = price
	if timestamp > sp.maxTimestamp {
		sp.maxTimestamp = timestamp
	}
}

func (sp *StockPrice) Current() int {
	return sp.timePriceMap[sp.maxTimestamp]
}
func (sp *StockPrice) Maximum() int {
	for {
		if p := sp.maxPrice[0]; -p.price == sp.timePriceMap[p.timestamp] {
			return -p.price
		}
		heap.Pop(&sp.maxPrice)
	}
}
func (sp *StockPrice) Minimum() int {
	for {
		if p := sp.minPrice[0]; p.price == sp.timePriceMap[p.timestamp] {
			return p.price
		}
		heap.Pop(&sp.minPrice)
	}
}

type timePrice struct{ price, timestamp int }
type heapPair []timePrice

func (h heapPair) Len() int            { return len(h) }
func (h heapPair) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h heapPair) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *heapPair) Push(v interface{}) { *h = append(*h, v.(timePrice)) }
func (h *heapPair) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
