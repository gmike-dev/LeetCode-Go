package _52_Open_the_Lock

import "strconv"

func openLock(deadends []string, target string) int {
	const N = 10000
	used := make([]bool, N)
	for i := 0; i < len(deadends); i++ {
		item, _ := strconv.Atoi(deadends[i])
		if item == 0 {
			return -1
		}
		used[item] = true
	}
	queueArr := [N]int{}
	queue := queueArr[:1]
	used[0] = true

	tryEnqueue := func(q []int, i int) []int {
		if !used[i] {
			q = append(q, i)
			used[i] = true
		}
		return q
	}

	targetNum, _ := strconv.Atoi(target)
	length := 0
	for len(queue) > 0 {
		items := make([]int, len(queue))
		copy(items, queue)
		queue = queueArr[:0]
		for i := 0; i < len(items); i++ {
			num := items[i]
			if num == targetNum {
				return length
			}
			k := 1
			for i := 0; i < 4; i++ {
				d := num / k % 10
				queue = tryEnqueue(queue, num+((d+1)%10-d)*k)
				queue = tryEnqueue(queue, num+((d+9)%10-d)*k)
				k *= 10
			}
		}
		length++
	}
	return -1
}
