package program

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
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
// 在柱形矩形图的基础上 操作 从上到下增加层数 并视为柱状图 当遇到0时其向上的柱状图崩裂为0（因为是矩形,遇到0就会成为不规则图形）4ms/3.8mb
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	maxNum := 0
	m, n := len(matrix), len(matrix[0])
	height := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//每一行去找1的高度
			//如果不是‘1’，则将当前高度置为0 - 保证其实柱状图
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				//是‘1’，则将高度加1
				height[j] = height[j] + 1
			}
		}
		//更新最大矩形的面积 -- 每个
		maxNum = max(maxNum, LargestRectangleArea(height))
	}
	return maxNum
}

// 二叉树的中序遍历 左->中->右---------------------------------------------------------------------------
// 解法一：迭代+stack的方法 0ms/2mb
func inorderTraversal(root *TreeNode) []int {
	reply := make([]int, 0)

	stack := make([]*TreeNode, 0)

	// 当stack存在 或者 root有值时(第一次||只有右支点的情况)
	for root != nil || 0 < len(stack) {

		// 追到当前root最左边的子结点(如果左边没支点了,插入他本身)，同时将路过的所有node存入stack中
		for root != nil {
			stack = append(stack, root) //入栈
			root = root.Left            //移至最左
		}

		// 将栈顶的数据给即 目前的最左数给arr（即左为nil 右为未知的子节点）
		index := len(stack) - 1
		reply = append(reply, stack[index].Val)
		root = stack[index].Right //右节点会进入下次循环，如果 =nil，继续出栈
		stack = stack[:index]     //出栈
	}
	return reply
}

// 解法二:递归 最简单的方式 0ms/2mb
func InorderTraversal(root *TreeNode) []int {
	reply := make([]int, 0)
	recur(root, &reply)
	return reply
}
func recur(root *TreeNode, arr *[]int) {
	if root != nil {
		recur(root.Left, arr)
		*arr = append(*arr, root.Val)
		recur(root.Right, arr)
	}
}

// 二叉树的锯齿形层次遍历 -递归---------------------------------------------------------------------------
// 广度优先算法BFS+递归+列表 ，将每层进行扫描，判断插入 0ms/2.5mb
func zigzagLevelOrder(root *TreeNode) [][]int {
	var doubleArr [][]int
	if root == nil {
		return doubleArr
	}

	lis := make([]*TreeNode, 0)
	lis = append(lis, root)
	leftJudge := true
	// 当lis有数时进入
	for len(lis) > 0 {
		l := len(lis)
		ans := make([]int, l)
		for i := 0; i < l; i++ {
			node := lis[i]
			if node.Left != nil {
				lis = append(lis, node.Left)
			}
			if node.Right != nil {
				lis = append(lis, node.Right)
			}

			if leftJudge {
				ans[i] = node.Val
			} else {
				ans[l-i-1] = node.Val
			}

		}
		// 反转
		leftJudge = !leftJudge
		doubleArr = append(doubleArr, ans)
		// 截断已经扫描过的该层各数
		lis = lis[l:]
	}
	return doubleArr
}

// 二叉树的前序遍历 中->左->右 -------------------------------------------------------------------------------------
// 在二叉树的中序遍历基础上修改 -- 栈放每次支点的右子节点 支点val每次都放入结果中,同时每次都往左追下一格 0ms/2mb\
// 方法一:迭代
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	stack := make([]*TreeNode, 0)

	// 当节点与 stack都没数据了就结束
	for root != nil || len(stack) != 0 {

		// 将所有根节点存入返回,右节点存入栈中 节点往左移动
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}

		// 弹出栈 将其赋值给root - 不用担心stack 没有走不到这来
		index := len(stack) - 1
		node := stack[index]
		stack = stack[:index]
		root = node
	}
	return res
}

// 方法二:递归 - 中左右 完成 0ms/2mb
func PreorderTraversal(root *TreeNode) []int {
	var arr = make([]int, 0)
	preorderRecur(root, &arr)
	return arr
}
func preorderRecur(root *TreeNode, arr *[]int) {
	if root != nil {
		// 到这里支点处就将支点值放入 - 然后找他的左支点与他的右支点
		*arr = append(*arr, root.Val)
		preorderRecur(root.Left, arr)
		preorderRecur(root.Right, arr)
	}
}

// 二叉树的后序遍历 左右中 --------------------------------------------------------------------------------------------
// 方法一:迭代 将每个节点的右左节点都存入栈中  先存右 再存左 判断过的将其左右节点置成nil 不用再存 从栈顶中弹出来继续 0ms/2mb
func PostorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	// 当栈中有元素时继续  直到数据全部被取完
	for len(stack) > 0 {

		node := stack[len(stack)-1]
		// 先插入右  再插入左  当节点是最子节点就等于找到底了 释放
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		// 找到底 弹出栈 并赋值res
		if node.Left == nil && node.Right == nil {
			res = append(res, node.Val)
			stack = stack[:len(stack)-1]
		}
		// 已经找过左右子节点的节点 不用再找
		node.Left = nil
		node.Right = nil
	}
	return res
}

// 方法二:递归 0ms/2mb
func postorderTraversal(root *TreeNode) []int {
	var arr = make([]int, 0)
	postorderRecur(root, &arr)
	return arr
}
func postorderRecur(root *TreeNode, arr *[]int) {
	if root != nil {
		// 先放其左子节点 在放其右子节点
		postorderRecur(root.Left, arr)
		postorderRecur(root.Right, arr)
		*arr = append(*arr, root.Val)
	}
}

// 二进制求和 ---------------------------------------------------------------------------------------------------------
// 思路用栈 匹配每一个数据 然后进制 0ms/2.6mb
func AddBinary(a string, b string) string {
	// 是否进位
	res := ""
	digit := false
	for len(a) != 0 || len(b) != 0 || digit {
		curA, curB := 0, 0
		if len(a) > 0 {
			curA, _ = strconv.Atoi(string(a[len(a)-1]))
			a = a[:len(a)-1]
		}
		if len(b) > 0 {
			curB, _ = strconv.Atoi(string(b[len(b)-1]))
			b = b[:len(b)-1]
		}
		curAll := 0
		if digit {
			curAll = curA + curB + 1
		} else {
			curAll = curA + curB
		}

		if curAll%2 == 0 && curAll/2 == 1 {
			res = "0" + res
			digit = true
		} else if curAll%2 == 0 && curAll/2 == 0 {
			if res == "" {
				res = "0" + res
			}
			digit = false
		} else if curAll%2 == 1 && curAll/2 == 1 {
			res = "1" + res
			digit = true
		} else {
			res = "1" + res
			digit = false
		}
	}
	// 去0
	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}
	return res
}

// 逆波兰表达式求值 ---------------------------------------------------------------------------------------------
// 将数字都存入栈中 - 当遇到运算符时 就数字弹出来 计算并存进去 只要他逆波兰写法是对的 那就没问题 4ms/4.8mb
func EvalRPN(tokens []string) int {
	n := len(tokens)
	if n == 0 {
		return 0
	}
	// 栈
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		// 将字符串转成数字 - 失败说明他是运算符
		curNum, err := strconv.Atoi(tokens[i])
		// 运算符 弹出两个数 进行计算
		if err != nil {
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res := 0
			//fmt.Println(a,b,tokens[i],res)
			switch tokens[i] {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			fmt.Println(a, b, tokens[i], res)
			stack = append(stack, res)
			continue
		}
		stack = append(stack, curNum)
	}
	return stack[0]
}

// 二叉搜索树迭代器---------------------------------------------------------------------------------------------------
// 解题思路: 用 中序遍历(栈+迭代) 将tree进行排序 然后每次弹出即可 44ms/12.5mb
type BSTIterator struct {
	tree []int
}

func newConstructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		tree: inorderTraversal(root),
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	res := this.tree[0]
	this.tree = this.tree[1:]
	return res
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return len(this.tree) > 0
}

// 基本计算器----------------------------------------------------------------------------------------------------------
// 思路:栈 (就将已经计算的结果存入栈中 与(前的预算符号存入栈中,遇到)将栈中之前结果与运算符弹出来 将 现在结果（运算符决定+-）与之前结果相加 0 ms/3.2mb
func Calculate(s string) int {
	num := 0  // 提取s中的数字
	res := 0  // 返回的计算结果
	sign := 1 // 标志记录运算符号
	stack := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// 提取s中的数字
			num = 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				// 注意处理多位数字
				num = 10*num + int(s[i]-'0')
			}
			// 根据前面的记录，进行运算
			res += sign * num
			// 此时 s[i] 已经不是数字了
			// for 语句中，会再＋1，所以这里先 -1
			i--
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			// 遇到 '(' 就把当前的 res 和 sign 入栈，保存好当前的运行环境
			stack = append(stack, res, sign)
			// 对 res 和 sign 赋予新的值
			res = 0
			sign = 1
		case ')':
			// 遇到')'出栈
			// sign 是与这个')'匹配的'('前的运算符号
			sign = stack[len(stack)-1]
			// temp 是sign前的运算结果
			temp := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			// '(' 与 ')' 之间的运算结果
			//          ↓
			res = sign*res + temp
		}
	}
	return res
}
