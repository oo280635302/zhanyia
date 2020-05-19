package program

import "fmt"

// 罗马数值转普通数字
// 枚举型
func romanToInt(s string) int {
	romans := [...]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	nums := [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	var start int
	var v int
	for i := len(romans) - 1; i >= 0; {
		charLen := len(romans[i])

		end := start + charLen
		fmt.Println(end)
		if end > len(s) || s[start:end] != romans[i] {
			i--
			continue
		}

		start += charLen
		v += nums[i]
	}
	return v
}

// 公共前缀 长度
// 取第一个 然后匹配
func longestCommonPrefix(arr []string) string {
	reply := ""
	if len(arr) < 1 {
		return reply
	} else if len(arr) == 1 {
		return arr[0]
	}

	for k, _ := range arr[0] {
		for k1, v := range arr[1:] {
			if len(v) > k && arr[0][k] == v[k] {
				if k1 == len(arr[1:])-1 {
					reply += string(arr[0][k])
				}
			} else {
				return reply
			}
		}
	}
	return reply
}

// 有效括号
// 用栈匹配 循环一次 将左括号都存起来 后存先匹配机制判断是否取完 - 同类都可以用栈知识
func isValid(s string) bool {
	if s == "" {
		return true
	}

	var stack []uint8

	m := map[uint8]uint8{
		'}': '{',
		')': '(',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			if m[s[i]] != stack[len(stack)-1] {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

// 合并链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	reply := &ListNode{}

	if l1 == nil && l2 == nil {
		return nil
	}

	nowNode := reply
	for {

		if l1 == nil && l2 == nil {
			nowNode = nil
			return reply
		}

		newNode := &ListNode{}

		if l1 == nil {
			nowNode.Val = l2.Val
			nowNode.Next = newNode
			l2 = l2.Next
		} else if l2 == nil {
			nowNode.Val = l1.Val
			nowNode.Next = newNode
			l1 = l1.Next
		} else {
			if l1.Val > l2.Val {
				nowNode.Val = l2.Val
				nowNode.Next = newNode
				l2 = l2.Next
			} else {
				nowNode.Val = l1.Val
				nowNode.Next = newNode
				l1 = l1.Next
			}
		}
		if l1 == nil && l2 == nil {
			nowNode.Next = nil
			return reply
		}
		nowNode = nowNode.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}
