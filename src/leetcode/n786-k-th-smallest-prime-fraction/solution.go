package n786_k_th_smallest_prime_fraction

import "slices"

type Fraction struct {
	Nom int
	Den int
}

func kthSmallestPrimeFractionV1(arr []int, k int) []int {
	n := len(arr)
	fractions := make([]Fraction, 0, n*(n-1)/2)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fractions = append(fractions, Fraction{arr[i], arr[j]})
		}
	}
	slices.SortFunc(fractions, func(a, b Fraction) int {
		return a.Nom*b.Den - b.Nom*a.Den
	})
	return []int{fractions[k-1].Nom, fractions[k-1].Den}
}

func kthSmallestPrimeFractionV2(arr []int, k int) []int {
	n := len(arr)
	fractions := make([]Fraction, 0, n*(n-1)/2)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			fractions = append(fractions, Fraction{arr[i], arr[j]})
		}
	}
	result := QuickSelect(fractions, k-1, func(a, b Fraction) int {
		return a.Nom*b.Den - b.Nom*a.Den
	})
	return []int{result.Nom, result.Den}
}

func QuickSelect(a []Fraction, k int, cmp func(a, b Fraction) int) Fraction {
	l := 0
	r := len(a) - 1
	for l < r {
		var q = Partition(a, l, r, cmp)
		if k <= q {
			r = q
		} else {
			l = q + 1
		}
	}
	return a[k]
}

func Partition(a []Fraction, l int, r int, cmp func(a, b Fraction) int) int {
	var pivot = a[l+(r-l)/2]
	var i = l - 1
	var j = r + 1
	for {
		i++
		for cmp(a[i], pivot) < 0 {
			i++
		}
		j--
		for cmp(a[j], pivot) > 0 {
			j--
		}
		if i >= j {
			return j
		}
		a[i], a[j] = a[j], a[i]
	}
}
