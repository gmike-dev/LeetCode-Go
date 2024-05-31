package leetcode

func singleNumber(nums []int) []int {
	s := make(map[int]int, 0)
	for _, num := range nums {
		if s[num] == 1 {
			delete(s, num)
		} else {
			s[num] = 1
		}
	}
	result := make([]int, 0)
	for num := range s {
		result = append(result, num)
	}
	return result
}
