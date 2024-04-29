package n2997_minimum_number_of_operations_to_make_array_xor_equal_to_k

import "math/bits"

func minOperations(nums []int, k int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return bits.OnesCount(uint(xor ^ k))
}
