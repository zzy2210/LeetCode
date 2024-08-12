package right_side

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, deep int) {
		if node == nil {
			return
		}
		if deep == len(ans) {
			ans = append(ans, node.Val)
		}
		dfs(node.Right, deep+1)
		dfs(node.Left, deep+1)
	}
	dfs(root, 0)
	return ans
}
