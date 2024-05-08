package n506_relative_ranks

import (
	"slices"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	n := len(score)
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	slices.SortFunc(p, func(a, b int) int {
		return score[b] - score[a]
	})
	ranks := make([]string, n)
	ranks[p[0]] = "Gold Medal"
	if n > 1 {
		ranks[p[1]] = "Silver Medal"
		if n > 2 {
			ranks[p[2]] = "Bronze Medal"
			for i := 3; i < n; i++ {
				ranks[p[i]] = strconv.Itoa(i + 1)
			}
		}
	}
	return ranks
}
