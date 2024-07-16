package core

import "strings"

func isEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func isNotEmpty(s string) bool {
	return !isEmpty(s)
}
func same(a, b string) bool {
	return a == b
}
func diff(a, b string) bool {
	return !same(a, b)
}
