package program

import (
	"strconv"
)

// 检测正方形
// 思路： 正方形的边长固定，跟据一个点 与 固定边长能计算出应该出现的正方形的位置 用hash保存所有点位出现频率，来求出正方形出现次数
type DetectSquares map[int]map[int]int

func ConstructorDetectSquares() DetectSquares {
	return DetectSquares{}
}
func (d DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if d[y] == nil {
		d[y] = make(map[int]int)
	}
	d[y][x]++
}
func (d DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	// 确定好Y轴
	if d[y] == nil {
		return 0
	}

	ans := 0
	yCnt := d[y] // 固定了y轴的所有x点位
	for col, colCnt := range d {
		if col != y {
			distance := col - y                                      // 正方形的边长
			ans += colCnt[x] * yCnt[x+distance] * colCnt[x+distance] // colCnt[x] (x,col) 个数 * yCnt[x+distance] (x+col,y) 个数 * colCnt[x+distance] (x+col,y+col) 个数
			ans += colCnt[x] * yCnt[x-distance] * colCnt[x-distance] // colCnt[x] (x,col) 个数 * yCnt[x+distance] (x-col,y) 个数 * colCnt[x+distance] (x-col,y-col) 个数
		}
	}
	return ans
}

// 网格照明
func gridIllumination(n int, lamps, queries [][]int) []int {
	type pair struct{ x, y int }
	points := map[pair]bool{}	// 某点上是否有灯
	row := map[int]int{}	// 行
	col := map[int]int{}	// 列
	diagonal := map[int]int{}	// 正对角线	- 最重点的点* 斜率
	antiDiagonal := map[int]int{}	// 反对角线
	for _, lamp := range lamps {
		r, c := lamp[0], lamp[1]
		p := pair{r, c}
		if points[p] {	// 重复灯只占一个数据
			continue
		}
		points[p] = true
		row[r]++
		col[c]++
		diagonal[r-c]++
		antiDiagonal[r+c]++
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		r, c := query[0], query[1]
		if row[r] > 0 || col[c] > 0 || diagonal[r-c] > 0 || antiDiagonal[r+c] > 0 {	// 只要在一个列上就说明被照亮了
			ans[i] = 1
		}
		for x := r - 1; x <= r+1; x++ {	// 将附近9个格子的灯都灭掉
			for y := c - 1; y <= c+1; y++ {
				if x < 0 || y < 0 || x >= n || y >= n || !points[pair{x, y}] {
					continue
				}
				delete(points, pair{x, y})
				row[x]--
				col[y]--
				diagonal[x-y]--
				antiDiagonal[x+y]--
			}
		}
	}
	return ans
}

// 保证文件名唯一
func getFolderNames(names []string) []string {
	d := make(map[string]int)
	res := make([]string, 0)
	for _, name := range names {
		s := name
		for d[s] > 0 { // 找到是否已经有相同的文件名，有就加（X）
			s = name + "(" + strconv.Itoa(d[name]) + ")"
			d[name]++
		}
		d[s] = 1
		res = append(res, s)
	}
	return res
}
