package n2444_count_subarrays_with_fixed_bounds

func countSubarrays(nums []int, minK int, maxK int) int64 {
	left := -1
	lastMinK := -1
	lastMaxK := -1
	numberOfSubarrays := int64(0)
	for right := 0; right < len(nums); right++ {
		var x = nums[right]
		if x < minK || x > maxK {
			left = right
		} else {
			if x == minK {
				lastMinK = right
			}
			if x == maxK {
				lastMaxK = right
			}
			if lastMinK != -1 && lastMaxK != -1 {
				numberOfSubarrays += int64(max(0, min(lastMinK, lastMaxK)-left))
			}
		}
	}
	return numberOfSubarrays
}
