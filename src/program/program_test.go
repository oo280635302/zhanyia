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
	l := []int{}
	for i := 0; i < 1000; i++ {
		l = append(l, i)
	}

	for i := 0; i < 1; i++ {
		Arr100(l)
	}
}

func Arr100(arr []int) {
	for i := 0; i < 1000; i++ {
		if i == 5 {
			arr = append(arr[:i], arr[i+1:]...)
			break
		}
	}
	arr = append(arr, 5)
}
