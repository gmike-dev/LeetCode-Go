package v1

func setZeroes(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	zeroRow := make([]bool, n)
	zeroCol := make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				zeroRow[i] = true
				zeroCol[j] = true
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if zeroRow[i] || zeroCol[j] {
				matrix[i][j] = 0
			}
		}
	}
}
