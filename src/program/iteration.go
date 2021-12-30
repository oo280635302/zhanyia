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

// 扰乱字符串-----------------------------------------------------------------------------------------------------------
// 思路:递归, 将数切割成 1个数和2个数时匹配  切割时:头和头匹配, 头和尾匹配 0ms/2.1mb
func isScramble(s1 string, s2 string) bool {
	length := len(s1)
	if !checkStrSame(s1, s2) {
		return false
	}
	if length == 1 {
		return s1[0] == s2[0]
	}
	if length == 2 {
		return s1[0] == s2[0] && s1[1] == s2[1] || s1[0] == s2[1] && s1[1] == s2[0]
	}

	for i := 1; i < len(s1); i++ {
		la, ra := s1[:i], s1[i:]
		lb, rb := s2[:i], s2[i:]
		nb, mb := s2[length-i:], s2[:length-i]
		// 将s2 切割成 头和尾模式 只要s1的头和s2的头或者s2的尾配对就 = 匹配成功
		if isScramble(la, lb) && isScramble(ra, rb) || isScramble(la, nb) && isScramble(ra, mb) {
			return true
		}
	}
	return false
}

func checkStrSame(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var ta, tb [26]int
	for i := 0; i < len(a); i++ {
		ta[a[i]-'a']++
		tb[b[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if ta[i] != tb[i] {
			return false
		}
	}
	return true
}

// 斐波那契数列---------------------------------------------------------------------------------------------------------
func fibonacci(x int) int {
	if x == 1 {
		return 0
	}
	if x == 2 || x == 3 {
		return 1
	}
	n, nf1, nf2 := 0, 1, 1
	for i := 3; i < x; i++ {
		n = nf1 + nf2 // 当前数 = 前1+前2
		nf2 = nf1     // 前2 = 前1
		nf1 = n       // 前1 = 当前数
	}
	return n
}
