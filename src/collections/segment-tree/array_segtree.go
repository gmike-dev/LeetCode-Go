package segment_tree

type ArraySegTree struct {
	n    int
	tree []int
	agg  func(int, int) int
}

func NewArraySegTree(data []int, agg func(int, int) int) *ArraySegTree {
	n := len(data)
	t := &ArraySegTree{
		n:    n,
		tree: make([]int, 4*n),
		agg:  agg,
	}
	t.build(data, 1, 0, n-1)
	return t
}

func (t *ArraySegTree) Query(l, r int) int {
	return t.query(1, 0, t.n-1, l, r)
}

func (t *ArraySegTree) Update(pos int, value int) {
	t.update(1, 0, t.n-1, pos, value)
}

func (t *ArraySegTree) update(v, tl, tr, pos int, value int) {
	if tl == tr {
		t.tree[v] = value
	} else {
		tm := tl + (tr-tl)/2
		if pos <= tm {
			t.update(v*2, tl, tm, pos, value)
		} else {
			t.update(v*2+1, tm+1, tr, pos, value)
		}
		t.tree[v] = t.agg(t.tree[v*2], t.tree[v*2+1])
	}
}

func (t *ArraySegTree) query(v, tl, tr, l, r int) int {
	if l > r {
		return 0
	}
	if l == tl && r == tr {
		return t.tree[v]
	}
	tm := tl + (tr-tl)/2
	return t.agg(t.query(v*2, tl, tm, l, min(r, tm)), t.query(v*2+1, tm+1, tr, max(l, tm+1), r))
}

func (t *ArraySegTree) build(data []int, v, tl, tr int) {
	if tl == tr {
		t.tree[v] = data[tl]
	} else {
		tm := tl + (tr-tl)/2
		t.build(data, v*2, tl, tm)
		t.build(data, v*2+1, tm+1, tr)
		t.tree[v] = t.agg(t.tree[v*2], t.tree[v*2+1])
	}
}
