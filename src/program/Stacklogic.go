package program

import (
	"container/list"
	"fmt"
	"math"
	"strings"
)

// 栈有关的算法问题 -- LeetCode

// 最小栈------------------------------------------------------------------------------------
// 使用辅助栈 来辅助表示每次插入栈时,最小的元素是谁,这样在推出后,辅助栈也推出 就能找到推出后最小的数 16 ms/7.8mb
type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 用队列实现栈----------------------------------------------------------------------------------
// 用list列表来实现即可 -- 功能都有 - 弹出最后一个数据并删除 0ms/2mb
type MyStack struct {
	*list.List
}

/** Initialize your data structure here. */
func QueueConstructor() MyStack {
	return MyStack{list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
// 删除并返回最后一个数
func (this *MyStack) Pop() int {
	if this.Len() == 0 {
		return -1
	}
	return this.Remove(this.Back()).(int)
}

/** Get the top element. */
// 返回最后一个数
func (this *MyStack) Top() int {
	if this.Len() == 0 {
		return -1
	}
	return this.Back().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Len() == 0
}

// 用栈实现队列----------------------------------------------------------------------------------
// 使用辅助栈完成 - 每次取值的是否将数据放入辅助栈 0ms/2.4mb
type MyQueue struct {
	PushStack []int
	PopStack  []int
}

/** Initialize your data structure here. */
func ConstructorQueue() MyQueue {
	return MyQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.PushStack = append(this.PushStack, x)
}

/** Removes the element from in front of queue and returns that element. */
// 第一次是n 后面都是1
func (this *MyQueue) Pop() int {
	if len(this.PopStack) == 0 {
		for i := len(this.PushStack) - 1; i >= 0; i-- {
			this.PopStack = append(this.PopStack, this.PushStack[i])
		}
		this.PushStack = nil
	}
	pop := this.PopStack[len(this.PopStack)-1]
	this.PopStack = this.PopStack[:len(this.PopStack)-1]
	return pop
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.PopStack) == 0 {
		for i := len(this.PushStack) - 1; i >= 0; i-- {
			this.PopStack = append(this.PopStack, this.PushStack[i])
		}
		this.PushStack = nil
	}
	return this.PopStack[len(this.PopStack)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.PushStack) == 0 && len(this.PopStack) == 0 {
		return true
	}
	return false
}

// 接雨水-------------------------------------------------------------------------------
// 使用栈来处理 -- 雨水始终保存在两个顶峰之间  -- 用栈来保存两个顶峰间的数据 4ms/2.8mb
func Trap(height []int) int {
	sum := 0
	stack := make([]int, 0)

	current := 0

	for current < len(height) {
		// 栈中没数据 跳过 直接将数据插入栈   ----   当 当前点 > 栈顶点时 进入判断 -- 直到将栈中数据全部计算完就退出
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {

			// 雨水格计算
			// 将纵向计算 如 1+2+1 转成 横向计算 1 + 0 + 3
			h := height[stack[len(stack)-1]] // h=低的位置

			// 将栈数据弹出来 - 弹出来后栈没数据,即到底了不用计算了
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				break
			}

			distance := current - stack[len(stack)-1] - 1            // 距离 当前点到
			min := min(height[stack[len(stack)-1]], height[current]) // min为顶峰的小值
			fmt.Println(current, stack[len(stack)-1], min, h)
			sum += distance * (min - h)
		}

		stack = append(stack, current)
		current++

	}
	return sum
}

// 简化路径-------------------------------------------------------------------------------
// 使用栈的思想 遇到/.. 就弹栈  遇到/字母就进栈  其余不管 然后展开栈 4ms/4.2mb
// 利用split 将目录 .. .提取出来
func SimplifyPath(path string) string {
	buf := strings.Split(path, "/")
	var stack []string

	for i := 0; i < len(buf); i++ {
		if buf[i] == "" || buf[i] == "." {
			continue
		}
		if buf[i] == ".." {
			if len(stack) > 0 {
				stack = stack[0 : len(stack)-1]
			}
		} else {
			stack = append(stack, buf[i])
		}
	}

	return "/" + strings.Join(stack, "/")
}

// 柱状图中最大的矩形---------------------------------------------------------------------
// 使用栈 当遇到当前点<栈顶元素时 就能去计算栈顶元素他在矩形面积  20ms/5.8mb
func LargestRectangleArea(heights []int) int {

	res := 0

	stack := []int{} // 通过slice维护一个单调递增栈

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			h := heights[stack[len(stack)-1]] // 以出栈元素为高，计算最大矩形的面积
			stack = stack[:len(stack)-1]

			var w int // 计算宽
			if len(stack) == 0 {
				w = i
			} else {
				w = i - stack[len(stack)-1] - 1
			}

			s := h * w
			res = max(res, s)
		}

		stack = append(stack, i)
	}

	// 清空栈内元素,确保以每个元素作为高，并计算其面积
	for len(stack) > 0 {
		h := heights[stack[len(stack)-1]] // 以出栈元素为高，计算最大矩形的面积
		stack = stack[:len(stack)-1]

		var w int // 计算宽
		if len(stack) > 0 {
			w = len(heights) - stack[len(stack)-1] - 1
		} else {
			w = len(heights)
		}

		s := h * w
		res = max(res, s)

	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	fmt.Println(x, y)
	return y
}

// 最大矩形图-----------------------------------------------------------------------------
func maximalRectangle(matrix [][]byte) int {

}
