package leetcode

func subsetXORSum(nums []int) int {
	n := len(nums)
	s := 0
	for mask := 1; mask < 1<<n; mask++ {
		m := mask
		xor := 0
		for i := 0; m != 0; i++ {
			if m&1 != 0 {
				xor ^= nums[i]
			}
			m >>= 1
		}
		s += xor
	}
	return s
}
