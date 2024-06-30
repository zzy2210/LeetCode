package main

func main() {

}

// 沿着 上,下左,右四个方向扫描
// 每次扫描完一类就移动一次对应的边界
// 这何尝不是一种滑动窗口
func spiralOrder(matrix [][]int) []int {
	var end []int

	// 定义矩阵的 上下,左右边界
	top, left := 0, 0
	bottom := len(matrix) - 1
	right := len(matrix[0]) - 1

	// 遍历扫描
	for top <= bottom && left <= right {
		// 扫描上横
		for i := left; i <= right; i++ {
			end = append(end, matrix[top][i])
		}
		// 扫描完成,上限下移
		top++
		// 扫描右竖
		for i := top; i <= bottom; i++ {
			end = append(end, matrix[i][right])
		}
		right--
		// 做一次界限判断
		if top > bottom || left > right {
			break
		}
		// 扫描下横
		for i := right; i >= left; i-- {
			end = append(end, matrix[bottom][i])
		}
		bottom--
		// 左竖
		for i := bottom; i >= top; i-- {
			end = append(end, matrix[i][left])
		}
		left++
	}
	return end
}
