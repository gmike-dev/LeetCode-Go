package leetcode

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	p := make([]int, 26)
	n := len(s1)
	for i := 0; i < 26; i++ {
		p[i] = i
	}
	for i := 0; i < n; i++ {
		x, y := int(s1[i]-'a'), int(s2[i]-'a')
		px, py := x, y
		for p[px] != px {
			px = p[px]
		}
		for p[py] != py {
			py = p[py]
		}
		if px < py {
			p[py] = px
		} else {
			p[px] = py
		}
	}
	m := len(baseStr)
	result := make([]byte, m)
	for i, c := range baseStr {
		pc := int(c - 'a')
		for p[pc] != pc {
			pc = p[pc]
		}
		result[i] = byte('a' + pc)
	}
	return string(result)
}
