package addtwonumbers

/*
	2. 两数相加

	不是返回结果值，是返回一个链表，每个 node 对应每个位数的结果

	加个进位数就ok


	普遍情况：直接遍历 l1 l2

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	head := result
	// 进位标志
	carry := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		// 考虑长度不等
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		sum := v1 + v2 + carry
		// 1x /10 = 1  x / 10 = 0 通过数学判断 carry > 10 ?
		carry = sum / 10
		sum = sum % 10
		result.Next = &ListNode{Val: sum}
		result = result.Next
	}
	if carry == 1 {
		result.Next = &ListNode{Val: carry}
	}

	return head.Next
}
