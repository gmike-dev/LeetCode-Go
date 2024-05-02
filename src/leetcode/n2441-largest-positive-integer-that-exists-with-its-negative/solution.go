package n2441_largest_positive_integer_that_exists_with_its_negative

func findMaxK(nums []int) int {
	exists := [1001]int{}
	for _, n := range nums {
		if n < 0 {
			exists[-n] |= 1
		} else {
			exists[n] |= 2
		}
	}
	for i := 1000; i >= 0; i-- {
		if exists[i] == 3 {
			return i
		}
	}
	return -1
}
