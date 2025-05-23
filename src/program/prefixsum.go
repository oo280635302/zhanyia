package program

import (
	"container/heap"
	"fmt"
	"sort"
)

func maxRemoval(nums []int, queries [][]int) int {
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})

	pq := &IHeap{}

	heap.Init(pq)

	deltaArray := make([]int, len(nums)+1)

	operations := 0

	for i, j := 0, 0; i < len(nums); i++ {
		operations += deltaArray[i]
		for j < len(queries) && queries[j][0] == i {
			heap.Push(pq, queries[j][1])
			j++
		}
		for operations < nums[i] && pq.Len() > 0 && (*pq)[0] >= i {
			operations += 1
			deltaArray[heap.Pop(pq).(int)+1] -= 1
		}
		if operations < nums[i] {
			return -1
		}
	}

	return pq.Len()
}

/*
现在有一个getValidPos函数：获取一边区域里面的可以放置建筑的合法点位
area是二维区域地图的二位数组，如果这个二位数组的值为0表示空闲 1表示有障碍物
sideX是物体占x的长度、sideY是物体占y的长度
现在需要一个性能优异尽量减少时间的算法获取合法点位
*/
func getValidPos(area [][]uint8, sideX, sideY int32) []pos {
	rows := len(area)
	if rows == 0 {
		return nil
	}
	cols := len(area[0])
	if cols == 0 {
		return nil
	}

	// 检查建筑尺寸是否超过区域
	if sideX > int32(rows) || sideY > int32(cols) {
		return nil
	}

	// 构建前缀和数组
	prefixSum := make([][]int, rows+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, cols+1)
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			prefixSum[i][j] = int(area[i-1][j-1]) + prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
		}
	}

	fmt.Println("前缀和", prefixSum)

	var validPositions []pos
	maxI := rows - int(sideX)
	maxJ := cols - int(sideY)

	for i := 0; i <= maxI; i++ {
		for j := 0; j <= maxJ; j++ {
			x2 := i + int(sideX)
			y2 := j + int(sideY)
			sum := prefixSum[x2][y2] - prefixSum[i][y2] - prefixSum[x2][j] + prefixSum[i][j]
			if sum == 0 {
				validPositions = append(validPositions, pos{X: int32(i), Y: int32(j)})
			}
		}
	}

	return validPositions
}

func getPrefixSum(area [][]uint8) {
	rows := len(area)
	if rows == 0 {
		return
	}
	cols := len(area[0])
	if cols == 0 {
		return
	}

	// 构建前缀和数组
	prefixSum := make([][]int, rows+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, cols+1)
	}

	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			prefixSum[i][j] = int(area[i-1][j-1]) + prefixSum[i-1][j] + prefixSum[i][j-1] - prefixSum[i-1][j-1]
		}
	}
}

// 二维树状数组
type FenwickTree2D struct {
	tree [][]int
	rows int
	cols int
}

func NewFenwickTree2D(rows, cols int) *FenwickTree2D {
	tree := make([][]int, rows+1)
	for i := range tree {
		tree[i] = make([]int, cols+1)
	}
	return &FenwickTree2D{
		tree: tree,
		rows: rows,
		cols: cols,
	}
}

// 单点更新
func (ft *FenwickTree2D) update(x, y, delta int) {
	for i := x; i <= ft.rows; i += i & -i {
		for j := y; j <= ft.cols; j += j & -j {
			ft.tree[i][j] += delta
		}
	}
}

// 矩形区域更新（差分技巧）
func (ft *FenwickTree2D) updateRange(x1, y1, x2, y2, delta int) {
	ft.update(x1, y1, delta)
	ft.update(x1, y2+1, -delta)
	ft.update(x2+1, y1, -delta)
	ft.update(x2+1, y2+1, delta)
}

// 前缀和查询
func (ft *FenwickTree2D) query(x, y int) int {
	sum := 0
	for i := x; i > 0; i -= i & -i {
		for j := y; j > 0; j -= j & -j {
			sum += ft.tree[i][j]
		}
	}
	return sum
}

// 判断区域是否全为空闲
func (ft *FenwickTree2D) isAreaFree(x1, y1, x2, y2 int) bool {
	sum := ft.query(x2, y2) - ft.query(x1-1, y2) - ft.query(x2, y1-1) + ft.query(x1-1, y1-1)
	return sum == 0
}

// 动态获取合法位置
func getValidPosDynamic(area [][]uint8, sideX, sideY int32) []pos {
	rows := len(area)
	if rows == 0 || sideX <= 0 || sideY <= 0 {
		return nil
	}
	cols := len(area[0])
	if cols == 0 || int(sideX) > rows || int(sideY) > cols {
		return nil
	}

	// 初始化树状数组（坐标从1开始）
	ft := NewFenwickTree2D(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if area[i][j] == 1 {
				ft.updateRange(i+1, j+1, i+1, j+1, 1)
			}
		}
	}

	var validPositions []pos
	sX := int(sideX)
	sY := int(sideY)

	for i := 0; i <= rows-sX; i++ {
		for j := 0; j <= cols-sY; j++ {
			x1, y1 := i+1, j+1   // 树状数组坐标
			x2, y2 := i+sX, j+sY // 矩形右下角
			if ft.isAreaFree(x1, y1, x2, y2) {
				validPositions = append(validPositions, pos{X: int32(i), Y: int32(j)})
			}
		}
	}

	return validPositions
}
