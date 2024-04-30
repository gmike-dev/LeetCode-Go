package n834_sum_of_distances_in_tree

var n int
var g [][]int
var cnt []int
var sum []int

func sumOfDistancesInTree(numberOfNodes int, edges [][]int) []int {
	n = numberOfNodes
	g = make([][]int, n)
	cnt = make([]int, n)
	sum = make([]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	for i := 0; i < n; i++ {
		cnt[i] = 1
	}
	CountSumChild(0, -1)
	AddParentSum(0, -1)
	return sum
}

func CountSumChild(v int, parent int) {
	for _, child := range g[v] {
		if child != parent {
			CountSumChild(child, v)
			cnt[v] += cnt[child]
			sum[v] += sum[child] + cnt[child]
		}
	}
}

func AddParentSum(v int, parent int) {
	for _, child := range g[v] {
		if child != parent {
			sum[child] = sum[v] + n - 2*cnt[child]
			AddParentSum(child, v)
		}
	}
}
