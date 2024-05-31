package leetcode

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	lsb := xor & -xor
	x := 0
	for _, num := range nums {
		if lsb&num == 0 {
			x ^= num
		}
	}
	return []int{x, xor ^ x}
}
