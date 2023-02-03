package program

import (
	"math/rand"
	"sort"
)

// 删除当前节点 -------------------------------------------------------------------------------------------------------
func deleteNode(node *ListNode) {
	if node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}

// 删除链表的倒数第N个节点 --------------------------------------------------------------------------------------------
// 思路:双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := head
	r := head
	for i := 0; i < n; i++ {
		r = r.Next
	}
	if r == nil {
		head = head.Next
		return head
	}

	for {
		if r.Next != nil {
			l = l.Next
			r = r.Next
		} else {
			l.Next = l.Next.Next
			break
		}
	}

	return head
}

// 反转链表------------------------------------------------------------------------------------------------------------
// 思路1：递归  因为链表特性当前链表只能获取下一个数不能获取上一个数，所以需要用递归特性保存到每层链表的值，然后从最后两个数开始交换，交换完后返回上一层继续交换
func reverseList(head *ListNode) *ListNode {
	// 遇到nil就等于到底了 返回
	if head == nil || head.Next == nil {
		return head
	}

	// tmp 指的当前函数的头的下个指向
	tmp := head.Next

	// 翻转第一步： next后的链表切分开
	reverse := reverseList(tmp)

	// 翻转第二步：将头赋到尾  这里的tmp是reverse的尾指针
	tmp.Next = head

	// 反转第三步：将尾清理干净
	head.Next = nil

	return reverse
}

// 思路2：栈
func reverseListByStack(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	res := &ListNode{}
	tmp := res

	for i := len(stack) - 1; i >= 0; i-- {
		tmp.Val = stack[i]
		if i > 0 {
			tmp.Next = &ListNode{}
		}
		tmp = tmp.Next

	}
	return res
}

// 思路3：迭代 辅助链表保存head除头的部分
func reverseListBySup(head *ListNode) *ListNode {
	var res *ListNode

	for head != nil {
		// 临时表保存head除头
		tmp := head.Next

		// 拔head头接res尾
		head.Next = res

		// 这个时候的head其实是res
		res = head

		// 将临时表的head剩余部分还给head
		head = tmp
	}

	return res
}

// 回文链表-----------------------------------------------------------------------------------------------------
// 思路：双指针 从中间比较  考虑到链表只能往后移动，
func isPalindromeByList(head *ListNode) bool {
	// 第一步：找中间点偏右点
	// 123的3 1234的3
	var l, r = head, head

	// 第二步：l、r从1开始跳，l跳一步，r跳两步，直到r到底
	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next.Next
	}

	// 如果r还有数，说明链表是奇数，l往下移动一位
	if r != nil {
		l = l.Next
	}

	// 第三步: 反转l后的数据,r从头开始
	r = head
	l = reverseList(l)

	// 第四步：依次对比数据
	for l != nil {
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}

	return true
}

// 环形链表-------------------------------------------------------------------------------------------------------
// 思路：双指针
func hasCycle(head *ListNode) bool {
	l, r := head, head

	for r != nil && r.Next != nil {
		r = r.Next
		if l == r {
			return true
		}
		if r.Next == nil {
			break
		}
		l = l.Next
		r = r.Next
	}

	return false
}

// 合并K个升序链表
// 思路： 排序，再组合   有更优解：分治2个链表2个链表进行合并
func mergeKLists(lists []*ListNode) *ListNode {
	arr := make([]int, 0)

	for _, list := range lists {
		for list != nil {
			arr = append(arr, list.Val)
			list = list.Next
		}
	}

	if len(arr) == 0 {
		return nil
	}

	sort.Ints(arr)

	res := &ListNode{Val: arr[0]}
	cur := res

	for _, v := range arr[1:] {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	return res
}

// 两两交换链表中的节点
// 思路：链表交换
func swapPairs(head *ListNode) *ListNode {
	tmp := &ListNode{Val: 0, Next: head} // 来个辅助头

	cur := tmp
	for cur.Next != nil && cur.Next.Next != nil { // 只有两个数的时候才交换
		n1 := cur.Next
		n2 := cur.Next.Next
		cur.Next = n2     // 先让2前移
		n1.Next = n2.Next // 再让1来接管2的小弟
		n2.Next = n1      // 再让1成为2的小弟
		cur = n1.Next     // 当前点往后移动2个身位
	}

	return tmp.Next
}

// 链表随机节点
// 思路：抽样水塘， 时间换空间，每次弹出都是o(n)
type solution struct {
	List *ListNode
}

func constructor(head *ListNode) solution {
	return solution{head}
}

func (this *solution) getRandom() int {
	i := 1
	ans := 0
	for tmp := this.List; tmp != nil; tmp = tmp.Next {
		if rand.Intn(1) == 0 {
			ans = tmp.Val
		}
		i++
	}
	return ans
}

// 合并两个链表
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	list := list1
	var head, tail, last *ListNode
	var cnt = 0
	for list != nil {
		if cnt == a {
			head = last
		}
		if cnt == b {
			tail = list.Next
			break
		}
		cnt++
		last = list
		list = list.Next
	}
	head.Next = list2

	for cur := list2; cur != nil; cur = cur.Next {
		if cur.Next == nil {
			cur.Next = tail
			break
		}
	}

	return list1
}
