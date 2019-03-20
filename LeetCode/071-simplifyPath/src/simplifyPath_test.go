package src

import "testing"

func TestOK(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"/home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}
	for _, test := range tests {
		if got := simplifyPath(test.input); got != test.want {
			t.Errorf("simplifyPath(%s) = %s", test.input, got)
		}
	}
}
