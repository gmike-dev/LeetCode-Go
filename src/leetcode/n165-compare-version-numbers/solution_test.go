package n165_compare_version_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareVersion1(t *testing.T) {
	actual := compareVersion("1.01", "1.001")
	expect := 0
	assert.Equal(t, expect, actual)
}

func TestCompareVersion2(t *testing.T) {
	actual := compareVersion("1.0", "1.0.0")
	expect := 0
	assert.Equal(t, expect, actual)
}

func TestCompareVersion3(t *testing.T) {
	actual := compareVersion("0.1", "1.1")
	expect := -1
	assert.Equal(t, expect, actual)
}

func TestCompareVersion5(t *testing.T) {
	actual := compareVersion("1.2.3", "1.2")
	expect := 1
	assert.Equal(t, expect, actual)
}
