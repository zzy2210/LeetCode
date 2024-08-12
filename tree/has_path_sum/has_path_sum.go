package has_path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == root.Right { // 这两个是指针，相等只有nil的情况
		return targetSum == 0
	}
	return hasPathSum(root.Right, targetSum) || hasPathSum(root.Left, targetSum)
}
