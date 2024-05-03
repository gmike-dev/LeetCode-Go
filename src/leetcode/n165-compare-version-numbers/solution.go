package n165_compare_version_numbers

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	n := len(v1)
	m := len(v2)
	for i := 0; i < max(n, m); i++ {
		var d1, d2 int
		if i < n {
			d1, _ = strconv.Atoi(v1[i])
		} else {
			d1 = 0
		}
		if i < m {
			d2, _ = strconv.Atoi(v2[i])
		} else {
			d2 = 0
		}
		if d1 < d2 {
			return -1
		}
		if d1 > d2 {
			return 1
		}
	}
	return 0
}
