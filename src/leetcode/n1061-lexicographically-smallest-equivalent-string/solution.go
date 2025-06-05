package leetcode

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	p := make([]int, 26)
	n := len(s1)
	for i := 0; i < 26; i++ {
		p[i] = i
	}
	for i := 0; i < n; i++ {
		x, y := int(s1[i]-'a'), int(s2[i]-'a')
		for p[x] != x {
			x = p[x]
		}
		for p[y] != y {
			y = p[y]
		}
		if x < y {
			p[y] = x
		} else {
			p[x] = y
		}
	}
	result := make([]byte, len(baseStr))
	for i, c := range baseStr {
		x := int(c - 'a')
		for p[x] != x {
			x = p[x]
		}
		result[i] = byte('a' + x)
	}
	return string(result)
}
