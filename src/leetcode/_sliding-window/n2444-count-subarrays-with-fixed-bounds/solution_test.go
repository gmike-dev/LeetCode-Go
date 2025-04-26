package n2444_count_subarrays_with_fixed_bounds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSubarrays1(t *testing.T) {
	assert.Equal(t, int64(2), countSubarrays([]int{1, 3, 5, 2, 7, 5}, 1, 5))
}

func TestCountSubarrays2(t *testing.T) {
	assert.Equal(t, int64(10), countSubarrays([]int{1, 1, 1, 1}, 1, 1))
}

func TestCountSubarrays3(t *testing.T) {
	assert.Equal(t, int64(81), countSubarrays([]int{35054, 398719, 945315, 945315, 820417, 945315, 35054, 945315,
		171832, 945315, 35054, 109750, 790964, 441974, 552913}, 35054, 945315))
}
