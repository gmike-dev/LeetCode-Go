package leetcode

import "math"

func isNStraightHand(hand []int, groupSize int) bool {
	if groupSize == 1 {
		return true
	}
	n := len(hand)
	if n%groupSize != 0 {
		return false
	}

	m := make(map[int]int)
	for _, val := range hand {
		m[val]++
	}

	for i := 0; i < n/groupSize; i++ {
		x := math.MaxInt
		for key := range m {
			x = min(x, key)
		}
		x--
		for j := 0; j < groupSize; j++ {
			x++
			if m[x] == 0 {
				return false
			}
			m[x]--
			if m[x] == 0 {
				delete(m, x)
			}
		}
	}
	return true
}
