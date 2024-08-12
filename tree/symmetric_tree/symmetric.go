package symmetric_tree

/*
101 对称二叉树 中等

检查轴对称

写一个查两个值是否相等的函数， 里面做比较，返回的时候再把两个值的 轴对称节点放进去比较，这样递归完事
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	// 走到这一层 一定不是全空
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
