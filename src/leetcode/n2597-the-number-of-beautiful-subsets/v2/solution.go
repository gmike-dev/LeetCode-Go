package leetcode

import "slices"

func beautifulSubsets(nums []int, k int) int {
	slices.Sort(nums)
	n := len(nums)
	s := make([]int, 0)
	result := 0
	for mask := 1; mask < 1<<n; mask++ {
		m := mask
		s = s[:0]
		for i := 0; m != 0; i++ {
			if m&1 != 0 {
				s = append(s, nums[i])
			}
			m >>= 1
		}
		if beautiful(s, k) {
			result++
		}
	}
	return result
}

func beautiful(s []int, k int) bool {
	j := 0
	for i := 0; i < len(s)-1; i++ {
		for ; j < len(s) && s[j]-s[i] < k; j++ {
		}
		if j == len(s) {
			break
		}
		if s[j]-s[i] == k {
			return false
		}
	}
	return true
}
