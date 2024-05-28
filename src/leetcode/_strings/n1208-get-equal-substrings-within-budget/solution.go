package leetcode

func equalSubstring(s string, t string, maxCost int) int {
	left := 0
	cost := 0
	maxLen := 0
	for right := 0; right < len(s); right++ {
		cost += abs(int(s[right]) - int(t[right]))
		for cost > maxCost {
			cost -= abs(int(s[left]) - int(t[left]))
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
