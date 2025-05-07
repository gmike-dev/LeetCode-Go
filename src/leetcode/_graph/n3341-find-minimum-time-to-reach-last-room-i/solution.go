package n3341_find_minimum_time_to_reach_last_room_i

import (
	"container/heap"
)

type HeapItem struct {
	time int
	pos  int
}

type MinHeap []HeapItem

func (h *MinHeap) Len() int {
	return len(*h)
}
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].time < (*h)[j].time
}
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(HeapItem))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

var d = []int{0, 1, 0, -1, 0}

func minTimeToReach(moveTime [][]int) int {
	m := len(moveTime)
	n := len(moveTime[0])
	visited := make([]bool, m*n)
	visited[0] = true

	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, HeapItem{0, 0})

	for h.Len() > 0 {
		item := heap.Pop(h).(HeapItem)
		r, c, time := item.pos/n, item.pos%n, item.time

		for i := 0; i < 4; i++ {
			nr, nc := r+d[i], c+d[i+1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			np := nr*n + nc
			if visited[np] {
				continue
			}
			nextTime := max(moveTime[nr][nc], time) + 1
			if np == len(visited)-1 {
				return nextTime
			}
			visited[np] = true
			heap.Push(h, HeapItem{nextTime, np})
		}
	}
	return -1
}
