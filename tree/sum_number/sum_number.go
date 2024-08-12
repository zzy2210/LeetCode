package sum_number

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

// x 为上一次计算的结果
func dfs(root *TreeNode, x int) int {
	if root == nil {
		return 0
	}
	x = x*10 + root.Val
	if root.Left == root.Right {
		return x
	}
	return dfs(root.Left, x) + dfs(root.Right, x)
}
