package program

import (
	"fmt"
)

func Ingress() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: -5,
			},
		},
		Right: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: 4,
			},
		},
	}
	res := sufficientSubset(node, -1)
	fmt.Println(res)
}
