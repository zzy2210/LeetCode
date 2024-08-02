package mergetwolists

/*

21. 合并两个有序链表

因为是升序链表，就每次都拿两边node

不需要纠结和原本的比，直接拿两个中的最小的往里放就ok了，后续的一定都比这个大

这题应该还能变式

一个是我经常写的 合并k个有序链表

解法就是归并，把k个拆成 k/2 组，两两合并上来

一个是 合并 k个无序链表。 还要加排序


*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dump := &ListNode{}
	head := dump
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			dump.Next = list1
			list1 = list1.Next
		} else {
			dump.Next = list2
			list2 = list2.Next
		}
		dump = dump.Next
	}
	if list1 != nil {
		dump.Next = list1
	} else if list2 != nil {
		dump.Next = list2
	}
	return head.Next
}
