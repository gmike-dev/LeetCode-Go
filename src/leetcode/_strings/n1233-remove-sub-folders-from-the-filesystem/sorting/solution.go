package sorting

import (
	"slices"
	"strings"
)

func removeSubfolders(folder []string) []string {
	slices.Sort(folder)
	var res []string
	res = append(res, folder[0])
	for i := 1; i < len(folder); i++ {
		if !strings.HasPrefix(folder[i], res[len(res)-1]+"/") {
			res = append(res, folder[i])
		}
	}
	return res
}
