package n1915_number_of_wonderful_substrings

func wonderfulSubstrings(word string) int64 {
	counter := make([]int64, 1024)
	counter[0] = 1
	mask := 0
	ans := int64(0)
	for _, c := range word {
		mask ^= 1 << (c - 'a')
		ans += counter[mask]
		counter[mask]++
		for i := 0; i < 10; i++ {
			ans += counter[mask^(1<<i)]
		}
	}
	return ans
}
