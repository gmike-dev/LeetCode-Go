package leetcode

import "strings"

type trieNode struct {
	next [26]*trieNode
	stop bool
}

func buildTrie(dictionary []string) *trieNode {
	t := &trieNode{}
	for _, w := range dictionary {
		n := t
		for _, c := range w {
			if n.stop {
				break
			}
			c -= 'a'
			if n.next[c] == nil {
				n.next[c] = &trieNode{}
			}
			n = n.next[c]
		}
		n.stop = true
	}
	return t
}

func replaceWords(dictionary []string, sentence string) string {
	t := buildTrie(dictionary)
	words := strings.Split(sentence, " ")
	result := make([]string, len(words))
	for i, w := range words {
		n := t
		for j, c := range w {
			n = n.next[c-'a']
			if n == nil {
				break
			}
			if n.stop {
				result[i] = w[:j+1]
				break
			}
		}
		if result[i] == "" {
			result[i] = w
		}
	}
	return strings.Join(result, " ")
}
