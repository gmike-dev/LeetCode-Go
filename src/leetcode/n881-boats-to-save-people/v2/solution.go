package n881_boats_to_save_people_v2

func numRescueBoats(people []int, limit int) int {
	cnt := make([]int, limit+1)
	for _, p := range people {
		cnt[p]++
	}

	left := 0
	right := limit
	boats := 0
	for left <= right {
		if cnt[left] == 0 {
			left++
		} else if cnt[right] == 0 {
			right--
		} else {
			boats++
			cnt[right]--
			if cnt[left] != 0 && left+right <= limit {
				cnt[left]--
			}
		}
	}
	return boats
}
