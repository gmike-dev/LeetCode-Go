package n3508_implement_router

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	router := Constructor(3)
	assert.True(t, router.AddPacket(1, 4, 90))
	assert.True(t, router.AddPacket(2, 5, 90))
	assert.False(t, router.AddPacket(1, 4, 90))
	assert.True(t, router.AddPacket(3, 5, 95))
	assert.True(t, router.AddPacket(4, 5, 105))
	assert.Equal(t, "[2,5,90]", tst.List2Str(router.ForwardPacket()))
	assert.True(t, router.AddPacket(5, 2, 110))
	assert.Equal(t, 1, router.GetCount(5, 100, 110))
}
