package leetcode

func partition(s string) [][]string {
	result := make([][]string, 0)
	curr := make([]string, 0)
	isPalindrome := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	var findPalindromes func(int)
	findPalindromes = func(l int) {
		if l >= len(s) {
			tmp := make([]string, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}
		for r := l; r < len(s); r++ {
			if isPalindrome(l, r) {
				curr = append(curr, s[l:r+1])
				findPalindromes(r + 1)
				curr = curr[:len(curr)-1]
			}
		}
	}
	findPalindromes(0)
	return result
}
