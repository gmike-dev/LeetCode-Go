package leetcode

func subsetXORSum(a []int) int {
	return subsetXor(a, 0, 0)
}

func subsetXor(a []int, i int, xor int) int {
	if i >= len(a) {
		return xor
	}
	return subsetXor(a, i+1, xor^a[i]) + subsetXor(a, i+1, xor)
}
