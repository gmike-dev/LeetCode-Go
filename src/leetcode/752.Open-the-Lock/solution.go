package _52_Open_the_Lock

import "strconv"

func openLock(deadends []string, target string) int {
	t, _ := strconv.Atoi(target)
	used := make([]bool, 10000)
	for i := 0; i < len(deadends); i++ {
		item, _ := strconv.Atoi(deadends[i])
		if item == 0 {
			return -1
		}
		used[item] = true
	}
	q := make([]int, 0)
	q = append(q, 0)
	used[0] = true
	length := 0
	for len(q) > 0 {
		items := make([]int, len(q))
		copy(items, q)
		q = nil
		for i := 0; i < len(items); i++ {
			item := items[i]
			if item == t {
				return length
			}
			k := 1
			for i := 0; i < 4; i++ {
				d := item / k % 10
				next1 := item - d*k + (d+1)%10*k
				next2 := item - d*k + (10+(d-1))%10*k
				if !used[next1] {
					q = append(q, next1)
					used[next1] = true
				}
				if !used[next2] {
					q = append(q, next2)
					used[next2] = true
				}
				k *= 10
			}
		}
		length++
	}
	return -1
}
