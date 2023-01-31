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
	//fmt.Println(x, y)
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

// 去除重复字母(字典序最小，保持相对位置)-------------------------------------------------------------------------------
// 思路:单调栈，根据字典序最小特性 尽量让小字母在左 -- 将stack中不存在的数据插入栈中 当新数据<栈顶元素|栈顶元素后面会出现 就将栈顶元素弹出 4ms/2.2mb
func RemoveDuplicateLetters(s string) string {
	// 存放每个字母出现的次数
	countMap := make(map[int32]int, 0)
	// 栈
	stack := make([]int32, 0)
	// 存放放入栈中的字母元素
	checkMap := make(map[int32]bool, 0)

	// 计算每个数据的次数
	for _, val := range s {
		countMap[val]++
	}

	// 遍历s
	// 已经有的数据不需要再存
	// 新的字符  ：如果<栈顶元素 同时栈顶元素后面还会出现 就让其代替栈顶元素
	//          其余：
	for _, val := range s {
		if checkMap[val] {
			countMap[val]--
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > val && countMap[stack[len(stack)-1]] > 0 {
			// 弹出
			checkMap[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		checkMap[val] = true
		stack = append(stack, val)
		countMap[val]--
	}

	return string(stack)
}

// 验证二叉树的前序序列化-----------------------------------------------------------------------------------------------
// 思路：根节点有2个空节点，每分出去一个支点多一个空子节点，同时要保证一个空节点数不能多余需要的空节点数如 1,#,#,#,2 是错的 0ms/2.7mb
func isValidSerialization(preorder string) bool {
	preArr := strings.Split(preorder, ",")
	check := 1
	for _, val := range preArr {
		check -= 1
		if check < 0 {
			return false
		}

		if val != "#" {
			check = check + 2
		}
	}
	return check == 0
}

// 下一个更大元素 I----------------------------------------------------------------------------------------------------
// 思路：单调栈，将num2 依次放入栈中 遇到栈顶小于本身就将栈顶弹出 放入map中，然后nums1直接找map就完事（不重复的特性）4ms/3.1mb
func NextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0)
	checkMap := make(map[int]int, len(nums2))

	for i := 0; i < len(nums2); i++ {

		// 栈不为空 同时 栈顶数据小于当前数据就将其弹出保存
		for len(stack) != 0 && stack[len(stack)-1] < nums2[i] {
			checkMap[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}

		// 插入数据到栈中
		stack = append(stack, nums2[i])

	}
	for _, val := range stack {
		checkMap[val] = -1
	}

	res := make([]int, 0)
	for _, val := range nums1 {
		res = append(res, checkMap[val])
	}
	return res
}

// 下一个更大元素 II----------------------------------------------------------------------------------------------------
// 思路：单调栈	倒序遍历2次（正序也一样） 目的是为了组成类环形帮后面数字找到其更大元素 2次是最小圈数 28ms/6.4mb
func NextGreaterElements(nums []int) []int {
	stack := make([]int, 0) // 栈保存索引
	n := len(nums)
	res := make([]int, n)
	for i, _ := range res {
		res[i] = -1
	}

	// 倒序遍历
	for i := n - 1; i >= 0; i-- {

		// 前一个数 >= 栈顶时 将栈顶元素弹出
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		// 这时栈还有数 说明 当前数 < 该数 赋值
		if len(stack) > 0 {
			res[i] = nums[stack[len(stack)-1]]
		}

		// 将值追加进栈中
		stack = append(stack, i)
	}

	// 倒序遍历
	for i := n - 1; i >= 0; i-- {

		// 前一个数 >= 栈顶时 将栈顶元素弹出
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		// 这时栈还有数 说明 当前数 < 该数 赋值
		if len(stack) > 0 {
			res[i] = nums[stack[len(stack)-1]]
		}

		// 将值追加进栈中
		stack = append(stack, i)
	}

	return res
}

// 用两个栈实现队列---------------------------------------------------------------------------------------------------
// 思路： 一个出栈 一个入栈 - 出栈没值了就把入栈的值倒进去出栈 244ms,8.4mb
type CQueue struct {
	In  []int
	Out []int
}

func twoStackConstructor() CQueue {
	return CQueue{In: make([]int, 0), Out: make([]int, 0)}
}

// tail In
func (this *CQueue) AppendTail(value int) {
	this.In = append(this.In, value)
}

// head Out
func (this *CQueue) DeleteHead() int {
	if len(this.Out) == 0 && len(this.In) == 0 {
		return -1
	}
	// 如果出栈 为0 将 入栈 倒入 出栈
	if len(this.Out) == 0 {
		for i := len(this.In) - 1; i >= 0; i-- {
			this.Out = append(this.Out, this.In[i])
		}
		this.In = this.In[:0]
	}
	res := this.Out[len(this.Out)-1]
	this.Out = this.Out[:len(this.Out)-1]
	return res
}

// 132模式----------------------------------------------------------------------------------------------------
// 思路：先获取到每个数的前最小数 13模式  然后用栈找2 倒序遍历 20ms/6mb
func Find132pattern(nums []int) bool {
	stack := make([]int, 0)
	n := len(nums)
	if n < 3 {
		return false
	}

	// 获取每个数的前面最小数
	minArr := make([]int, n)
	minNum := nums[0]
	minArr[0] = minNum
	for i := 1; i <= n-1; i++ {

		minArr[i] = min(minNum, nums[i])
		minNum = minArr[i]
	}

	for i := n - 1; i >= 0; i-- {
		// 当栈顶元素 比 当前数的前面最小数都小时 说明1>2了 不符合要求 不用存在于栈顶了
		for len(stack) > 0 && stack[len(stack)-1] <= minArr[i] {
			stack = stack[0 : len(stack)-1]
		}
		// 如果栈顶还有 说明 1<2 已经成立 如果正好3 > 2 就找到了
		if len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			return true
		}
		// 每次都 将2插入进入
		stack = append(stack, nums[i])
	}

	return false
}

// 字符串解码--------------------------------------------------------------------------------------------------
// 思路：一个数字栈，一个字符串栈  遇到] 取字符串栈中数据 直到[然后取数字栈来判定重复数 然后推入栈中 最后组合 0ms/2.2mb
func DecodeString(s string) string {
	if s == "" {
		return ""
	}

	stack := make([]string, 0)
	numStack := make([]int32, 0)
	cur, numCur := "", ""
	// 便利s
	for _, val := range s {

		switch val {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			numCur += string(val)
		case '[':
			if cur != "" {
				stack = append(stack, cur)
				cur = ""
			}
			n, _ := strconv.Atoi(numCur)
			numStack = append(numStack, int32(n))
			numCur = ""
			stack = append(stack, "[")
		case ']':
			for len(stack) > 0 && stack[len(stack)-1] != "[" {
				s := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				cur = s + cur
			}
			stack = stack[:len(stack)-1]
			n := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			a := ""
			for i := 0; i < int(n); i++ {
				a += cur
			}
			stack = append(stack, a)
			cur = ""
		default:
			cur += string(val)
		}

	}
	res := ""
	for _, val := range stack {
		res += val
	}
	res += cur
	return res
}

//

// 移掉K位数字-----------------------------------------------------------------------------------------------
// 思路： 贪心算法+栈  根据1个特性 如 132 右边的数<左边的数时 去掉左边的数就会使剩下的数小来移除 0ms/3.1mb
func removeKdigits(num string, k int) string {

	stack := make([]int32, 0)

	for _, val := range num {
		for len(stack) > 0 && k > 0 && val < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, val)
	}

	stack = stack[:len(stack)-k]
	res := strings.TrimLeft(string(stack), "0")
	if res == "" {
		return "0"
	}
	return res
}

// 标签验证器-----------------------------------------------------------------------------------------------------------
// 思路：就是单纯的栈 中间会设计到很多的判断 0ms/2.1mb
func IsValidTag(code string) bool {
	n := len(code)
	if n < 2 || code[0] != '<' || code[1] == '!' {
		return false
	}

	tagStack := make([]string, 0)
	s := ""
	for key, v := range code {

		// 遇到特殊处理将其内容置空
		if len(s) >= 11 && s[:9] == "<![CDATA[" && s[len(s)-2:] == "]]" && v == '>' {
			s = ""
			continue
		}

		if v == '<' && s == "" {
			s = "<"
			continue
		}

		if v == '>' {
			if s == "" {
				continue
			}
			s += ">"
			if checkTag(s) {
				if s[0:2] == "</" {
					if len(tagStack) > 0 && tagStack[len(tagStack)-1] == s[2:len(s)-1] {
						tagStack = tagStack[:len(tagStack)-1]
						s = ""
						// 开头 只能被 最后 干掉
						if key != n-1 && len(tagStack) == 0 {
							return false
						}
						continue
					} else {
						return false
					}
				} else if s[0:2] != "<!" {
					tagStack = append(tagStack, s[1:len(s)-1])
					s = ""
					continue
				}
			} else {
				return false
			}
		}

		if len(s) > 0 {
			s += string(v)
		}
	}

	if s == "" && len(tagStack) == 0 {
		return true
	}
	return false
}

func checkTag(s string) bool {
	if len(s) >= 9 && s[:9] == "<![CDATA[" {
		return true
	}
	if len(s) <= 2 {
		return false
	}

	count := 0
	for key, val := range s[1 : len(s)-1] {
		if val == '/' || key == 0 {
			continue
		}
		if val < 65 || val > 90 {
			return false
		}
		count++
	}
	if count > 9 {
		return false
	}

	return true
}

// 函数的独占时间--------------------------------------------------------------------------------------------------------
// 思路:因为是单线程 意思同一时间只能有一个函数被调用:一个函数被调用的时间= 结束时间-开始时间-中途调用其他函数消耗的时间 12ms/6.1mb
type CallStack struct {
	funcId    int
	startTime int
	extraTime int
}

func ExclusiveTime(n int, logs []string) []int {
	result := make([]int, n)
	stack := make([]CallStack, 0)

	for _, log := range logs {
		split := strings.Split(log, ":")
		if split[1] == "start" {
			id, _ := strconv.Atoi(split[0])
			num, _ := strconv.Atoi(split[2])
			stack = append(stack, CallStack{
				funcId:    id,
				startTime: num,
				extraTime: 0,
			})
		} else {
			es, _ := strconv.Atoi(split[2])
			cs := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			runtime := es - cs.startTime + 1 - cs.extraTime
			result[cs.funcId] += runtime
			if len(stack) > 0 {
				stack[len(stack)-1].extraTime += es - cs.startTime + 1
			}
		}
	}
	return result
}

// 棒球比赛 -----------------------------------------------------------------------------------------------------------
// 思路:简单的栈调用 C就将栈顶弹出 D就将栈顶*2插入 +就将顶1顶2加起来插入 最后求栈的和 4ms/2.6mb
func calPoints(ops []string) int {
	stack := make([]int, 0)

	for _, v := range ops {
		if v == "C" {
			stack = stack[:len(stack)-1]
			continue
		}
		if v == "D" {
			stack = append(stack, stack[len(stack)-1]*2)
			continue
		}

		if v == "+" {
			stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
			continue
		}
		n, _ := strconv.Atoi(v)
		stack = append(stack, n)
	}
	res := 0
	for _, v := range stack {
		res += v
	}

	return res
}

// 原子的数量 -----------------------------------------------------------------------------------------------------------
// 思路:
//func countOfAtoms(formula string) string {
//	 stack := make([]string,0)
//	 numMap := make(map[string]int,0)
//
//
//}

// 行星碰撞 -----------------------------------------------------------------------------------------------------------
// 思路: 遇到栈顶数据<0,自身又是>0的 说明会相撞将栈顶弹出用绝对值比较 >就将栈顶返还 =结束   <就继续撞   12ms/4.7mb
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0)

	for _, v := range asteroids {

		// 当栈顶与当前数+-不同时 进行Pk
		if len(stack) > 0 {
			for len(stack) > 0 {
				if checkDoublePN(stack[len(stack)-1], v) {
					stack = append(stack, v)
					break
				} else {
					n := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					absN := int(math.Abs(float64(n)))
					absV := int(math.Abs(float64(v)))
					if absN > absV {
						stack = append(stack, n)
						break
					} else if absN == absV {
						break
					} else if absN < absV && len(stack) == 0 {
						stack = append(stack, v)
						break
					}
				}
			}
		} else {
			stack = append(stack, v)
		}

	}
	return stack
}

// 是否不发生碰撞
func checkDoublePN(a, b int) bool {
	if a > 0 && b < 0 {
		return false
	}

	return true
}

// 每日温度-----------------------------------------------------------------------------------------------------------
// 思路：倒序遍历 当前值>栈顶就一直弹 直到找到<的情况 top-i 或者栈没值了 0  88ms/6.7mb
func dailyTemperatures(T []int) []int {
	n := len(T)
	stack := make([]int, 0)
	res := make([]int, n)

	for i := n - 1; i >= 0; i-- {

		// 找stack中比他大的数
		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			// 找到了就直接出去
			if T[cur] > T[i] {
				break
			}
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		}
		stack = append(stack, i)
	}
	return res
}

// 比较含退格的字符串----------------------------------------------------------------------------------------------------
// 思路: 两个栈  遍历S,T两个数 遇到#就把栈顶数据弹出 其余就追加 0ms/2mb
func backspaceCompare(S string, T string) bool {
	sStack := make([]int32, 0)
	tStack := make([]int32, 0)

	for _, v := range S {
		if v == '#' {
			if len(sStack) > 0 {
				sStack = sStack[:len(sStack)-1]
			}
			continue
		}
		sStack = append(sStack, v)
	}

	for _, v := range T {
		if v == '#' {
			if len(tStack) > 0 {
				tStack = tStack[:len(tStack)-1]
			}
			continue
		}
		tStack = append(tStack, v)
	}
	return string(sStack) == string(tStack)
}

// 括号的分数-----------------------------------------------------------------------------------------------------------
// 思路：只有()才是真正的值1 (() 会使1*2 当遇到() 只需要知道其前面生效的(有多少个 就让1移动多少位数(剩2) 0ms/2mb
func ScoreOfParentheses(S string) int {
	res, count := 0, uint(0)

	for k, v := range S {
		if v == '(' {
			count++
		} else {
			count--
			if S[k-1] == '(' {
				res += 1 << count
				fmt.Println(count, res)
			}
		}
	}
	return res
}

func XthPower() int {
	fmt.Println(uint(1) >> 3)
	return 1
}

// 索引处的解码字符串---------------------------------------------------------------------------------------------------
// 思路:有误!
func DecodeAtIndex(S string, K int) string {
	stack := make([]string, 0)
	cur := ""
	for _, val := range S {
		if val >= '2' && val <= '9' {
			if cur != "" {
				stack = append(stack, cur)
				cur = ""
			}
			str := stack[len(stack)-1]
			n := ""
			for i := 0; i < int(val)-48; i++ {
				n += str
			}
			// 替换栈顶
			stack[len(stack)-1] = n

		} else {
			cur += string(val)
		}
	}

	realS := ""
	for _, v := range stack {
		realS += v
	}
	k := K % len(realS)
	//unicode.IsDigit(c)
	return string(realS[k-1])
}

// 计算右侧小于当前元素的个数----------------------------------------------------------------------
// 思路：暴力 从右到左遍历 每个数都往右边找比他小的数计数 找到相等的就不用找了 直接加上其计数值 440ms/4.3mb
func CountSmaller(nums []int) []int {
	n := len(nums)

	res := make([]int, n)

	for i := n - 2; i >= 0; i-- {

		count := 0
		for k, v := range nums[i+1:] {
			if nums[i] > v {
				count++
			} else if nums[i] == v {
				count += res[k+i+1]
				break
			}
		}
		res[i] = count

	}
	return res
}

// 去除最外层的括号----------------------------------------------------------------------------
// 思路：左右括号一定会闭合，因此跳过每次最外层的左括号，让其他括号去配对，配对不上的就是跳过的左括号的右括号让他剔除
func removeOuterParentheses(S string) string {
	s := strings.Builder{}
	cnt := 0
	for _, v := range S {
		if v == '(' {
			if cnt > 0 {
				s.WriteRune(v)
			}
			cnt++
		} else {
			cnt--
			if cnt > 0 {
				s.WriteRune(v)
			}
		}
	}
	return s.String()
}

// 相邻字母不能相同------------------------------------------------------------------------------
// 思路：栈
func removeDuplicates(S string) string {
	r := make([]uint8, 0)

	for i := 0; i < len(S); i++ {
		if len(r) == 0 {
			r = append(r, S[i])
			continue
		}
		if S[i] == r[len(r)-1] {
			r = r[:len(r)-1]
		} else {
			r = append(r, S[i])
		}
	}
	return string(r)
}

// 子数组范围和--------------------------------------------------------------------------------
// 思路：暴力循环  固定左边界 移动右边界 每移动一步就计算一步
func subArrayRanges(nums []int) int64 {
	var ans int64
	n := len(nums)

	for i := 0; i < n; i++ {
		minNum, maxNum := nums[i], nums[i]

		for j := i + 1; j < n; j++ {
			minNum = min(minNum, nums[j])
			maxNum = max(maxNum, nums[j])
			ans += int64(maxNum - minNum)
		}

	}

	return ans
}

// 最大频率栈----------------------------------------------------------------------------------
// 思路：数字只会一个一个增加，用频率群组栈来保存
func ConstructorFreqStack() FreqStack {
	return FreqStack{map[int]int{}, map[int][]int{}, 0}
}

type FreqStack struct {
	freq    map[int]int   // 每个数字对应的出现频率
	group   map[int][]int // 频率群组
	maxFreq int           // 指向最大频率群组
}

func (f *FreqStack) Push(val int) {
	f.freq[val]++
	f.group[f.freq[val]] = append(f.group[f.freq[val]], val)
	f.maxFreq = max(f.maxFreq, f.freq[val])
}
func (f *FreqStack) Pop() int {
	res := f.group[f.maxFreq][len(f.group[f.maxFreq])-1]
	f.freq[res]--
	f.group[f.maxFreq] = f.group[f.maxFreq][:len(f.group[f.maxFreq])-1]
	if len(f.group[f.maxFreq]) == 0 {
		f.maxFreq--
	}
	return res
}

// 替换字符串中的括号内容
func evaluate(s string, knowledge [][]string) string {
	m := map[string]string{}
	for _, v := range knowledge {
		m[v[0]] = v[1]
	}

	ans := strings.Builder{}
	flag := false
	cur := -1
	for i, b := range s {
		switch b {
		case '(':
			cur = i
			flag = true
		case ')':
			flag = false
			v := m[s[cur+1:i]]
			if v == "" {
				v = "?"
			}
			ans.WriteString(v)
			cur = -1
		default:
			if !flag {
				ans.WriteRune(b)
			}
		}
	}

	return ans.String()
}

// 统计星号
func countAsterisks(s string) int {
	ans := 0
	open := 0
	for _, v := range s {
		if v == '|' {
			open ^= 1
		} else if v == '*' && open == 0 {
			ans++
		}
	}
	return ans
}

// 判断矩阵是否是一个 X 矩阵
func checkXMatrix(grid [][]int) bool {
	ans := true

	for x, in := range grid {
		for y, val := range in {
			if x == y || x+y == len(grid) {
				if val == 0 {
					ans = false
					break
				}
			} else {
				if val != 0 {
					ans = false
					break
				}
			}
		}
	}

	return ans
}
