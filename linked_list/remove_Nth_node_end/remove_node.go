package removenthnodeend

/*
19.删除链表倒数第N个节点
快慢指针
快指针比慢快 n+1 步
快指针为空时，慢指针就是倒数第 n-1 个节点

考虑到  只有1 删1 的情况，需要有一个头节点索引，让 slow.Next = head
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := dummy.Next, dummy

	for fast != nil {
		if n > 0 {
			fast = fast.Next
			n--
		} else {
			fast = fast.Next
			slow = slow.Next
		}
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	} else {
		slow.Next = nil
	}

	return dummy.Next
}
