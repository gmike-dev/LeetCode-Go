package trie

import "strings"

func removeSubfolders(folder []string) []string {
	root := newTrieNode("")
	add := func(f string) {
		node := root
		s := strings.Split(f, "/")
		for i := 1; i < len(s); i++ {
			sub := s[i]
			child, ok := node.child[sub]
			if !ok {
				child = newTrieNode(sub)
				node.child[sub] = child
			} else if child.folder != "" {
				return
			}
			node = child
		}
		node.folder = f
		node.child = nil
	}
	for _, f := range folder {
		add(f)
	}
	var res []string
	var collect func(*trieNode)
	collect = func(node *trieNode) {
		if node.folder != "" {
			res = append(res, node.folder)
		} else {
			for _, v := range node.child {
				collect(v)
			}
		}
	}
	collect(root)
	return res
}

type trieNode struct {
	val    string
	child  map[string]*trieNode
	folder string
}

func newTrieNode(val string) *trieNode {
	return &trieNode{
		val:   val,
		child: make(map[string]*trieNode),
	}
}
