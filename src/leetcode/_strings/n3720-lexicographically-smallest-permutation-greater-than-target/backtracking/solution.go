package backtracking

func lexGreaterPermutation(s string, t string) string {
	cnt := make([]int, 128)
	n := len(s)
	for i := range n {
		cnt[s[i]]++
	}
	a := make([]byte, n)

	var f func(int) bool

	f = func(pos int) bool {
		if pos >= n {
			return false
		}
		tc := t[pos]
		if cnt[tc] > 0 {
			a[pos] = t[pos]
			cnt[tc]--
			if f(pos + 1) {
				return true
			}
			cnt[tc]++
		}
		for c := tc + 1; c <= 'z'; c++ {
			if cnt[c] > 0 {
				a[pos] = c
				cnt[c]--
				pos++
				for i := 'a'; i <= 'z'; i++ {
					for cnt[i] > 0 {
						a[pos] = byte(i)
						pos++
						cnt[i]--
					}
				}
				return true
			}
		}
		return false
	}

	if f(0) {
		return string(a)
	}
	return ""
}
