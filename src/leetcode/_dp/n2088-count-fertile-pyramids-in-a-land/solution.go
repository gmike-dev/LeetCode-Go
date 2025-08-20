package n2088_count_fertile_pyramids_in_a_land

func countPyramids(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	if n < 2 || m < 3 {
		return 0
	}
	res := count(grid, n, m)
	flip(grid, n)
	res += count(grid, n, m)
	return res
}

func count(grid [][]int, n, m int) int {
	res := 0
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
	}
	for i := 1; i < n; i++ {
		for j := 2; j < m; j++ {
			if grid[i][j] == 1 && grid[i-1][j-1] == 1 && grid[i][j-2] == 1 && grid[i][j-1] == 1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-2]) + 1
				res += dp[i][j]
			}
		}
	}
	return res
}

func flip(grid [][]int, n int) {
	for i := range n / 2 {
		grid[i], grid[n-i-1] = grid[n-i-1], grid[i]
	}
}
