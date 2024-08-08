package sametree

/*
100 相同的树 简单
这个也是递归 每次都是递归它的 左右节点

当 相等的时候没必要返回正确，直接走到下一个就ok
如果一路相等总会走到 空的

如果路上有不等 就返回了
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
