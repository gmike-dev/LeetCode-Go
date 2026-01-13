package binary_search

func separateSquares(squares [][]int) float64 {
	l := 0.0
	r := 0.0
	halfArea := 0.0
	for _, s := range squares {
		_, y, a := float64(s[0]), float64(s[1]), float64(s[2])
		r = max(r, y+a)
		halfArea += a * a
	}
	halfArea /= 2.0
	for r-l > 1e-5 {
		m := l + (r-l)/2.0
		downArea := 0.0
		for _, s := range squares {
			_, y, a := float64(s[0]), float64(s[1]), float64(s[2])
			downArea += max(0.0, min(a, m-y)) * a
			if downArea > halfArea {
				break
			}
		}
		if downArea < halfArea {
			l = m
		} else {
			r = m
		}
	}
	return l
}
