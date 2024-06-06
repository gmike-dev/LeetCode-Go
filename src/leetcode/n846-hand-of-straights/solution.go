package leetcode

func isNStraightHand(hand []int, groupSize int) bool {
	if groupSize == 1 {
		return true
	}
	n := len(hand)
	if n%groupSize != 0 {
		return false
	}

}
