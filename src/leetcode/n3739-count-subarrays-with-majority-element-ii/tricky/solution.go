package tricky

func countMajoritySubarrays(nums []int, target int) int64 {
	n := len(nums)
	count := make([]int, 2*n+2)
	countPrefixSum := make([]int, 2*n+2)
	s := n + 1
	result := int64(0)
	countPrefixSum[s] = 1
	count[s] = 1
	for _, num := range nums {
		if num == target {
			s++
		} else {
			s--
		}
		count[s]++
		countPrefixSum[s] = countPrefixSum[s-1] + count[s]
		result += int64(countPrefixSum[s-1])
	}
	return result
}
