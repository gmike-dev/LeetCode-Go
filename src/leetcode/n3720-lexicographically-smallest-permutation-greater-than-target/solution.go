package n3720_lexicographically_smallest_permutation_greater_than_target

func lexGreaterPermutation(s string, t string) string {
	cnt := make([]int, 26)
	n := len(s)
	for i := range n {
		cnt[s[i]-'a']++
	}
	a := make([]byte, n)

	var f func(int) bool

	f = func(pos int) bool {
		if pos == n {
			return false
		}
		tc := t[pos] - 'a'
		if cnt[tc] > 0 {
			a[pos] = t[pos]
			cnt[tc]--
			if f(pos + 1) {
				return true
			}
			cnt[tc]++
		}
		for c := tc + 1; c < 26; c++ {
			if cnt[c] > 0 {
				a[pos] = c + 'a'
				cnt[c]--
				pos++
				for i := byte(0); i < 26; i++ {
					for cnt[i] > 0 {
						a[pos] = i + 'a'
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
