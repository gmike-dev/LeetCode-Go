package leetcode

func sortColors(nums []int) {
	cnt := [3]int{}
	for _, num := range nums {
		cnt[num]++
	}
	j := 0
	for i := 0; i < 3; i++ {
		for k := 0; k < cnt[i]; k++ {
			nums[j] = i
			j++
		}
	}
}
