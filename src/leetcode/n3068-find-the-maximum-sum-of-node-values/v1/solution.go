package v1

import (
	"slices"
)

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	n := len(nums)
	d := make([]int, n)
	sum := int64(0)
	for i := 0; i < n; i++ {
		d[i] = (nums[i] ^ k) - nums[i]
		sum += int64(nums[i])
	}
	slices.SortFunc(d, func(a, b int) int {
		return b - a
	})
	maxAdd := int64(0)
	for i := 0; i < n-1; i += 2 {
		add := d[i] + d[i+1]
		if add <= 0 {
			break
		}
		maxAdd += int64(add)
	}
	return sum + maxAdd
}
