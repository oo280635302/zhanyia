package program

import "fmt"

func Ingress() {
	fmt.Println(mergeInBetween(&ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}, 1, 2, &ListNode{Val: 991, Next: &ListNode{Val: 992, Next: &ListNode{Val: 993}}}))
}
