package leetcode

func longestPalindrome(s string) int {
	cnt := [128]int{}
	for _, c := range s {
		cnt[c]++
	}
	odd := 0
	length := 0
	for _, c := range cnt {
		length += c & -2
		odd |= c & 1
	}
	return length + odd
}
