package linear

func lexGreaterPermutation(s string, t string) string {
	cnt := make([]int, 128)
	n := len(s)
	for i := range n {
		cnt[s[i]]++
	}
	a := make([]byte, 0, n)
	lastGreatPos := -1
	lastGreatCh := byte(0)
	for i := range n {
		tc := t[i]
		for c := tc + 1; c <= 'z'; c++ {
			if cnt[c] > 0 {
				lastGreatPos = i
				lastGreatCh = c
				break
			}
		}
		if cnt[tc] > 0 {
			a = append(a, t[i])
			cnt[tc]--
		} else {
			break
		}
	}
	if lastGreatPos == -1 {
		return ""
	}
	for i := lastGreatPos; i < len(a); i++ {
		cnt[a[i]]++
	}
	a = a[:lastGreatPos]
	a = append(a, lastGreatCh)
	cnt[lastGreatCh]--
	for c := 'a'; c <= 'z'; c++ {
		for cnt[c] > 0 {
			a = append(a, byte(c))
			cnt[c]--
		}
	}
	return string(a)
}
