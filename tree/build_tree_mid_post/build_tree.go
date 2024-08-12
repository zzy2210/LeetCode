package build_tree_mid_post

import "slices"

/*
106 从中序与后序构造二叉树 中等
先拿根节点，根节点是后序的最后一个

这里 左子树 右子树是  中序的左右 都好判断
后序呢 后序的左边一定长度就是左子树 后边是右子树? 还真是
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	leftSize := slices.Index(inorder, postorder[n-1])
	left := buildTree(inorder[:leftSize], postorder[:leftSize])
	right := buildTree(inorder[1+leftSize:], postorder[leftSize:n-1])
	return &TreeNode{Val: postorder[n-1], Left: left, Right: right}
}
