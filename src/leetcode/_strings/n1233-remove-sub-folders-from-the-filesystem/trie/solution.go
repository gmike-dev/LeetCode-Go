package trie

import "strings"

func removeSubfolders(folder []string) []string {
	root := newTrieNode("")
	for _, f := range folder {
		node := root
		s := strings.Split(f, "/")
		substring := false
		for _, sub := range s {
			if sub != "" {
				child, ok := node.child[sub]
				if ok {
					if child.folder != "" {
						substring = true
						break
					}
				} else {
					child = newTrieNode(sub)
					node.child[sub] = child
				}
				node = child
			}
		}
		if substring {
			continue
		}
		node.folder = f
	}
	return collectFolders(root, []string{})
}

func collectFolders(node *trieNode, res []string) []string {
	if node.folder != "" {
		return append(res, node.folder)
	}
	for _, v := range node.child {
		res = collectFolders(v, res)
	}
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
