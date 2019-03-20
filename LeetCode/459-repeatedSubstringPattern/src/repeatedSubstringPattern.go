package src

import "strings"

func repeatedSubstringPattern(s string) bool {
	if len(s) < 2 {
		return false
	}

	newStr := s + s
	return strings.Contains(newStr[1:len(newStr)-1], s)
}
