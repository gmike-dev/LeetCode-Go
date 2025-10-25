package n2048_next_greater_numerically_balanced_number

import (
	"math"
	"slices"
)

func nextBeautifulNumber(n int) int {
	res := math.MaxInt
	for m := 1; m < 1<<6; m++ {
		var digits []int
		digit := 1
		for k := m; k != 0; k >>= 1 {
			if k&1 != 0 {
				for range digit {
					digits = append(digits, digit)
				}
			}
			digit++
		}
		if len(digits) < 8 {
			for {
				num := toNum(digits)
				if num > n {
					res = min(res, num)
				}
				if !nextPermutation(digits) {
					break
				}
			}
		}
	}
	return res
}

func toNum(digits []int) int {
	res := 0
	for i := range digits {
		res = res*10 + digits[i]
	}
	return res
}

func nextPermutation(a []int) bool {
	n := len(a)
	var left int
	for left = n - 2; left >= 0; left-- {
		if a[left] < a[left+1] {
			break
		}
	}
	if left < 0 {
		return false
	}
	var right int
	for right = n - 1; left < right; right-- {
		if a[right] > a[left] {
			break
		}
	}
	a[left], a[right] = a[right], a[left]
	slices.Reverse(a[left+1:])
	return true
}
