package leetcode

func numRabbits(answers []int) int {
	counter := make(map[int]int)
	for _, num := range answers {
		counter[num]++
	}
	var sum int
	for num, count := range counter {
		if num < count {
			sum += (count + num) / (num + 1) * (num + 1)
		} else {
			sum += num + 1
		}
	}
	return sum
}
