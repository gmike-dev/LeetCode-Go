package n3748_count_stable_subarrays

func countStableSubarrays(nums []int, queries [][]int) []int64 {
	n := len(nums)
	prefixSum := make([]int64, n)
	prefixSum[0] = 1
	for i := 1; i < n; i++ {
		if nums[i-1] <= nums[i] {
			prefixSum[i] = prefixSum[i-1] + 1
		} else {
			prefixSum[i] = 1
		}
	}
	for i := 1; i < n; i++ {
		prefixSum[i] += prefixSum[i-1]
	}

	end := make([]int, n)
	end[n-1] = n - 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			end[i] = end[i+1]
		} else {
			end[i] = i
		}
	}

	m := len(queries)
	result := make([]int64, m)
	for i, q := range queries {
		l, r := q[0], q[1]
		ll := min(r, end[l])
		k := int64(ll - l + 1)
		result[i] = k*(k+1)/2 + prefixSum[r] - prefixSum[ll]
	}
	return result
}
