package n514_freedom_trail

import "math"

var ring string
var key string
var dp [][]int
var n int
var m int

func findRotateSteps(_ring string, _key string) int {
	ring, key = _ring, _key
	n, m = len(ring), len(key)
	dp = make([][]int, n)
	InitDp()
	return Steps(0, 0)
}

func Steps(ringIndex int, keyIndex int) int {
	if keyIndex == m {
		return 0
	}
	if dp[ringIndex][keyIndex] != -1 {
		return dp[ringIndex][keyIndex]
	}
	result := math.MaxInt
	for i := 0; i < n; i++ {
		if ring[i] == key[keyIndex] {
			steps := 1 + Steps(i, keyIndex+1)
			if i != ringIndex {
				steps += Dist(i, ringIndex)
			}
			result = min(result, steps)
		}
	}
	dp[ringIndex][keyIndex] = result
	return result
}

func Dist(i int, j int) int {
	d := Abs(i - j)
	return min(d, n-d)
}

func InitDp() {
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
