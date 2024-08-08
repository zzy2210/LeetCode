package maxdeep

/*
104 二叉树最大深度 简单

递归了一下

一路下到最低，返回0 然后往上一次+1 在往上就是送最小值 一路送上去

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
