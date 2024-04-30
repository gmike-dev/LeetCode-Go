package n310_minimum_height_trees

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	g := make([][]int, n)
	cnt := make([]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		cnt[a]++
		cnt[b]++
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}
	q := make([]int, 0)
	for i := 0; i < n; i++ {
		if cnt[i] == 1 {
			q = append(q, i)
		}
	}
	remain := n
	for remain > 2 {
		size := len(q)
		remain -= size
		for i := 0; i < size; i++ {
			v := q[0]
			q = q[1:]
			for _, next := range g[v] {
				cnt[next]--
				if cnt[next] == 1 {
					q = append(q, next)
				}
			}
		}
	}
	return q
}
