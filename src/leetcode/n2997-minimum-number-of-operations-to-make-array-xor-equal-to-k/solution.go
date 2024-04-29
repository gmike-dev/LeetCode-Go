package n2997_minimum_number_of_operations_to_make_array_xor_equal_to_k

func minOperations(nums []int, k int) int {
	xor := 0
	for i := 0; i < len(nums); i++ {
		xor ^= nums[i]
	}
	ans := 0
	diff := xor ^ k
	for diff != 0 {
		ans += diff & 1
		diff >>= 1
	}
	return ans
}
