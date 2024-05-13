package v1

func matrixScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	s := 0
	for j := 0; j < n; j++ {
		ones := 0
		for i := 0; i < m; i++ {
			if grid[i][j] == grid[i][0] {
				ones++
			}
		}
		s += max(ones, m-ones) * (1 << (n - j - 1))
	}
	return s
}
