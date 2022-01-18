package program

import "fmt"

// 将有序数组转换为二叉搜索树-------------------------------------------------------------------------------------------
// 思路：递归 每次都从正中间切开	4ms/4.4mb
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:]),
	}
}

func sortedArrayToBSTBy2(nums []int) *TreeNode {
	return array2treeRecur(nums, 0, len(nums)-1)
}

func array2treeRecur(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	tree := &TreeNode{Val: nums[mid]}
	tree.Left = array2treeRecur(nums, left, mid-1)
	tree.Right = array2treeRecur(nums, mid+1, right)

	return tree
}

// 路径总和-----------------------------------------------------------------------------------------------------------
// 思路1: 递归 操作找到每个数的子节点上判断是否相等   8ms/4.8mb
func hasPathSum(root *TreeNode, sum int) bool {
	return checkPath(root, 0, sum)
}

func checkPath(root *TreeNode, nowSum, sum int) bool {
	tSum := root.Val + nowSum
	if root.Left == nil && root.Right == nil {
		return tSum == sum
	}
	a, b := false, false
	if root.Left != nil && tSum < sum {
		a = checkPath(root.Left, tSum, sum)
	}
	if root.Right != nil && tSum < sum {
		b = checkPath(root.Right, tSum, sum)
	}
	return a || b
}

// 更优解:迭代  8ms/4.8mb
func hasPathSum2(root *TreeNode, sum int) bool {
	if root == nil { // 如果树为空
		return false // 返回false
	}

	var stack []*TreeNode            // 保存树节点的栈
	var sumStack []int               // 保存每个节点的和
	stack = append(stack, root)      // 树的根节点入栈
	sumStack = append(sumStack, sum) // 初始的和sum入栈

	for len(stack) > 0 {
		// 出栈
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		s := sumStack[len(sumStack)-1]
		sumStack = sumStack[:len(sumStack)-1]

		// 如果当前节点已经为叶子节点
		if n.Left == nil && n.Right == nil && n.Val == s {
			return true
		}

		if n.Left != nil {
			stack = append(stack, n.Left)
			sumStack = append(sumStack, s-n.Val)
		}

		if n.Right != nil {
			stack = append(stack, n.Right)
			sumStack = append(sumStack, s-n.Val)
		}
	}

	return false
}

// 不同的二叉搜索树-------------------------------------------------------------------------------------------------
// 思路: 动态规划 每个数都与之前数有关 如 2 = 0*1 + 1*0 以此推论每次都保存前一个数 0ms/1.9mb
func numTrees1(n int) int {
	if n <= 1 {
		return 1
	}
	r := make([]int, n+1)
	r[0], r[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			// 获取 Ci = C0*Ci-1 + ...+ Ci-1*C0
			r[i] += r[j] * r[i-j-1]
		}
	}
	fmt.Println(r)
	return r[n]
}

// 更优解：类似于 (数学问题1.卡特兰数) -> 变式 0ms/1.9mb
func NumTrees1(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		fmt.Println(c, i)
		c = c * (4*i + 2) / (i + 2)
	}
	return c
}

// 不同的二叉搜索树2---------------------------------------------------------------------------------------------------
// 思路：卡特兰树 一个节点为顶点时,所有变化次数 = dfs左数变化次数 * dfs右数变化次数  特殊情况:在边界上时有一段会出现s>e的情况 返回有一个nil的数组以便计算 4ms/4.5mb
func GenerateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreesDfs(1, n)
}

func generateTreesDfs(start, end int) []*TreeNode {
	res := make([]*TreeNode, 0)
	// 当数据s>e = 返回有nil的数组
	if start > end {
		res = append(res, nil)
		return res
	}
	// 遇到相同的之间返回 避免继续递归 浪费时间/空间
	if start == end {
		res = append(res, &TreeNode{Val: start})
		return res
	}

	// 其他情况 找当前s->e的所有节点 当一个数的固定时,其所有变化次数 = 左数变化次数 * 右数变化次数
	for i := start; i <= end; i++ {
		for _, v := range generateTreesDfs(start, i-1) {
			for _, v1 := range generateTreesDfs(i+1, end) {
				oneTree := &TreeNode{
					Val:   i,
					Left:  v,
					Right: v1,
				}
				res = append(res, oneTree)
			}
		}
	}
	return res
}

// 树的深度-------------------------------------------------------------------------------------------------------------
// 思路1：广度搜索 将每层节点统计出来 层层递进 算出最终层数
func maxDepthByBreadth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur := []*TreeNode{root}
	ds := 0

	for {
		if len(cur) == 0 {
			break
		}

		node := make([]*TreeNode, 0)
		for _, v := range cur {
			if v.Left != nil {
				node = append(node, v.Left)
			}
			if v.Right != nil {
				node = append(node, v.Right)
			}
		}
		cur = node

		ds++
	}

	return ds
}

// 思路2：深度优先
func maxDepthByDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepthByDepth(root.Left), maxDepthByDepth(root.Right)) + 1
}

// 二叉树的层序遍历------------------------------------------------------------------------------------------------------
// 思路：广度优先 一层一层遍历出来
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	lv := make([][]int, 0)
	//lv = append(lv,[]int{root.Val})
	cur := []*TreeNode{root}

	for len(cur) > 0 {
		node := make([]*TreeNode, 0)
		curInt := make([]int, 0)

		for _, v := range cur {
			if v.Left != nil {
				node = append(node, v.Left)
			}
			if v.Right != nil {
				node = append(node, v.Right)
			}
			curInt = append(curInt, v.Val)
		}
		cur = node
		lv = append(lv, curInt)
	}

	return lv
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

// 二叉树的中序遍历 左->中->右---------------------------------------------------------------------------
// 解法一：迭代+stack的方法 0ms/2mb
// 中序遍历应用场景:可以用来做表达式树，在编译器底层实现的时候用户可以实现基本的加减乘除，比如 a*b+c
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

// 解法三：morris 创建 一个连接 将树中 所有父节点的子节点里的最右节点指向父节点
//     然后从最左节点开始进组，左进组了父进组，父进组了右进组，右进组了他指向父节点，然后父进组，又右进组,依次循环
func inOrderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil {

			// 找最右点
			pre := root.Left
			for pre.Right != nil && pre.Right != root {
				pre = pre.Right
			}

			// 只有pre已经组成循环的时候才会=nil，将他的指针指向root
			if pre.Right == nil {
				pre.Right = root
				root = root.Left

				// 当他!=nil,说明是其之前的右节点都已经被处理了指向，同时到达了最左点
			} else {
				res = append(res, root.Val)
				pre.Right = nil
				root = root.Right
			}

		} else {
			// 左边已经被处理完了，处理右边
			res = append(res, root.Val)
			// 所有左都进组了，进右组，或者从右跳到其父节点
			root = root.Right
		}
	}
	return res
}

// 二叉树的前序遍历 中->左->右 -------------------------------------------------------------------------------------
// 在二叉树的中序遍历基础上修改 -- 栈放每次支点的右子节点 支点val每次都放入结果中,同时每次都往左追下一格 0ms/2mb\
// 方法一:迭代
// 前序遍历的应用场景:可以用来实现目录结构的显示
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

// 解法二:递归 - 中左右 完成 0ms/2mb
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
// 后序遍历的应用场景: 计算目录内的文件占用的数据大小
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

// 解法二:递归 0ms/2mb
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

// 验证二叉搜索树-------------------------------------------------------------------------------------------------------
// 思路：中序遍历，判断是否返回结果是递增的
func isValidBST(root *TreeNode) bool {
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
		if len(reply) > 0 && reply[len(reply)-1] >= stack[index].Val {
			return false
		}
		reply = append(reply, stack[index].Val)
		root = stack[index].Right //右节点会进入下次循环，如果 =nil，继续出栈
		stack = stack[:index]     //出栈
	}

	return true
}

// 判断二叉树是否对称----------------------------------------------------------------------------------------------------
// 思路：递归，
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricRecur(root.Left, root.Right)
}

func isSymmetricRecur(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricRecur(left.Left, right.Right) && isSymmetricRecur(left.Right, right.Left)
}

// 单词接龙 II  从wordList里面找到beginWord到endWord最短的变化路线
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	graph := map[string]map[string]bool{}
	wordMap := map[string]bool{}
	for _, v := range wordList {
		wordMap[v] = true
	}

	// 字典里面没有endWord 直接GG
	if !wordMap[endWord] {
		return [][]string{}
	}

	//1.广度优先遍历创图
	queue := []string{beginWord}
	layer := 0
	wordMap[beginWord] = false
	isFound := false
	for len(queue) > 0 {
		layer++
		newQueue := []string{}

		// 层序遍历
		for _, qv := range queue {
			s := []byte(qv)
			canGo := map[string]bool{}

			// 依次变化单个字符
			for i, v := range s {
				// 只存在小写字母
				for c := 'a'; c <= 'z'; c++ {

					s[i] = byte(c)
					newWord := string(s)
					// 如果正好找到endWord那就可以结束层次遍历了
					if newWord == endWord {
						isFound = true
					}
					// 找到了可以变化的单词
					if wordMap[newWord] {
						canGo[newWord] = true
						newQueue = append(newQueue, newWord)
					}
				}
				// 将变化了的字符变回去
				s[i] = v
			}

			// 该层能变化到的单词都加入该单词的图下
			graph[qv] = canGo
		}

		// 如果找到endWord 直接结束
		if isFound {
			break
		}

		// 去掉已遍历的单词
		for _, vi := range newQueue {
			wordMap[vi] = false
		}
		queue = newQueue
	}

	//2.深度优先遍历寻找路径
	res := make([][]string, 0)
	var dfs = func(beg string, num int) {}
	path := []string{beginWord}

	dfs = func(beg string, num int) {
		if num == layer {
			// 找到了
			if beg == endWord {
				// 拷贝 不能直接把Path加入结果中 因为回溯时Path会变化
				dst := make([]string, len(path))
				copy(dst, path)
				res = append(res, dst)
			}
			return
		}
		for k, _ := range graph[beg] {

			// 将当前路径插入
			path = append(path, k)
			// 递归找
			dfs(k, num+1)
			// 找到了就把之前的路径清理掉
			path = path[:len(path)-1]
		}
	}

	dfs(beginWord, 0)

	return res
}
