package src

import (
	"testing"
)

func TestOk(t *testing.T) {
	var tests = []struct {
		s1   string
		s2   string
		want bool
	}{
		{"", "", true},
		{"ab", "ab", true},
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	}
	for _, test := range tests {
		if got := checkInclusion(test.s1, test.s2); got != test.want {
			t.Errorf("checkInclusion(%s, %s) = %v", test.s1, test.s2, got)
		}
	}

}
