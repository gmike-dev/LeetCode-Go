package segment_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraySegTree_Query(t *testing.T) {
	data := []int{0, 1, 10, 5, 0, 1, 2, 17, 18, 16}
	st := NewArraySegTree(data, func(x, y int) int { return x + y })
	checkQuery(t, st, data)
}

func TestArraySegTree_Update(t *testing.T) {
	data := []int{0, 1, 10, 5, 0, 1, 2, 17, 18, 16}
	st := NewArraySegTree(data, func(x, y int) int { return x + y })
	for i := range data {
		data[i] += 3
		st.Update(i, data[i])
		checkQuery(t, st, data)
	}
}

func checkQuery(t *testing.T, st *ArraySegTree, data []int) {
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			assert.Equal(t, sum(data[i:j+1]), st.Query(i, j))
		}
	}
}

func sum(values []int) int {
	s := 0
	for _, val := range values {
		s += val
	}
	return s
}
