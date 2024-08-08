package minbst

import "math"

/*
530 二叉搜索树的最小绝对差 简单

给了一个二叉搜索树，算最小差

中序遍历的 二叉搜索树 是一个递增序列

所以
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt
	res := []int{}
	geinorderTraversal(root, &res)
	for i := 1; i < len(res); i++ {
		tem := res[i] - res[i-1]
		if tem < min {
			min = tem
		}
	}
	return min

}

func geinorderTraversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	geinorderTraversal(node.Left, res)
	*res = append(*res, node.Val)
	geinorderTraversal(node.Right, res)
}
