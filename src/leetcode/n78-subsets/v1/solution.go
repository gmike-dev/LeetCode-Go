package leetcode

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	findSubsets(nums, 0, []int{}, &result)
	return result
}

func findSubsets(nums []int, i int, curr []int, result *[][]int) {
	if i >= len(nums) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*result = append(*result, tmp)
		return
	}
	findSubsets(nums, i+1, curr, result)
	findSubsets(nums, i+1, append(curr, nums[i]), result)
}
