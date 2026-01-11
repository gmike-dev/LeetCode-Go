package monotonic_stack

func maximalRectangle(a [][]byte) int {
	n, m := len(a), len(a[0])
	h := make([]int, m+1)
	res := 0
	sh := make([]int, m+1)
	si := make([]int, m+1)
	sn := 1
	for i := 1; i <= n; i++ {
		sn = 1
		for j := 1; j <= m; j++ {
			if a[i-1][j-1] != '0' {
				h[j]++
				k := j
				for sh[sn-1] >= h[j] {
					k = si[sn-1]
					sn--
				}
				sh[sn] = h[j]
				si[sn] = k
				sn++
				for l := range sn {
					res = max(res, (j-si[l]+1)*sh[l])
				}
			} else {
				h[j] = 0
				sn = 1
			}
		}
	}
	return res
}
