package zigzag_level_order

/*
103 二叉树锯齿层次遍历 中等

就是正常层次遍历
如果是奇数层就反着来，偶数层正着来

这里用队列比较方便
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	p := []*TreeNode{root}
	for len(p) > 0 {
		n := len(p)
		tmpSlice := make([]int, n)
		for i := range tmpSlice {
			node := p[0]
			p = p[1:]
			if len(res)%2 == 1 {
				tmpSlice[n-i-1] = node.Val
			} else {
				tmpSlice[i] = node.Val
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		res = append(res, tmpSlice)
	}
	return res
}
