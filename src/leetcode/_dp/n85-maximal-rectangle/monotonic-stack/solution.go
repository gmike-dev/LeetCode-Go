package monotonic_stack

func maximalRectangle(a [][]byte) int {
	n, m := len(a), len(a[0])
	h := make([]int, m+1)
	s := make([]int, m+1)
	sn := 1
	res := 0
	for i := 1; i <= n; i++ {
		s[0] = 0
		sn = 1
		for j := 1; j <= m; j++ {
			if a[i-1][j-1] != '0' {
				h[j]++
				for h[s[sn-1]] >= h[j] {
					sn--
				}
				s[sn] = j
				sn++
				for k := 1; k < sn; k++ {
					res = max(res, (j-s[k-1])*h[s[k]])
				}
			} else {
				h[j] = 0
				s[0] = j
				sn = 1
			}
		}
	}
	return res
}
