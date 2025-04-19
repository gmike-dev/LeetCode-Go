package n2563_count_the_number_of_fair_pairs

import (
	"slices"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	slices.Sort(nums)
	var result int64
	n := len(nums)
	var left, right int
	for i := n - 1; i > left; i-- {
		right = min(right, i-1)
		for left < i && nums[left]+nums[i] < lower {
			left++
		}
		for right < i && nums[right]+nums[i] <= upper {
			right++
		}
		result += int64(right - left)
	}
	return result
}
