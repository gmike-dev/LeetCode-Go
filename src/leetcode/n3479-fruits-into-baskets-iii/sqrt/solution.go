package sqrt

import "math"

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(fruits)
	segmentSize := int(math.Sqrt(float64(n)))
	segments := make([]int, (n+segmentSize-1)/segmentSize)
	for i := range n {
		j := i / segmentSize
		segments[j] = max(segments[j], baskets[i])
	}
	unplaced := n
	for i := range n {
		for segment := 0; segment < len(segments); segment++ {
			if segments[segment] >= fruits[i] {
				start := segment * segmentSize
				end := min(n, start+segmentSize)
				for j := start; j < end; j++ {
					if baskets[j] >= fruits[i] {
						baskets[j] = 0
						mx := 0
						for k := start; k < end; k++ {
							mx = max(mx, baskets[k])
						}
						segments[segment] = mx
						break
					}
				}
				unplaced--
				break
			}
		}
	}
	return unplaced
}
