package matrix

// 36. 有效的数独

func isValidSudoku(board [][]byte) bool {
	// 先计算行
	// 再计算列
	// 再计矩阵

	// 第一个代表第几行/列,第二个代表是否出现过,不需要序列,因为数组的序列就是key
	rows := [9][9]int{}
	lists := [9][9]int{}

	// 大横坐标 大纵坐标
	// 3x3 的 九个矩阵,根据 小的 x y 坐标可以算出对应再某个 矩阵
	// 比如 一个数是 2,6 那么他就在 (0,2) 这个矩形上
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, v := range row {
			if v == '.' {
				continue
			}

			// 不做这一步会走是 ascii 码-1
			key := v - '1'

			// 判断行
			if rows[i][key] == 0 {
				rows[i][key] = 1
			} else {
				return false
			}
			// 判断列
			if lists[j][key] == 0 {
				lists[j][key] = 1
			} else {
				return false
			}

			if subboxes[i/3][j/3][key] == 0 {
				subboxes[i/3][j/3][key] = 1
			} else {
				return false
			}
		}
	}

	return true
}
