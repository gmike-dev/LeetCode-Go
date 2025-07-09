package dp

import (
	"cmp"
	"slices"
)

func maxValue(events [][]int, k int) int {
	slices.SortFunc(events, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})
	n := len(events)
	prev := make([]int, n)
	for i, e := range events {
		l, r := 0, i
		for l < r {
			m := l + (r-l)/2
			if events[m][1] < e[0] {
				l = m + 1
			} else {
				r = m
			}
		}
		prev[i] = r - 1
	}
	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 1; i <= k; i++ {
		for j := i - 1; j < n; j++ {
			dp[i][j] = max(dp[i][max(0, j-1)], events[j][2])
			if prev[j] >= 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][prev[j]]+events[j][2])
			}
		}
	}
	ans := 0
	for _, a := range dp {
		ans = max(ans, a[n-1])
	}
	return ans
}
