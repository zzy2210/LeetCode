package level_order

/*

102 二叉树层序遍历

第一个解法是双数组
第一个数组表达当前层的所有节点，第二个数组表达下一层的全部节点

队列的方法就是将戴白哦当前层数的数组改成队列
但是感觉意义不高
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	// q 是 当前层数的全部节点
	q := []*TreeNode{root}
	// i 是层数
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		// 临时变量，代表下一层，结尾给替换上
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}

func levelOrderByQueue(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	// q 是 全部节点按顺序放进的队列
	q := []*TreeNode{root}
	// 不断从队列取，直到取完
	for len(q) > 0 {
		n := len(q)
		tmpSlice := make([]int, n)
		for i := range tmpSlice {
			// pop
			node := q[0]
			q = q[1:]
			tmpSlice[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, tmpSlice)
	}

	return res
}
