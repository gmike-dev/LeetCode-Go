package n1992_find_all_groups_of_farmland

func findFarmland(land [][]int) [][]int {
	n := len(land)
	m := len(land[0])
	k := 0
	result := make([][]int, 0, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if land[i][j] == 1 {
				left, right, up, down := near(land, n, m, i, j)
				if up == 0 && left == 0 {
					result = append(result, []int{i, j, 0, 0})
					k++
					land[i][j] = k
				} else {
					land[i][j] = max(up, left)
				}
				if right == 0 && down == 0 {
					result[land[i][j]-1][2] = i
					result[land[i][j]-1][3] = j
				}
			}
		}
	}
	return result
}

func near(land [][]int, n int, m int, i int, j int) (int, int, int, int) {
	left := 0
	if j > 0 {
		left = land[i][j-1]
	}
	right := 0
	if j != m-1 {
		right = land[i][j+1]
	}
	up := 0
	if i > 0 {
		up = land[i-1][j]
	}
	down := 0
	if i != n-1 {
		down = land[i+1][j]
	}
	return left, right, up, down
}
