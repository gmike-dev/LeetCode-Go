package recursion

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
	cache := make([][]int, n)
	for i := range cache {
		cache[i] = make([]int, k+1)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i int, m int) int {
		if i < 0 || m == 0 {
			return 0
		}
		if cache[i][m] != -1 {
			return cache[i][m]
		}
		ans := max(f(i-1, m), f(prev[i], m-1)+events[i][2])
		cache[i][m] = ans
		return ans
	}
	return f(n-1, k)
}
