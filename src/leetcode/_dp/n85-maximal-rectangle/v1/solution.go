package v1

func maximalRectangle(matrix [][]byte) int {
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, m+1)
	}
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if matrix[i-1][j-1] != '0' {
				dp[i][j] = dp[i-1][j] + 1
				h := dp[i][j]
				for w := 1; h > 0; w++ {
					res = max(res, w*h)
					h = min(h, dp[i][j-w])
				}
			}
		}
	}
	return res
}
