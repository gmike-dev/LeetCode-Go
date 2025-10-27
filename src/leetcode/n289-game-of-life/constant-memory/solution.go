package n289_game_of_life

func gameOfLife(board [][]int) {
	n, m := len(board), len(board[0])
	d := []int{-1, 0, 1, 1, -1, 1, 0, -1, -1}
	neighbours := func(r, c int) int {
		count := 0
		for i := range 8 {
			nr := r + d[i]
			nc := c + d[i+1]
			if nr >= 0 && nc >= 0 && nr < n && nc < m {
				count += board[nr][nc] & 1
			}
		}
		return count
	}
	for i := range n {
		for j := range m {
			nc := neighbours(i, j)
			if board[i][j]&1 == 0 && nc == 3 || board[i][j]&1 == 1 && (nc == 2 || nc == 3) {
				board[i][j] |= 2
			}
		}
	}
	for i := range n {
		for j := range m {
			board[i][j] >>= 1
		}
	}
}
