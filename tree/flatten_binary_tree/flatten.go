package flatten_binary_tree

/*
114 二叉树展开链表 中等

这种可以用dfs 的 考虑一个最小单位就好，最小单位能跑完，直接dfs递归
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	// 逻辑
	// 如果左不为空，把左边移到右边
	if root.Left != nil {
		l := root.Left
		for l.Right != nil {
			// 在 左树的右数不为空的情况下，一路遍历到最右的叶子节点
			l = l.Right
		}
		// 叶子节点接上，右树
		// 按前序遍历，左树的最右叶子节点下一个就是右树
		l.Right = root.Right
		// 左树移到右边，将右边置空
		root.Right = root.Left
		root.Left = nil
	}
	flatten(root.Right)
}
