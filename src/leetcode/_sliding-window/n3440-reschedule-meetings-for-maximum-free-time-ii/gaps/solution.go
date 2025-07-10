package fake_meetings

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)

	gaps := make([]int, n+1)
	gaps[0] = startTime[0]
	for i := 1; i < n; i++ {
		gaps[i] = startTime[i] - endTime[i-1]
	}
	gaps[n] = eventTime - endTime[n-1]

	maxGapAtRight := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		maxGapAtRight[i] = max(maxGapAtRight[i+1], gaps[i+1])
	}

	maxGapAtLeft := 0
	res := 0
	for i := 0; i < n; i++ {
		freeTime := gaps[i] + gaps[i+1]
		duration := endTime[i] - startTime[i]
		if maxGapAtLeft >= duration || maxGapAtRight[i+1] >= duration {
			freeTime += duration
		}
		res = max(res, freeTime)
		maxGapAtLeft = max(maxGapAtLeft, gaps[i])
	}
	return res
}
