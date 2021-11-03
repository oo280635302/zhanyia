package program

import "fmt"

func Ingress() {
	a := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	isPalindromeByList(a)

	for a != nil {
		fmt.Println(a)
		a = a.Next
	}
}

// 中序遍历 morris
