package v1

func matrixScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	matrix := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i] = (matrix[i] << 1) + grid[i][j]
		}
	}

	invertRow := func(row int) {
		matrix[row] = ^matrix[row] & ((1 << n) - 1)
	}

	invertCol := func(col int) {
		for i := 0; i < m; i++ {
			matrix[i] ^= 1 << (n - col - 1)
		}
	}

	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			invertRow(i)
		}
	}

	for j := 1; j < n; j++ {
		ones := 0
		for i := 0; i < m; i++ {
			if matrix[i]&(1<<(n-j-1)) != 0 {
				ones++
			}
		}
		if ones+ones < m {
			invertCol(j)
		}
	}

	s := 0
	for i := 0; i < m; i++ {
		s += matrix[i]
	}
	return s
}
