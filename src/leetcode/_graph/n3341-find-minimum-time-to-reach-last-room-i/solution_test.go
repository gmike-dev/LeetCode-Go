package n3341_find_minimum_time_to_reach_last_room_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinTimeToReach_1(t *testing.T) {
	actual := minTimeToReach([][]int{
		{0, 4},
		{4, 4},
	})
	expected := 6
	assert.Equal(t, expected, actual)
}

func TestMinTimeToReach_2(t *testing.T) {
	actual := minTimeToReach([][]int{
		{0, 0, 0},
		{0, 0, 0},
	})
	expected := 3
	assert.Equal(t, expected, actual)
}

func TestMinTimeToReach_3(t *testing.T) {
	actual := minTimeToReach([][]int{
		{0, 1},
		{1, 2},
	})
	expected := 3
	assert.Equal(t, expected, actual)
}
