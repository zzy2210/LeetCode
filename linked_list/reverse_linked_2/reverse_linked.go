package reverselinked2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: head}
	count := 0
	n1, n2 := dummyHead, dummyHead

	for count < left-1 {
		n1 = n1.Next
		n2 = n2.Next
		count++
	}
	for count < right {
		n2 = n2.Next
		count++
	}
	dummy2 := &ListNode{Next: n2.Next}
	// 断开
	n2.Next = nil

	resverNode := n1.Next
	for resverNode != nil {
		tmp := resverNode.Next
		resverNode.Next = dummy2.Next
		dummy2.Next = resverNode
		resverNode = tmp
	}

	n1.Next = dummy2.Next

	return dummyHead.Next
}
