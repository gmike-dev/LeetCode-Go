package n166_fraction_to_recurring_decimal

import (
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var res string
	if (numerator > 0) != (denominator > 0) {
		res += "-"
	}
	numerator = abs(numerator)
	denominator = abs(denominator)
	res += strconv.Itoa(numerator / denominator)
	rem := numerator % denominator
	if rem == 0 {
		return res
	}
	res += "."
	used := make(map[int]int)
	for rem != 0 {
		if pos, ok := used[rem]; ok {
			return res[:pos] + "(" + res[pos:] + ")"
		}
		used[rem] = len(res)
		rem *= 10
		res += strconv.Itoa(rem / denominator)
		rem %= denominator
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
