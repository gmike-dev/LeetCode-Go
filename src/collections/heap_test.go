package collections

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildMaxHeap1(t *testing.T) {
	heap := []int{1, 2, 3}
	BuildMaxHeap(heap)
	assert.Equal(t, []int{3, 2, 1}, heap)
}

func TestBuildMaxHeap2(t *testing.T) {
	heap := []int{4, 1, 2, 3, 65, 7, 2, 3, 1, 2, 3, 6, -1, 23, 0, 0, -13}
	BuildMaxHeap(heap)
	assert.Equal(t, []int{65, 4, 23, 3, 3, 7, 2, 3, 1, 2, 1, 6, -1, 2, 0, 0, -13}, heap)
}
