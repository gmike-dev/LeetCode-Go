package manually_impl_binary_search

func smallestUniqueSubarray(nums []int) int {
	n := len(nums)

	// Rolling hash
	const p = 91138233 // must be > size of dictionary
	prefixHash := make([]int, n+1)
	powers := make([]int, n+1)
	powers[0] = 1
	for i := range n {
		powers[i+1] = powers[i] * p
	}
	for i := range n {
		prefixHash[i+1] = prefixHash[i]*p + nums[i]
	}

	hasUnique := func(m int) bool {
		count := make(map[int]int)
		for i := n - m; i >= 0; i-- {
			h := prefixHash[i+m] - prefixHash[i]*powers[m]
			count[h]++
		}
		for _, v := range count {
			if v == 1 {
				return true
			}
		}
		return false
	}

	l := 1
	r := n
	for l <= r {
		m := (l + r) / 2
		if hasUnique(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
