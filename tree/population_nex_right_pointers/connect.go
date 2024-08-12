package population_nex_right_pointers

/*
117 填充每个节点的下一节点 中等

给每个节点多了一个指针，指向它的右节点

如果右边没有就指向nil

dfs一下， 把每个的节点放到数组里面
a[i].new =node
a[i] = node
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	pre := []*Node{}
	var dfs func(*Node, int)
	dfs = func(node *Node, deepth int) {
		if node == nil {
			return
		}
		if deepth == len(pre) {
			pre = append(pre, node)
		} else {
			pre[deepth].Next = node
			pre[deepth] = node
		}
		dfs(node.Left, deepth+1)
		dfs(node.Right, deepth+1)
	}
	dfs(root, 0)
	return root
}
