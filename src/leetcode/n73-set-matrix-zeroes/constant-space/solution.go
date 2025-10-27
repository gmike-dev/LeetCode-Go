package constant_space

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	zeroInFirstCol := false
	zeroInFirstRow := false
	for i := range n {
		if matrix[i][0] == 0 {
			zeroInFirstCol = true
			break
		}
	}
	for j := range m {
		if matrix[0][j] == 0 {
			zeroInFirstRow = true
			break
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if zeroInFirstCol {
		for i := range n {
			matrix[i][0] = 0
		}
	}
	if zeroInFirstRow {
		for j := range m {
			matrix[0][j] = 0
		}
	}
}
