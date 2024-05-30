package leetcode

func countTriplets(arr []int) int {
	n := len(arr)
	xor := make([]int, n+1)
	for i := 0; i < n; i++ {
		xor[i+1] = xor[i] ^ arr[i]
	}
	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			x := xor[j] ^ xor[i]
			for k := j; k < n; k++ {
				y := xor[k+1] ^ xor[j]
				if x == y {
					count++
				}
			}
		}
	}
	return count
}
