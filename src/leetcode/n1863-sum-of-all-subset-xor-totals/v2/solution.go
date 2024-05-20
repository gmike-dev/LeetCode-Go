package leetcode

func subsetXORSum(a []int) int {
	s := 0
	subsetXor(a, 0, 0, &s)
	return s
}

func subsetXor(a []int, i int, xor int, s *int) {
	if i < len(a) {
		*s += xor ^ a[i]
		subsetXor(a, i+1, xor^a[i], s)
		subsetXor(a, i+1, xor, s)
	}
}
