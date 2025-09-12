package n3495_minimum_operations_to_make_array_elements_zero

import "math/bits"

func minOperations(queries [][]int) int64 {
	var res int64
	for _, q := range queries {
		l, r := q[0], q[1]
		divCount := div4Count(r) - div4Count(l-1)
		res += (divCount + 1) / 2
	}
	return res
}

func div4Count(n int) int64 {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	b := int64(bits.Len(uint(n)))
	x := int64(n - (1 << (b - 1)) + 1)
	op := x*((b+1)/2) + 1
	cnt := int64(2)
	for i := int64(2); i < b; i++ {
		op += cnt * ((i + 1) / 2)
		cnt <<= 1
	}
	return op
}
