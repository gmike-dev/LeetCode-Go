package n2081_sum_of_k_mirror_numbers

func kMirror(k int, n int) int64 {
	var (
		count = 0
		pow   = 1
		s     int64
	)
	for m := 1; ; m++ {
		oddLen := m%2 != 0
		for i := pow; i < 10*pow; i++ {
			p := buildPalindrome(i, oddLen)
			if isPalindrome(p, k) {
				s += p
				count++
				if count == n {
					return s
				}
			}
		}
		if !oddLen {
			pow *= 10
		}
	}
}

func buildPalindrome(num int, oddLen bool) int64 {
	res := int64(num)
	if oddLen {
		num /= 10
	}
	for num != 0 {
		d := int64(num % 10)
		res = res*10 + d
		num /= 10
	}
	return res
}

var baseK = [100]int64{}

func isPalindrome(num int64, k int) bool {
	n := 0
	for num != 0 {
		baseK[n] = num % int64(k)
		n++
		num /= int64(k)
	}
	for i := 0; i+i < n; i++ {
		if baseK[i] != baseK[n-i-1] {
			return false
		}
	}
	return true
}
