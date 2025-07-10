package fake_meetings

import "slices"

func maxFreeTime(eventTime int, startTime []int, endTime []int) int {
	n := len(startTime)
	startTime = slices.Concat([]int{0}, startTime, []int{eventTime})
	endTime = slices.Concat([]int{0}, endTime, []int{eventTime})
	maxGapAtRight := make([]int, n+2)
	for i := n; i >= 0; i-- {
		maxGapAtRight[i] = max(maxGapAtRight[i+1], startTime[i+1]-endTime[i])
	}
	maxGapAtLeft := 0
	res := 0
	for i := 1; i <= n; i++ {
		duration := endTime[i] - startTime[i]
		if maxGapAtLeft >= duration || maxGapAtRight[i+1] >= duration {
			res = max(res, startTime[i+1]-endTime[i-1])
		} else {
			res = max(res, startTime[i+1]-endTime[i-1]-duration)
		}
		maxGapAtLeft = max(maxGapAtLeft, startTime[i]-endTime[i-1])
	}
	return res
}
