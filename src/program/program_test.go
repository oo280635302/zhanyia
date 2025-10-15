package program

import (
	"math/rand"
	"testing"
)

func TestList(t *testing.T) {
	var arr []*ExamRoom
	for i := 0; i < 100; i++ {
		arr = append(arr, &ExamRoom{n: i})
	}
	arr = cs2(arr, &ExamRoom{n: 50})

	for i := 0; i < len(arr); i++ {
		t.Log(arr[i], arr[i].n)
	}
}

// 构造随机路径
func generatePath(n int) []MapPointFloat {
	points := make([]MapPointFloat, n)
	for i := 0; i < n; i++ {
		points[i] = MapPointFloat{
			X: rand.Float64() * 1000,
			Y: rand.Float64() * 1000,
		}
	}
	return points
}

func TestSubPathByCircle(t *testing.T) {
	point := MapPointFloat{X: 391.9, Y: 1450.35}
	radius := 1.0
	path := []MapPointFloat{{X: 400, Y: 1447.5}, {X: 392.88, Y: 1450.5123}} // 模拟路径点数

	res := subPathByCircle(path, point, radius)
	t.Log(path)
	t.Log(res)
}

func BenchmarkSub(b *testing.B) {
	point := MapPointFloat{X: 808.6, Y: 857.52}
	radius := 100.0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		path := generatePath(4) // 模拟路径点数
		subPathByCircle(path, point, radius)
	}
}

func cs1(arr []*ExamRoom, room *ExamRoom) []*ExamRoom {
	arr = append(arr, &ExamRoom{})
	copy(arr[6:], arr[5:])
	arr[5] = room
	return arr
}

func cs2(arr []*ExamRoom, room *ExamRoom) []*ExamRoom {
	arr = append(arr[:5], append([]*ExamRoom{room}, arr[6:]...)...)
	return arr
}
