package merge_sort

func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)
	a := make([]int, n+1)
	for i, num := range nums {
		if num == target {
			a[i+1] = a[i] + 1
		} else {
			a[i+1] = a[i] - 1
		}
	}
	return countInversions(a)
}

func mergeSortAndCount(a []int) int64 {
	n := len(a)
	b := make([]int, 0, n)
	m := n / 2
	l, r := 0, m
	var count int64
	for l < m && r < n {
		if a[l] >= a[r] {
			b = append(b, a[l])
			l++
		} else {
			b = append(b, a[r])
			r++
			count += int64(m - l)
		}
	}
	b = append(b, a[l:m]...)
	b = append(b, a[r:]...)
	copy(a, b)
	return count
}

func countInversions(a []int) int64 {
	n := len(a)
	if n <= 1 {
		return 0
	}
	m := n / 2
	var count int64
	count += countInversions(a[:m])
	count += countInversions(a[m:])
	count += mergeSortAndCount(a)
	return count
}
