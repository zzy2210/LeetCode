package invert_tree

/*
226  翻转二叉树 简单

有手就像 递归一下



*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return root

}

func invertTreeBySelf(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = invertTreeBySelf(root.Left)
	root.Right = invertTreeBySelf(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
