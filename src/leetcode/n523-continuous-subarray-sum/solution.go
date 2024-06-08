package leetcode

func checkSubarraySum(nums []int, k int) bool {
	s := make(map[int]int)
	s[0] = -1
	prefixMod := 0
	for i, n := range nums {
		prefixMod = (prefixMod + n) % k
		mod, ok := s[prefixMod]
		if ok {
			if i-mod > 1 {
				return true
			}
		} else {
			s[prefixMod] = i
		}
	}
	return false
}
