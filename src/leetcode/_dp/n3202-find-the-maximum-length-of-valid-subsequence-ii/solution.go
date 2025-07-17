package n3202_find_the_maximum_length_of_valid_subsequence_ii

func maximumLength(nums []int, k int) int {
	res := 0
	for mod := range k {
		dp := make([]int, k)
		for _, num := range nums {
			m := num % k
			prev := (mod - m + k) % k
			dp[m] = dp[prev] + 1
			res = max(res, dp[m])
		}
	}
	return res
}
