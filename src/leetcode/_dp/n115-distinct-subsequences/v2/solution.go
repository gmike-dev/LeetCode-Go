package v2

func numDistinct(s string, t string) int {
	m := len(s)
	n := len(t)
	if m < n {
		return 0
	}
	dp := make([]int, n+1)
	dp[n] = 1
	for i := m - 1; i >= 0; i-- {
		for j := range n {
			if s[i] == t[j] {
				dp[j] = dp[j+1] + dp[j]
			}
		}
	}
	return dp[0]
}
