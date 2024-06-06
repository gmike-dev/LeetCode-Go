package collections

import "cmp"

func BuildMaxHeap(a []int) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		PushDown(a, i, n, cmp.Compare[int])
	}
}

func BuildMinHeap(a []int) {
	n := len(a)
	invCompare := func(x, y int) int {
		return -cmp.Compare(x, y)
	}
	for i := n/2 - 1; i >= 0; i-- {
		PushDown(a, i, n, invCompare)
	}
}

func PushDown(a []int, i, n int, cmp func(int, int) int) {
	for {
		largest := i
		l := 2*i + 1
		r := l + 1
		if l < n && cmp(a[l], a[largest]) > 0 {
			largest = l
		}
		if r < n && cmp(a[r], a[largest]) > 0 {
			largest = r
		}
		if largest == i {
			break
		}
		a[largest], a[i] = a[i], a[largest]
		i = largest
	}
}
