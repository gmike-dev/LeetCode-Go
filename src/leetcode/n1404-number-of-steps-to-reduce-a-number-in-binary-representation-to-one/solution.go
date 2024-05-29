package leetcode

func numSteps(s string) int {
	count := len(s) - 1
	carry := 0
	for i := len(s) - 1; i > 0; i-- {
		c := int(s[i] - '0')
		if carry+c == 1 {
			count++
			carry = 1
		}
	}
	return count + carry
}
