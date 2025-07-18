package n2163_minimum_difference_in_sums_after_removal_of_elements

import (
	"cmp"
	"container/heap"
	"math"
)

func minimumDifference(nums []int) int64 {
	n := len(nums) / 3
	var sl, sr int64
	minHeap := NewMinHeap[int](n)
	maxHeap := NewMaxHeap[int](n)
	minSum := make([]int64, len(nums))
	maxSum := make([]int64, len(nums))
	for i := range n {
		heap.Push(minHeap, nums[i])
		sl += int64(nums[i])
		minSum[i] = sl
	}
	for i := n; i < 2*n; i++ {
		sl += int64(nums[i])
		heap.Push(minHeap, nums[i])
		sl -= int64(heap.Pop(minHeap).(int))
		minSum[i] = sl
	}
	for i := len(nums) - 1; i >= 2*n; i-- {
		heap.Push(maxHeap, nums[i])
		sr += int64(nums[i])
		maxSum[i] = sr
	}
	for i := 2*n - 1; i >= n; i-- {
		sr += int64(nums[i])
		heap.Push(maxHeap, nums[i])
		sr -= int64(heap.Pop(maxHeap).(int))
		maxSum[i] = sr
	}
	res := int64(math.MaxInt64)
	for i := n - 1; i < 2*n; i++ {
		res = min(res, minSum[i]-maxSum[i+1])
	}
	return res
}

type Heap[T cmp.Ordered] struct {
	data []T
	less func(T, T) bool
}

func NewMinHeap[T cmp.Ordered](cap int) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0, cap),
		less: func(x, y T) bool { return x > y },
	}
}

func NewMaxHeap[T cmp.Ordered](cap int) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0, cap),
		less: func(x, y T) bool { return x < y },
	}
}

func (h *Heap[T]) Len() int           { return len(h.data) }
func (h *Heap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h *Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *Heap[T]) Push(x interface{}) {
	h.data = append(h.data, x.(T))
}

func (h *Heap[T]) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}
