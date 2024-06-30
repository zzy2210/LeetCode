package main

import "fmt"

func main() {
	//[[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
	a := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(a)
	fmt.Println(a)

}

/*
循环每个值
如果一个数 值 从0 -> 1 设为2,这样其他数就看到他就将它视为0
反之则设为-1

假设一个中间的数,它的上一个则是  [i-1][j]
右边是 [i][j+1]
下面是 [i+1][j]
左边是 [i][j-1]

就tm邪门,还有斜的
*/
func gameOfLife(board [][]int) {
	for i, row := range board {
		for j, v := range row {
			// 上下左右的缓存值
			var a, b, c, d int
			var e, f, g, h int
			// 上
			if i > 0 {
				a = board[i-1][j]
			}

			// 右
			if j < len(row)-1 {
				b = board[i][j+1]
			}

			// 下
			if i < len(board)-1 {
				c = board[i+1][j]

			}
			// 左
			if j > 0 {
				d = board[i][j-1]
			}

			// 左上
			if i > 0 && j > 0 {
				e = board[i-1][j-1]
			}
			// 右上
			if j < len(row)-1 && i > 0 {
				f = board[i-1][j+1]
			}
			// 左下
			if j > 0 && i < len(board)-1 {
				g = board[i+1][j-1]
			}
			// 右下
			if j < len(row)-1 && i < len(board)-1 {
				h = board[i+1][j+1]
			}
			a, b, c, d, e, f, g, h = getReal(a), getReal(b), getReal(c), getReal(d), getReal(e), getReal(f), getReal(g), getReal(h)
			total := a + b + c + d + e + f + g + h
			if total < 2 || total > 3 {
				if v == 1 {
					board[i][j] = -1
				}
			}
			if total == 3 {
				if v == 0 {
					board[i][j] = 2
				}
			}

		}
	}

	for i, row := range board {
		for j, v := range row {
			if v == -1 {
				board[i][j] = 0
			} else if v == 2 {
				board[i][j] = 1
			}
		}
	}
}

func getReal(v int) int {
	if v == -1 {
		return 1
	} else if v == 2 {
		return 0
	}
	return v
}
