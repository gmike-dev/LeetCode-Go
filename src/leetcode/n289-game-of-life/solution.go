package n289_game_of_life

func gameOfLife(board [][]int) {
	n, m := len(board), len(board[0])
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			nc := neighbours(board, i, j)
			if board[i][j] == 0 && nc == 3 {
				next[i][j] = 1
			} else if board[i][j] == 1 && (nc == 2 || nc == 3) {
				next[i][j] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		copy(board[i], next[i])
	}
}

func neighbours(board [][]int, r, c int) int {
	n, m := len(board), len(board[0])
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			nr := r + i
			nc := c + j
			if nr >= 0 && nc >= 0 && nr < n && nc < m {
				count += board[nr][nc]
			}
		}
	}
	return count - board[r][c]
}
