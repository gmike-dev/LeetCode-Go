package leetcode

func longestPalindrome(s string) int {
	mask := int64(0)
	for _, c := range s {
		mask ^= int64(1) << (c - 'A')
	}
	odd := popCount(mask)
	length := len(s) - odd
	if odd != 0 {
		length++
	}
	return length
}

func popCount(x int64) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}
