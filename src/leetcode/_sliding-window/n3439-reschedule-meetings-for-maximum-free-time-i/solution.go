package n3439_reschedule_meetings_for_maximum_free_time_i

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	m := n + 1
	gaps := make([]int, m)

	gaps[0] = startTime[0]
	for i := 1; i < n; i++ {
		gaps[i] = startTime[i] - endTime[i-1]
	}
	gaps[m-1] = eventTime - endTime[n-1]

	gapWindowSum := 0
	for i := 0; i < k; i++ {
		gapWindowSum += gaps[i]
	}
	res := 0
	for i := k; i < m; i++ {
		gapWindowSum += gaps[i]
		res = max(res, gapWindowSum)
		gapWindowSum -= gaps[i-k]
	}
	return res
}
