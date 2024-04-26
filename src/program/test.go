package program

import "fmt"

func Ingress() {
	fmt.Println(amountOfTime(
		&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}, 3))
}
