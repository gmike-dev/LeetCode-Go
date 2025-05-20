package n1931_painting_a_grid_with_three_different_colors

func colorTheGrid(m int, n int) int {
	const Mod = int(1e9) + 7
	var pow3 = [...]int{1, 3, 9, 27, 81, 243}
	var distinct []int
	for i := 0; i < pow3[m]; i++ {
		if isDistinct(i, m) {
			distinct = append(distinct, i)
		}
	}
	mc := len(distinct)
	canJoin := make([][]bool, mc)
	for i := 0; i < mc; i++ {
		canJoin[i] = make([]bool, mc)
	}
	for i := 0; i < mc; i++ {
		for j := 0; j < i; j++ {
			if isCompatible(distinct[i], distinct[j], m) {
				canJoin[i][j] = true
				canJoin[j][i] = true
			}
		}
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, mc)
	}
	for i := 0; i < mc; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < mc; j++ {
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

func isDistinct(x, m int) bool {
	for j := 0; j < m-1; j++ {
		if x%3 == x/3%3 {
			return false
		}
		x /= 3
	}
	return true
}

func isCompatible(x, y, m int) bool {
	for k := 0; k < m; k++ {
		if x%3 == y%3 {
			return false
		}
		x /= 3
		y /= 3
	}
	return true
}
