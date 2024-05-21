package leetcode

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	for mask := 0; mask < 1<<len(nums); mask++ {
		subset := make([]int, 0)
		m := mask
		for i := 0; m != 0; i++ {
			if m&1 != 0 {
				subset = append(subset, nums[i])
			}
			m >>= 1
		}
		result = append(result, subset)
	}

	return result
}
