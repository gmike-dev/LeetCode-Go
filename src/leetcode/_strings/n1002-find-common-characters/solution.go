package leetcode

func commonChars(words []string) []string {
	cnt := [26]int{}
	for _, c := range words[0] {
		cnt[c-'a']++
	}
	n := len(words)
	for i := 1; i < n; i++ {
		wc := [26]int{}
		for _, c := range words[i] {
			wc[c-'a']++
		}
		for j := 0; j < 26; j++ {
			cnt[j] = min(cnt[j], wc[j])
		}
	}

	common := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < cnt[i]; j++ {
			common = append(common, string(rune('a'+i)))
		}
	}
	return common
}
