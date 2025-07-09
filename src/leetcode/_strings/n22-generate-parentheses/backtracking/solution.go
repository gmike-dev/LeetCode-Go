package bitmasks

func generateParenthesis(n int) []string {
	var (
		res []string
		s   = make([]byte, 2*n)
	)
	var gen func(int, int)
	gen = func(open, close int) {
		i := open + close
		if i == 2*n {
			res = append(res, string(s))
		}
		if open < n {
			s[i] = '('
			gen(open+1, close)
		}
		if close < open {
			s[i] = ')'
			gen(open, close+1)
		}
	}
	gen(0, 0)
	return res
}
