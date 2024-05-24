package leetcode

func maxScoreWords(words []string, letters []byte, score []int) int {
	n := len(words)
	var letterFreq [26]int
	for _, letter := range letters {
		letterFreq[letter-'a']++
	}
	calcScore := func(m int) int {
		s := 0
		var freq [26]int
		for i := 0; m != 0; i++ {
			if m&1 != 0 {
				for _, letter := range words[i] {
					letter -= 'a'
					if freq[letter] == letterFreq[letter] {
						return 0
					}
					s += score[letter]
					freq[letter]++
				}
			}
			m >>= 1
		}
		return s
	}
	maxScore := 0
	for mask := 1; mask < 1<<n; mask++ {
		maxScore = max(maxScore, calcScore(mask))
	}
	return maxScore
}
