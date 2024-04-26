package n1289_minimum_falling_path_sum_ii

import (
	"slices"
)

// O^2 solution
func minFallingPathSumV2(grid [][]int) int {
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}
	for i := 1; i < n; i++ {
		prev := grid[i-1]
		min1, min2 := -1, -1
		for j := 0; j < n; j++ {
			if min1 == -1 || prev[j] <= prev[min1] {
				min2 = min1
				min1 = j
			} else if min2 == -1 || prev[j] < prev[min2] {
				min2 = j
			}
		}
		for j := 0; j < n; j++ {
			var minS int
			if j == min1 {
				minS = min2
			} else {
				minS = min1
			}
			grid[i][j] += prev[minS]
		}
	}
	return slices.Min(grid[n-1])
}
