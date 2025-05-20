package n1931_painting_a_grid_with_three_different_colors

import "math"

func colorTheGrid(m int, n int) int {
	const Mod = int(1e9) + 7
	mc := int(math.Pow(3, float64(m)))
	isDistinct := make([]bool, mc)
	for i := 0; i < mc; i++ {
		isDistinct[i] = true
		x := i
		for j := 0; j < m-1; j++ {
			if x%3 == x/3%3 {
				isDistinct[i] = false
				break
			}
			x /= 3
		}
	}

	canJoin := make([][]bool, mc)
	for i := 0; i < mc; i++ {
		canJoin[i] = make([]bool, mc)
	}
	for i := 0; i < mc; i++ {
		if !isDistinct[i] {
			continue
		}
		for j := 0; j < i; j++ {
			if !isDistinct[j] {
				continue
			}
			canJoin[i][j] = true
			canJoin[j][i] = true
			x, y := i, j
			for k := 0; k < m; k++ {
				if x%3 == y%3 {
					canJoin[i][j] = false
					canJoin[j][i] = false
					break
				}
				x /= 3
				y /= 3
			}
		}
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, mc)
	}
	for i := 0; i < mc; i++ {
		if isDistinct[i] {
			dp[0][i] = 1
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < mc; j++ {
			if !isDistinct[j] {
				continue
			}
			for k := 0; k < j; k++ {
				if canJoin[j][k] {
					dp[i][j] = (dp[i][j] + dp[i-1][k]) % Mod
					dp[i][k] = (dp[i][k] + dp[i-1][j]) % Mod
				}
			}
		}
	}
	sum := 0
	for i := 0; i < mc; i++ {
		sum = (sum + dp[n-1][i]) % Mod
	}
	return sum
}
