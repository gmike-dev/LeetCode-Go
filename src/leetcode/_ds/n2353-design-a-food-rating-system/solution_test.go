package n2353_design_a_food_rating_system

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	fr := Constructor(
		[]string{"kimchi", "miso", "sushi", "moussaka", "ramen", "bulgogi"},
		[]string{"korean", "japanese", "japanese", "greek", "japanese", "korean"},
		[]int{9, 12, 8, 15, 14, 7})
	assert.Equal(t, "kimchi", fr.HighestRated("korean"))
	assert.Equal(t, "ramen", fr.HighestRated("japanese"))
	fr.ChangeRating("sushi", 16)
	assert.Equal(t, "sushi", fr.HighestRated("japanese"))
	fr.ChangeRating("ramen", 16)
	assert.Equal(t, "ramen", fr.HighestRated("japanese"))
}
