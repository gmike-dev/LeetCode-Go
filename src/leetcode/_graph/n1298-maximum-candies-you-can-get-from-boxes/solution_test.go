package n1298_maximum_candies_you_can_get_from_boxes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxCandies1(t *testing.T) {
	assert.Equal(t, 16, maxCandies(
		[]int{1, 0, 1, 0},
		[]int{7, 5, 4, 100},
		[][]int{{}, {}, {1}, {}},
		[][]int{{1, 2}, {3}, {}, {}},
		[]int{0}))
}

func Test_maxCandies2(t *testing.T) {
	assert.Equal(t, 6, maxCandies(
		[]int{1, 0, 0, 0, 0, 0},
		[]int{1, 1, 1, 1, 1, 1},
		[][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
		[][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
		[]int{0}))
}
