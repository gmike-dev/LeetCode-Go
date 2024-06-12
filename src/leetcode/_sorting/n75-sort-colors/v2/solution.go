package leetcode

func sortColors(nums []int) {
	n := len(nums)
	r, w, b := 0, 0, n-1
	for w <= b {
		if nums[w] == 0 {
			nums[r], nums[w] = nums[w], nums[r]
			w++
			r++
		} else if nums[w] == 1 {
			w++
		} else {
			nums[w], nums[b] = nums[b], nums[w]
			b--
		}
	}
}
