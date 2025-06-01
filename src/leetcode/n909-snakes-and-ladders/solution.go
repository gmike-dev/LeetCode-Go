package n909_snakes_and_ladders

import "math"

func snakesAndLadders(board [][]int) int {
	var (
		n     = len(board)
		count = make([]int, n*n)
		f     func(int, int)
	)
	for i := 0; i < n*n; i++ {
		count[i] = math.MaxInt
	}
	f = func(pos, num int) {
		if pos == n*n-1 {
			count[pos] = min(count[pos], num)
			return
		}
		if count[pos] <= num {
			return
		}
		count[pos] = num
		for dice := 6; dice > 0; dice-- {
			target := pos + dice
			if target >= n*n {
				continue
			}
			if target != n*n-1 {
				tr, tc := getRC(target, n)
				if board[tr][tc] != -1 {
					target = board[tr][tc] - 1
				}
			}
			f(target, num+1)
		}
	}
	f(0, 0)
	if count[n*n-1] == math.MaxInt {
		return -1
	}
	return count[n*n-1]
}

func getRC(pos, n int) (int, int) {
	var r, c int
	r = n - pos/n - 1
	if (n-r-1)%2 == 0 {
		c = pos % n
	} else {
		c = n - pos%n - 1
	}
	return r, c
}
