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
	s.Arr = s.Arr[:len(s.Arr)-1]
	s.Top--
	return val, nil
}

// 恢复至栈顶
func (s *Stack) ReserveTop() {
	s.Top = len(s.Arr) - 1
}

// 移动至栈指定位置
func (s *Stack) MoveLocal(local int) {
	s.Top = local
}

// 将栈转换成string
func (s *Stack) StackToString() string {
	var reply string
	for i := len(s.Arr) - 1; i >= 0; i-- {
		reply += string(s.Arr[i] + 48)
	}
	return reply
}

// 删顶栈
func (s *Stack) DelTopStack() {
	s.Arr = s.Arr[:len(s.Arr)-1]
}

// 获得栈指定位置val
func (s *Stack) GetDesignVal(local int) int {
	if len(s.Arr) <= local {
		return -1
	}
	return s.Arr[local]
}

// 替换栈索引内容
func (s *Stack) ReplaceVal(index int, val int) {
	s.Arr[index] = val
}

// 检测栈指定索引是否存在值
func (s *Stack) CheckIndex(index int) bool {
	if len(s.Arr) <= index {
		return false
	}
	return true
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
// 1000/ms
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

// 栈小数加法运算
// 1000/ms
func StackDecimalAdd(num1 string, num2 string) (string, bool) {
	// 小数位数对齐
	diffDigit := len(num1) - len(num2)
	if diffDigit > 0 {
		for i := 0; i < diffDigit; i++ {
			num2 += "0"
		}
	} else if diffDigit < 0 {
		for i := 0; i > diffDigit; i-- {
			num1 += "0"
		}
	}
	realNum := StackIntegerAdd(num1, num2)
	if len(realNum) > len(num1) {
		return realNum, true
	}
	return realNum, false
}

// 栈加法运算
// 200/ms
func StackAdd(num1 string, num2 string) string {
	x := JudgeStringIsFloat(num1)
	y := JudgeStringIsFloat(num2)
	if x && y {
		num1s := strings.Split(num1, ".")
		num2s := strings.Split(num2, ".")
		decimal, carry := StackDecimalAdd(num1s[1], num2s[1])
		if carry {
			return StackIntegerAdd(StackIntegerAdd(num1s[0], num2s[0]), "1") + "." + decimal[1:]
		}

		return StackIntegerAdd(num1s[0], num2s[0]) + "." + decimal
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

// 栈乘法运算
// 100/ms
func StackMulti(num1 string, num2 string) string {
	// 将浮点数 转成 正数 并获取他的小数位数
	rNum1, decimal1 := TakeEveryNumToInt(num1)
	rNum2, decimal2 := TakeEveryNumToInt(num2)

	// 数字1,结果栈
	num1Stack := NewStack()
	resultStack := NewStack()

	// 存入数字1栈
	PushStringToStack(num1Stack, rNum1)

	// 检测真实的位数
	var num2Num int

	for x := 0; x < len(rNum1); x++ {

		// 获得新的数字2
		num2Stack := NewStack()
		PushStringToStack(num2Stack, rNum2)

		b, err1 := num1Stack.Pop()
		// 没有了就结束
		if err1 != nil {
			break
		}
		// 小数点直接跳过
		if b < 0 {
			continue
		}

		var carry, ac int

		for {
			a, err2 := num2Stack.Pop()
			if err2 != nil {
				break
			}
			// 当前的数字
			curNum := a*b + carry
			// 当结果栈对应位数有值时,
			// 		取栈val,与当前计算结果相加取模
			// 当结果栈没有对应位数值时
			// 		push模
			if resultStack.CheckIndex(num2Num + ac) {
				curStack := resultStack.GetDesignVal(num2Num + ac)
				resultStack.ReplaceVal(num2Num+ac, (curNum+curStack)%10)
				carry = (curNum + curStack) / 10
			} else {
				resultStack.Push(curNum % 10)
				carry = curNum / 10
			}
			ac++
		}
		// 进入下个循环还有carry,push进结果栈
		if carry > 0 {
			resultStack.Push(carry)
		}

		num2Num++
	}

	if decimal1+decimal2 > 0 && resultStack.StackToString()[:len(resultStack.StackToString())-decimal1-decimal2] != "" {
		return RmStringNumRemainZero(resultStack.StackToString()[:len(resultStack.StackToString())-decimal1-decimal2] + "." + resultStack.StackToString()[len(resultStack.StackToString())-decimal1-decimal2:])

	} else if decimal1+decimal2 > 0 && resultStack.StackToString()[:len(resultStack.StackToString())-decimal1-decimal2] == "" {
		return "0" + "." + resultStack.StackToString()[len(resultStack.StackToString())-decimal1-decimal2:]
	}

	return RmStringNumRemainZero(resultStack.StackToString())
}
