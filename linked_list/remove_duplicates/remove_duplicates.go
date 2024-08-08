package removeduplicates

/*
82. 删除排序链表重复元素 2

最简单的一集 已经排序了，直接删

哦 还是有问题的 是把本身也删了

那就是拿下一个和下下个 如果相同了就写一个循环一直删下去

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	node := dummyHead
	for node.Next != nil && node.Next.Next != nil {
		if node.Next.Val == node.Next.Next.Val {
			v := node.Next.Val
			for node.Next != nil && v == node.Next.Val {
				node.Next = node.Next.Next
			}
		} else {
			node = node.Next
		}
	}
	return dummyHead.Next
}
