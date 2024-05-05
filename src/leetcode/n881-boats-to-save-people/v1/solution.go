package n881_boats_to_save_people_v1

import "slices"

func numRescueBoats(people []int, limit int) int {
	slices.Sort(people)
	left := 0
	right := len(people) - 1
	boats := 0
	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		boats++
		right--
	}
	return boats
}
