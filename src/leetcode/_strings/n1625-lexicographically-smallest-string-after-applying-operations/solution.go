package n1625_lexicographically_smallest_string_after_applying_operations

import (
	"maps"
	"slices"
)

func findLexSmallestString(s string, a int, b int) string {
	used := make(map[string]bool)

	add := func(s string) string {
		x := []byte(s)
		for i := 1; i < len(x); i += 2 {
			x[i] = (x[i]-'0'+byte(a))%10 + '0'
		}
		return string(x)
	}
	shift := func(s string) string {
		return s[b:] + s[:b]
	}

	var dfs func(string)
	dfs = func(s string) {
		if used[s] {
			return
		}
		used[s] = true
		dfs(add(s))
		dfs(shift(s))
	}

	dfs(s)
	return slices.Min(slices.Collect(maps.Keys(used)))
}
