package program

import (
	"fmt"
	"math"
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
	points := map[pair]bool{}     // 某点上是否有灯
	row := map[int]int{}          // 行
	col := map[int]int{}          // 列
	diagonal := map[int]int{}     // 正对角线	- 最重点的点* 斜率
	antiDiagonal := map[int]int{} // 反对角线
	for _, lamp := range lamps {
		r, c := lamp[0], lamp[1]
		p := pair{r, c}
		if points[p] { // 重复灯只占一个数据
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
		if row[r] > 0 || col[c] > 0 || diagonal[r-c] > 0 || antiDiagonal[r+c] > 0 { // 只要在一个列上就说明被照亮了
			ans[i] = 1
		}
		for x := r - 1; x <= r+1; x++ { // 将附近9个格子的灯都灭掉
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
// 思路：用哈希来保存文件名使用次数， 循环直到找到没有文件名的
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

// 最简分数
// 思路： 主要是跳过有公约数的分子分母，用哈希表保存分子/分母，如果有相同的说明是由公约数的分子分母组合
func simplifiedFractions(n int) []string {
	m := make(map[float64]bool)

	res := make([]string, 0)

	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			//fmt.Println(i,j)
			if !m[float64(j)/float64(i)] {
				res = append(res, fmt.Sprintf("%d/%d", j, i))
				m[float64(j)/float64(i)] = true
			}
		}
	}

	return res
}

// 差的绝对值为 K 的数对数目
func countKDifference(nums []int, k int) int {
	m := make(map[int]int)

	for _, v := range nums { // 计算每个数目出现的次数，并用哈希保存
		m[v]++
	}
	ans := 0
	for key, _ := range m { // 跟据每个数据出现的次数 * （数据+k）出现的次数
		ans += m[key] * m[key+k]
	}

	return ans
}

// 找出星型图的中心节点
func findCenter(edges [][]int) int {
	m := make(map[int]int, 0)

	for _, edg := range edges[:2] {
		m[edg[0]]++
		if m[edg[0]] == 2 { // 只要中心节点会出现2次以上，其他都只出现1次
			return edg[0]
		}
		m[edg[1]]++
		if m[edg[1]] == 2 {
			return edg[1]
		}
	}

	return 0
}

// 判断一个数的数字计数是否等于数位的值
func digitCount(num string) bool {
	m := map[int]int{}
	for _, b := range num {
		m[int(b-'0')]++
	}

	for i, v := range num {
		if m[i] != int(v-'0') {
			return false
		}
	}

	return true
}

// 按列翻转得到最大值等行数  -fuck crazy
func maxEqualRowsAfterFlips(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	mp := make(map[string]int)
	for i := 0; i < m; i++ {
		arr := make([]byte, n)
		for j := 0; j < n; j++ {
			// 如果 matrix[i][0] 为 1，则对该行元素进行翻转
			if matrix[i][j]^matrix[i][0] == 0 {
				arr[j] = '0'
			} else {
				arr[j] = '1'
			}
		}
		s := string(arr)
		mp[s]++
	}

	res := 0

	for _, value := range mp {
		if value > res {
			res = value
		}
	}
	return res
}

// 与对应负数同时存在的最大正整数
func findMaxK(nums []int) int {
	h := map[int]bool{}

	res := -1
	for _, v := range nums {
		if h[-v] {
			res = max(res, int(math.Abs(float64(v))))
		}
		h[v] = true
	}

	return res
}
