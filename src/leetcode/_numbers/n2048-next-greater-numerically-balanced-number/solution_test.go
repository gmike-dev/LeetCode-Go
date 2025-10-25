package n2048_next_greater_numerically_balanced_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_nextBeautifulNumber1(t *testing.T) {
	assert.Equal(t, 22, nextBeautifulNumber(1))
}

func Test_nextBeautifulNumber2(t *testing.T) {
	assert.Equal(t, 1333, nextBeautifulNumber(1000))
}

func Test_nextBeautifulNumber3(t *testing.T) {
	assert.Equal(t, 3133, nextBeautifulNumber(3000))
}

func Test_nextBeautifulNumber4(t *testing.T) {
	assert.Equal(t, 1224444, nextBeautifulNumber(748601))
}
