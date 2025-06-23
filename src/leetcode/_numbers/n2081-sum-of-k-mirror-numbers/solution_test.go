package n2081_sum_of_k_mirror_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kMirror1(t *testing.T) {
	assert.Equal(t, int64(25), kMirror(2, 5))
}

func Test_kMirror2(t *testing.T) {
	assert.Equal(t, int64(499), kMirror(3, 7))
}

func Test_kMirror3(t *testing.T) {
	assert.Equal(t, int64(20379000), kMirror(7, 17))
}
