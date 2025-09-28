package n3699_number_of_zigzag_arrays_i

const mod = int(1e9) + 7

func zigZagArrays(n int, l int, r int) int {
	less := make([]int, 2002)
	more := make([]int, 2002)
	for x := l; x <= r; x++ {
		less[x] = 1
		more[x] = 1
	}
	prefix := make([]int, 2002)
	suffix := make([]int, 2002)
	for range n - 1 {
		for x := l; x <= r; x++ {
			prefix[x] = (prefix[x-1] + less[x]) % mod
		}
		for x := r; x >= l; x-- {
			suffix[x] = (suffix[x+1] + more[x]) % mod
		}
		for x := l; x <= r; x++ {
			less[x] = suffix[x+1]
			more[x] = prefix[x-1]
		}
	}
	res := 0
	for x := l; x <= r; x++ {
		res = (res + less[x]) % mod
		res = (res + more[x]) % mod
	}
	return res
}
