package n2000_reverse_prefix_of_word

import "strings"

func reversePrefix(word string, ch byte) string {
	index := strings.IndexByte(word, ch)
	if index == -1 || index == 0 {
		return word
	}
	chars := []byte(word)
	for i := 0; i+i <= index; i++ {
		chars[i], chars[index-i] = chars[index-i], chars[i]
	}
	return string(chars)
}
