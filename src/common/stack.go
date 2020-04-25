package common

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxTop int      // 栈的最大数量
	Top    int      // 栈顶
	arr    [100]int // 栈数据
}

// 入栈
func (s *Stack) Push(val int) error {
	// 当满了就不能进栈了
	if s.Top == s.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	s.Top++
	s.arr[s.Top] = val
	return nil
}

// 出栈
func (s *Stack) Pop() (int, error) {
	// 没有就不出了
	if s.Top == -1 {
		fmt.Println("stack empty")
		return -1, errors.New("stack empty")
	}
	val := s.arr[s.Top]
	s.Top--
	return val, nil
}

// 遍历当前栈情况
func (s *Stack) List() {
	if s.Top == -1 {
		return
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Println(i, "=", s.arr[i])
	}
}
