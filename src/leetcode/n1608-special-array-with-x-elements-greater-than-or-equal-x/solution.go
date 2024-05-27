package leetcode

func specialArray(nums []int) int {
	n := len(nums)
	cnt := make([]int, n+1)
	for _, num := range nums {
		cnt[min(n, num)]++
	}
	for i := 0; i < len(cnt); i++ {
		if i == n {
			return i
		}
		n -= cnt[i]
	}
	return -1
}
