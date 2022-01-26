package program

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
