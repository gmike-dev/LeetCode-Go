package n752_open_the_Lock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenLock1(t *testing.T) {
	actual := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	expected := 6
	assert.Equal(t, expected, actual)
}

func TestOpenLock2(t *testing.T) {
	actual := openLock([]string{"8888"}, "0009")
	expected := 1
	assert.Equal(t, expected, actual)
}

func TestOpenLock3(t *testing.T) {
	actual := openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	expected := -1
	assert.Equal(t, expected, actual)
}

func TestOpenLock4(t *testing.T) {
	actual := openLock([]string{"0234"}, "1234")
	expected := 10
	assert.Equal(t, expected, actual)
}

func TestOpenLock5(t *testing.T) {
	actual := openLock([]string{"0000"}, "8888")
	expected := -1
	assert.Equal(t, expected, actual)
}
