package leetcode

import "container/heap"

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	dist := Bfs(grid)
	used := make([][]bool, n)
	for i := 0; i < n; i++ {
		used[i] = make([]bool, n)
	}
	q := make(PriorityQueue, 0)
	heap.Push(&q, &Item{[]int{0, 0}, dist[0][0]})
	used[0][0] = true
	for len(q) > 0 {
		item := heap.Pop(&q).(*Item)
		x, y, f := item.value[0], item.value[1], item.priority
		if x == n-1 && y == n-1 {
			return f
		}
		for i := 0; i < 4; i++ {
			nx, ny := x+dx[i], y+dy[i]
			if nx >= 0 && ny >= 0 && nx < n && ny < n && !used[nx][ny] {
				heap.Push(&q, &Item{[]int{nx, ny}, min(f, dist[nx][ny])})
				used[nx][ny] = true
			}
		}
	}
	return -1
}

func Bfs(grid [][]int) [][]int {
	n := len(grid)
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = n
		}
	}
	q := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j, 0})
			}
		}
	}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		x, y, d := p[0], p[1], p[2]
		if x < 0 || y < 0 || x >= n || y >= n || d >= dist[x][y] {
			continue
		}
		dist[x][y] = d
		q = append(q,
			[]int{x - 1, y, d + 1},
			[]int{x + 1, y, d + 1},
			[]int{x, y - 1, d + 1},
			[]int{x, y + 1, d + 1})
	}
	return dist
}

type Item struct {
	value    []int
	priority int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int { return len(*pq) }

func (pq *PriorityQueue) Less(i, j int) bool { return (*pq)[i].priority > (*pq)[j].priority }

func (pq *PriorityQueue) Swap(i, j int) { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }

func (pq *PriorityQueue) Push(x any) { *pq = append(*pq, x.(*Item)) }

func (pq *PriorityQueue) Pop() any {
	q := *pq
	n := len(q)
	item := q[n-1]
	q[n-1] = nil
	*pq = q[0 : n-1]
	return item
}
