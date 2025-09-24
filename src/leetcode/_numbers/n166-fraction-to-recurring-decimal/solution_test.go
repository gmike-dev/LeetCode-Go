package n166_fraction_to_recurring_decimal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fractionToDecimal(t *testing.T) {
	assert.Equal(t, "0.(12)", fractionToDecimal(4, 33))
	assert.Equal(t, "0.(012)", fractionToDecimal(4, 333))
	assert.Equal(t, "0.1(6)", fractionToDecimal(1, 6))
	assert.Equal(t, "-0.58(3)", fractionToDecimal(7, -12))
	assert.Equal(t, "0", fractionToDecimal(0, 3))
	assert.Equal(t, "0", fractionToDecimal(0, -3))
	assert.Equal(t, "0.00(000000465661289042462740251655654056577585848337359161441621040707904997124914069194026549138227660723878669455195477065427143370461252966751355553982241280310754777158628319049732085502639731402098131932683780538602845887105337854867197032523144157689601770377165713821223802198558308923834223016478952081795603341592860749337303449725)", fractionToDecimal(1, 214748364))
	assert.Equal(t, "0.01(6)", fractionToDecimal(1, 60))
	assert.Equal(t, "0.001(6)", fractionToDecimal(1, 600))
	assert.Equal(t, "0.00(5)", fractionToDecimal(1, 180))
	assert.Equal(t, "0.(001)", fractionToDecimal(1, 999))
	assert.Equal(t, "0.0000000004656612873077392578125", fractionToDecimal(-1, -2147483648))
	assert.Equal(t, "2", fractionToDecimal(2, 1))
	assert.Equal(t, "21.6", fractionToDecimal(216, 10))
}
