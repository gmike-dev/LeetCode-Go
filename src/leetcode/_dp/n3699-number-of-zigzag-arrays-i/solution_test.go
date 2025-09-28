package n3699_number_of_zigzag_arrays_i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_zigZagArrays1(t *testing.T) {
	assert.Equal(t, 2, zigZagArrays(3, 4, 5))
}

func Test_zigZagArrays2(t *testing.T) {
	assert.Equal(t, 10, zigZagArrays(3, 1, 3))
}

func Test_zigZagArrays3(t *testing.T) {
	assert.Equal(t, 182, zigZagArrays(3, 8, 14))
}

func Test_zigZagArrays4(t *testing.T) {
	assert.Equal(t, 476062160, zigZagArrays(44, 462, 1570))
}

func Test_zigZagArrays5(t *testing.T) {
	assert.Equal(t, 702864277, zigZagArrays(54, 1998, 2000))
}
