package bitmasks

import "math/bits"

func generateParenthesis(n int) []string {
	var res []string
	for m := 0; m < 1<<(2*n); m++ {
		if valid(m, n) {
			res = append(res, parentheses(m, n))
		}
	}
	return res
}

func parentheses(m, n int) string {
	res := make([]byte, 2*n)
	for i := 0; i < 2*n; i++ {
		if m&1 == 0 {
			res[i] = '('
		} else {
			res[i] = ')'
		}
		m >>= 1
	}
	return string(res)
}

func valid(m, n int) bool {
	if bits.OnesCount(uint(m)) != n {
		return false
	}
	balance := 0
	for i := 0; i < 2*n; i++ {
		if m&1 == 0 {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
		m >>= 1
	}
	return balance == 0
}
