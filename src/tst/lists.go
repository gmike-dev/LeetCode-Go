package tst

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ListEqual[T any](t *testing.T, expected string, actual []T) {
	assert.Equal(t, expected, List2String(actual))
}

func MatrixEqual[T any](t *testing.T, expected string, actual [][]T) {
	assert.Equal(t, expected, Mx2String(actual))
}

func Str2List(s string) []int {
	s = s[1 : len(s)-1]
	if len(s) == 0 {
		return []int{}
	}
	items := strings.Split(s, ",")
	res := make([]int, len(items))
	for i := range res {
		num, err := strconv.ParseInt(items[i], 10, 32)
		if err != nil {
			panic(err)
		}
		res[i] = int(num)
	}
	return res
}

func Str2Matrix(s string) [][]int {
	s = s[1 : len(s)-1]
	if len(s) == 0 {
		return [][]int{}
	}
	r, err := regexp.Compile(`(\[.*?\])*`)
	if err != nil {
		panic(err)
	}
	ss := r.FindAllString(s, -1)
	res := make([][]int, len(ss))
	for i, row := range ss {
		res[i] = Str2List(row)
	}
	return res
}

func List2String[T any](a []T) string {
	s := strings.Builder{}
	s.WriteString("[")
	for i, x := range a {
		if i != 0 {
			s.WriteString(",")
		}
		s.WriteString(fmt.Sprintf("%v", x))
	}
	s.WriteString("]")
	return s.String()
}

func Mx2String[T any](a [][]T) string {
	s := strings.Builder{}
	s.WriteString("[")
	for i, x := range a {
		if i != 0 {
			s.WriteString(",")
		}
		s.WriteString(List2String(x))
	}
	s.WriteString("]")
	return s.String()
}
