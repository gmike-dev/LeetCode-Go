package union_find

type SimpleUnionFind []int

func NewSimpleUnionFind(n int) SimpleUnionFind {
	root := make(SimpleUnionFind, n)
	for i := 0; i < n; i++ {
		root[i] = i
	}
	return root
}

func (root SimpleUnionFind) Find(x int) int {
	for root[x] != x {
		root[x] = root[root[x]]
		x = root[x]
	}
	return x
}

func (root SimpleUnionFind) Union(x, y int) {
	root[root.Find(y)] = root.Find(x)
}

func (root SimpleUnionFind) Connected(x, y int) bool {
	return root.Find(x) == root.Find(y)
}
