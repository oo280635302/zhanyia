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
func NumTrees(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		fmt.Println(c, i)
		c = c * (4*i + 2) / (i + 2)
	}
	return c
}
