package n1353_maximum_number_of_events_that_can_be_attended

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxEvents1(t *testing.T) {
	actual := maxEvents([][]int{
		{1, 2}, {2, 3}, {3, 4},
	})
	assert.Equal(t, 3, actual)
}

func Test_maxEvents2(t *testing.T) {
	actual := maxEvents([][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 2},
	})
	assert.Equal(t, 4, actual)
}

func Test_maxEvents3(t *testing.T) {
	actual := maxEvents([][]int{
		{1, 4}, {4, 4}, {2, 2}, {3, 4}, {1, 1},
	})
	assert.Equal(t, 4, actual)
}

func Test_maxEvents4(t *testing.T) {
	actual := maxEvents([][]int{
		{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3},
	})
	assert.Equal(t, 5, actual)
}
