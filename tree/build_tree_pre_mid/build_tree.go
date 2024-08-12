package build_tree_pre_mid

import "slices"

/*
105 从前序与中序构造二叉树

递归制作，一路递归到叶子节点构造它，然后一路逆上来
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	// 拿根节点在 中序遍历中的地方，这样左边都是左子树，右边都是右子树
	leftSize := slices.Index(inorder, preorder[0])
	left := buildTree(preorder[1:1+leftSize], inorder[:leftSize])
	right := buildTree(preorder[1+leftSize:], inorder[1+leftSize:])
	return &TreeNode{preorder[0], left, right}
}
