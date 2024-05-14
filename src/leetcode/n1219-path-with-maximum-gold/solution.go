package leetcode

var grid [][]int
var visited [][]bool
var m, n int

func getMaximumGold(g [][]int) int {
	grid = g
	m = len(g)
	n = len(g[0])
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	maxGold := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			maxGold = max(maxGold, getGold(i, j))
		}
	}
	return maxGold
}

func getGold(row, col int) int {
	if row < 0 || col < 0 || row >= m || col >= n || grid[row][col] == 0 || visited[row][col] {
		return 0
	}
	visited[row][col] = true
	gold := grid[row][col] + max(
		max(getGold(row-1, col), getGold(row+1, col)),
		max(getGold(row, col-1), getGold(row, col+1)))
	visited[row][col] = false
	return gold
}
