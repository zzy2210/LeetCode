package validparentheses

func main() {

}

func setZeroes(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, rows := range matrix {
		for j, v := range rows {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i, rows := range matrix {
		for j := range rows {
			if row[i] || col[j] {
				rows[j] = 0
			}
		}
	}
}
