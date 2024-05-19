package v1

import "math"

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	sum := int64(0)
	sub := math.MaxInt
	negCnt := 0
	for _, num := range nums {
		sum += int64(max(num, num^k))
		s := num - (num ^ k)
		if s < 0 {
			negCnt++
		}
		sub = min(sub, abs(s))
	}
	if negCnt%2 != 0 {
		sum -= int64(sub)
	}
	return sum
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
