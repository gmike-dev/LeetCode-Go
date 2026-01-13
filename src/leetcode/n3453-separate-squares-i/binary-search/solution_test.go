package binary_search

import (
	"LeetCode-Go/tst"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_separateSquares1(t *testing.T) {
	assert.InEpsilon(t, 1.00000, separateSquares(tst.Str2Matrix("[[0,0,1],[2,2,1]]")), 1e-5)
}

func Test_separateSquares2(t *testing.T) {
	assert.InEpsilon(t, 1.16667, separateSquares(tst.Str2Matrix("[[0,0,2],[1,1,1]]")), 1e-5)
}

func Test_separateSquares3(t *testing.T) {
	assert.InEpsilon(t, 954521423.80202, separateSquares(tst.Str2Matrix("[[522261215,954313664,225462],"+
		"[628661372,718610752,10667],[619734768,941310679,44788],[352367502,656774918,289036],"+
		"[860247066,905800565,100123],[817623994,962847576,71460],[691552058,782740602,36271],"+
		"[911356,152015365,513881],[462847044,859151855,233567],[672324240,954509294,685569]]")), 1e-5)
}
