package n1353_maximum_number_of_events_that_can_be_attended

import (
	"container/heap"
	"math"
	"slices"
)

func maxEvents(events [][]int) int {
	slices.SortFunc(events, func(s1, s2 []int) int { return s1[0] - s2[0] })
	minDay, maxDay := math.MaxInt, 0
	for _, e := range events {
		minDay = min(minDay, e[0])
		maxDay = max(maxDay, e[1])
	}
	h := &Heap{}
	ans := 0
	i := 0
	n := len(events)
	for day := minDay; day <= maxDay; day++ {
		for i < n && events[i][0] <= day {
			heap.Push(h, events[i][1])
			i++
		}
		for h.Len() != 0 && (*h)[0] < day {
			heap.Pop(h)
		}
		if h.Len() != 0 {
			ans++
			heap.Pop(h)
		}
	}
	return ans
}

var _ heap.Interface = (*Heap)(nil)

type Heap []int

func (p *Heap) Len() int { return len(*p) }

func (p *Heap) Less(i, j int) bool { return (*p)[i] < (*p)[j] }

func (p *Heap) Swap(i, j int) { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }

func (p *Heap) Push(x any) { *p = append(*p, x.(int)) }

func (p *Heap) Pop() any {
	n := len(*p)
	res := (*p)[n-1]
	*p = (*p)[:n-1]
	return res
}
