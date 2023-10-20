package program

import (
	"container/list"
	"testing"
)

func BenchmarkList(b *testing.B) {
	l := list.List{}
	for i := 0; i < 1000; i++ {
		l.PushBack(i)
	}

	for i := 0; i < 1; i++ {
		List100(l)
	}
}

func List100(l list.List) {
	elem := l.Front()
	for elem != nil {
		if elem.Value.(int) == 5 {
			l.Remove(elem)
			break
		}
		elem = elem.Next()
	}
	l.PushBack(5)
}

func BenchmarkArr(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	m := map[int]bool{5: true, 7: true}
	for i := 0; i < b.N; i++ {
		FindForMap(arr, m)
	}
}

func FindForMap(arr []int, m map[int]bool) bool {
	for _, v := range arr {
		if m[v] {
			return true
		}
	}
	return false
}

func FindForArr(arr, arr1 []int) bool {
	for _, v := range arr {
		for _, v2 := range arr1 {
			if v == v2 {
				return true
			}
		}
	}
	return false
}
