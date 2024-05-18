package n979_distribute_coins_in_binary_tree

import (
	. "LeetCode-Go/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistributeCoins1(t *testing.T) {
	assert.Equal(t, 2, distributeCoins(StrToTree("3,0,0")))
}

func TestDistributeCoins2(t *testing.T) {
	assert.Equal(t, 3, distributeCoins(StrToTree("0,3,0")))
}

func TestDistributeCoins3(t *testing.T) {
	assert.Equal(t, 5, distributeCoins(StrToTree("2,0,3,0,null,1,0")))
}
