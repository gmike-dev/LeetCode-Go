package segmenttree

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	n := len(fruits)
	t := newSegmentTree(baskets)
	unplaced := 0
	for i := range n {
		l, r := 0, n
		for l < r {
			m := l + (r-l)/2
			if t.QueryMax(0, m) < fruits[i] {
				l = m + 1
			} else {
				r = m
			}
		}
		if l == n {
			unplaced++
		} else {
			t.Update(l, 0)
		}
	}
	return unplaced
}

type segmentTree struct {
	values []int
	size   int
}

func newSegmentTree(data []int) *segmentTree {
	n := len(data)
	t := &segmentTree{
		values: make([]int, 4*n),
		size:   n,
	}
	var build func(int, int, int)
	build = func(v, tl, tr int) {
		if tl == tr {
			t.values[v] = data[tl]
			return
		}
		tm := tl + (tr-tl)/2
		build(2*v+1, tl, tm)
		build(2*v+2, tm+1, tr)
		t.values[v] = max(t.values[2*v+1], t.values[2*v+2])
	}
	build(0, 0, n-1)
	return t
}

func (t *segmentTree) Update(index, value int) {
	t.update(0, 0, t.size-1, index, value)
}

func (t *segmentTree) QueryMax(left, right int) int {
	return t.queryMax(0, 0, t.size-1, left, right)
}

func (t *segmentTree) queryMax(v, tl, tr, l, r int) int {
	if l > r {
		return 0
	}
	if tl == l && tr == r {
		return t.values[v]
	}
	tm := tl + (tr-tl)/2
	return max(
		t.queryMax(2*v+1, tl, tm, l, min(r, tm)),
		t.queryMax(2*v+2, tm+1, tr, max(tm+1, l), r))
}

func (t *segmentTree) update(v, tl, tr, index, value int) {
	if tl == tr {
		t.values[v] = value
		return
	}
	tm := tl + (tr-tl)/2
	if index <= tm {
		t.update(2*v+1, tl, tm, index, value)
	} else {
		t.update(2*v+2, tm+1, tr, index, value)
	}
	t.values[v] = max(t.values[2*v+1], t.values[2*v+2])
}
