package n1289_minimum_falling_path_sum_ii

import (
	"math"
	"slices"
)

// O^3 solution
func minFallingPathSumV1(grid [][]int) int {
	n := len(grid)
	for i := 1; i < n; i++ {
		prev := grid[i-1]
		for j := 0; j < n; j++ {
			minS := math.MaxInt
			for k := 0; k < n; k++ {
				if k != j {
					minS = min(minS, prev[k])
				}
			}
			grid[i][j] += minS
		}
	}
	return slices.Min(grid[n-1])
}
