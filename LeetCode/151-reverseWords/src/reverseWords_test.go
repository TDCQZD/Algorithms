package src

import "testing"

func TestOk(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"  hello world!  ", "world! hello"},
		{"the sky is blue", "blue is sky the"},
		{"a good   example", "example good a"},
	}
	for _, test := range tests {
		if got := reverseWords(test.input); got != test.want {
			t.Errorf("reverseWords(%s) = %s want=%s", test.input, got, test.want)
		}
	}
}
