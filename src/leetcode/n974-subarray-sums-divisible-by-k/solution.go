package leetcode

func subarraysDivByK(nums []int, k int) int {
	prefixCount := make([]int, k)
	prefixCount[0] = 1
	prefixMod := 0
	subarrays := 0
	for _, num := range nums {
		prefixMod = ((prefixMod+num)%k + k) % k
		subarrays += prefixCount[prefixMod]
		prefixCount[prefixMod]++
	}
	return subarrays
}
