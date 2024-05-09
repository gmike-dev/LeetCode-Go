package n3075_maximize_happiness_of_selected_children

import "slices"

func maximumHappinessSum(happiness []int, k int) int64 {
	slices.SortFunc(happiness, func(a, b int) int { return b - a })
	var h int64
	var s int
	for i := 0; i < k; i++ {
		h += int64(max(0, happiness[i]-s))
		s++
	}
	return h
}
