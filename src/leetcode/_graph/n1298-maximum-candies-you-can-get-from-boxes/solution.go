package n1298_maximum_candies_you_can_get_from_boxes

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	var (
		q   []int
		ans int
	)
	has := make([]bool, len(status))
	for _, b := range initialBoxes {
		has[b] = true
		if status[b] != 0 {
			q = append(q, b)
		}
	}
	for len(q) != 0 {
		box := q[0]
		q = q[1:]
		ans += candies[box]
		for _, b := range containedBoxes[box] {
			has[b] = true
			if status[b] != 0 {
				q = append(q, b)
			}
		}
		for _, b := range keys[box] {
			if status[b] == 0 {
				status[b] = 1
				if has[b] {
					q = append(q, b)
				}
			}
		}
	}
	return ans
}
