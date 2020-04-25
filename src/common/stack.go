package common

import (
	"errors"
	"fmt"
	"strings"
)

type Stack struct {
	//MaxTop int    // 栈的最大数量
	Top      int   // 栈顶
	Arr      []int // 栈数据
	PointNum int   // 记录浮点在栈的位置
}

// 创建个新栈
func NewStack() *Stack {
	return &Stack{
		Top:      -1,
		Arr:      make([]int, 0),
		PointNum: -1,
	}
}

// 将string数字传入栈
func PushStringToStack(s *Stack, str string) {
	for i := 0; i < len(str); i++ {
		s.Push(int(str[i]) - 48)
	}
}

// 入栈
func (s *Stack) Push(val int) {
	s.Top++
	s.Arr = append(s.Arr, val)
}

// 出栈
func (s *Stack) Pop() (int, error) {
	// 没有就不出了
	if s.Top == -1 {
		//fmt.Println("stack empty")
		return -1, errors.New("stack empty")
	}
	val := s.Arr[s.Top]
	s.Top--
	return val, nil
}

// 将栈转换成string
func (s *Stack) StackToString() string {
	var reply string
	for i := len(s.Arr) - 1; i >= 0; i-- {
		reply += string(s.Arr[i] + 48)
	}
	return reply
}

// 遍历当前栈情况
func (s *Stack) List() {
	if s.Top == -1 {
		return
	}
	for i := s.Top; i >= 0; i-- {
		fmt.Println(i, "=", s.Arr[i])
	}
}

// 栈整数加法运算
func StackIntegerAdd(num1 string, num2 string) string {

	// 数字1,数字2栈,结果栈
	num1Stack := NewStack()
	num2Stack := NewStack()
	resultStack := NewStack()

	// 存入栈
	PushStringToStack(num1Stack, num1)
	PushStringToStack(num2Stack, num2)

	// 进位数的值 0,1
	var carry int
	// 计算
	for {
		a, err1 := num1Stack.Pop()
		b, err2 := num2Stack.Pop()
		// 都没有 = 结束
		if err1 != nil && err2 != nil && carry == 0 {
			break
		}
		if err1 != nil {
			a = 0
		}
		if err2 != nil {
			b = 0
		}
		curNum := (a + b + carry) % 10
		carry = (a + b + carry) / 10
		resultStack.Push(curNum)
	}

	return resultStack.StackToString()
}

// 栈加法运算
func StackAdd(num1 string, num2 string) string {
	x := JudgeStringIsFloat(num1)
	y := JudgeStringIsFloat(num2)
	if x && y {
		num1s := strings.Split(num1, ".")
		num2s := strings.Split(num2, ".")
		integer := StackIntegerAdd(num1s[0], num2s[0])
		decimal := StackIntegerAdd(ReverseString(num1s[1]), ReverseString(num2s[1]))
		return integer + "." + ReverseString(decimal)
	} else if x && !y {
		num1s := strings.Split(num1, ".")
		integer := StackIntegerAdd(num1s[0], num2)
		return integer + "." + num1s[1]
	} else if !x && y {
		num2s := strings.Split(num2, ".")
		integer := StackIntegerAdd(num1, num2s[0])
		return integer + "." + num2s[1]
	} else {
		return StackIntegerAdd(num1, num2)
	}
}

// 栈乘法运算 todo
func StackMulti() {

}
