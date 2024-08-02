package copyrandomlist

/*

	138 随机链表复制

	这个链表的特点是除了NEXT 还有个随机指针，指向任一个点

	题解给了两种解法


	第一种我看懂了，就是遇到要新建的一律新建，然后保存在hash表里，如果遇到了已经保存了的情况，那就跳过
	就是 deepCopyByHash 拷贝当前节点，它的下一个与随机数。这样的话就可以一直递归到全是nil

	第二种方法是在 原节点后面接一个拷贝节点
	第一次把 next 拷贝好 也就是一个完整的 除了random都好了的
	第二次把random补充了
	第三次给它取出来，或则说串起来
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 跨两步，第一步是拷贝节点
	for node := head.Next; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}

	for node := head.Next; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}

	newHead := head.Next
	for node := head; node != nil; node = node.Next {
		newNode := node.Next
		node.Next = node.Next.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
	}
	return newHead
}

var nodeHash map[*Node]*Node

func copyRandomListByHash(head *Node) *Node {
	nodeHash = make(map[*Node]*Node)
	return deepCopByhash(head)
}

func deepCopByhash(node *Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := nodeHash[node]; ok {
		return v
	}
	newNode := &Node{Val: node.Val}
	nodeHash[node] = newNode
	newNode.Next = deepCopByhash(node.Next)
	newNode.Random = deepCopByhash(node.Random)
	return newNode
}
