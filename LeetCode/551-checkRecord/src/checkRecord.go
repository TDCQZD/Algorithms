package src

import "strings"

func checkRecord(s string) bool {
	countA := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			countA++
		}
		if s[i] == 'L' {
			if i+2 < len(s) && s[i+1] == 'L' && s[i+2] == 'L' {
				return false
			}
		}
	}

	if countA > 1 {
		return false
	}
	return true
}

func checkRecord2(s string) bool {
	return strings.IndexByte(s, 'A') == strings.LastIndexByte(s, 'A') && !strings.Contains(s, "LLL")
}
