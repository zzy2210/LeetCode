package linkedlistcycle

/*
*
  - Definition for singly-linked list.

141 环形链表

久违了，我的链表 这题快慢指针秒了

快的走一次两步，慢的一次一步。 如果有环，那快的总会在某个时间追上/ 超过慢的 一圈 然后相遇
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	low, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		low = low.Next
		if fast == low {
			return true
		}
	}
	return false
}
