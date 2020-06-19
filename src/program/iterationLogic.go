package program

// 迭代算法
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从先序遍历还原二叉树 ---------------------------------------
// 特征:有了左支点才有右支点 有了上支点才有下支点 使用辅助指针 用于判断支点高度
func RecoverFromPreorder(S string) *TreeNode {
	// path 数组树
	// pos 指向当前操作的索引
	path, pos := []*TreeNode{}, 0
	for pos < len(S) {
		level := 0
		// 根据-数量找到 支链等级
		for S[pos] == '-' {
			level++
			pos++
		}

		// 获取数字 -
		value := 0
		for ; pos < len(S) && S[pos] >= '0' && S[pos] <= '9'; pos++ {
			value = value*10 + int(S[pos]-'0')
		}

		// 获取当前支点信息
		node := &TreeNode{Val: value}

		// 当 当前等级  与 位置正好相等时是新增的左支点(根节点除外) - 相反就为右节点
		// 插入时找到 其父节点 - 左节点-父节点的位置为其前一个索引位置并按 0 - 1 - 2 -3的等级分布
		// 右节点位置 - 其父节点的位置在level-1 -- 同时说明其父节点的左节点已充填完 将等级上升到父节点位置插入,方便左节点的定位
		if level == len(path) {
			if len(path) > 0 {
				path[len(path)-1].Left = node
			}
		} else {
			path = path[:level]
			path[len(path)-1].Right = node
		}
		path = append(path, node)
		//for _,v:= range path{
		//	fmt.Print(v.Val," ")
		//}
		//fmt.Println()
	}

	return path[0]
}
