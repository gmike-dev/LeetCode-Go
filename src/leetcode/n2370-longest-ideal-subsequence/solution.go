package n2370_longest_ideal_subsequence

import "slices"

func longestIdealString(s string, k int) int {
	dp := make([]int, 26)
	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		length := 0
		for j := max(0, c-k); j <= min(25, c+k); j++ {
			length = max(length, dp[j])
		}
		dp[c] = length + 1
	}
	return slices.Max(dp)
}
